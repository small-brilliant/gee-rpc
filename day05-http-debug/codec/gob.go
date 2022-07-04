package codec

import (
	"bufio"
	"encoding/gob"
	"io"
	"log"
)

type GobCodec struct {
	conn io.ReadWriteCloser //conn 是由构建函数传入，是通过TCP或者Unix建立socket时得到的连接实例
	buf  *bufio.Writer      //buf 是为了防止阻塞而创建的带缓冲的 Writer，一般这么做能提升性能
	dec  *gob.Decoder
	enc  *gob.Encoder
}

func NewGobCodec(conn io.ReadWriteCloser) Codec {
	buf := bufio.NewWriter(conn)
	return &GobCodec{
		conn: conn,
		buf:  buf,
		dec:  gob.NewDecoder(conn),
		enc:  gob.NewEncoder(buf),
	}
}

func (g GobCodec) Close() error {
	return g.conn.Close()
}

func (g GobCodec) ReadHeader(h *Header) error {
	return g.dec.Decode(h)
}

func (g GobCodec) ReadBody(body interface{}) error {
	return g.dec.Decode(body)
}

func (g GobCodec) Write(h *Header, body interface{}) error {
	err := g.enc.Encode(h)
	if err != nil {
		log.Println("rpc codec: gob error encoding header:", err)
		return err
	}
	err = g.enc.Encode(body)
	if err != nil {
		log.Println("rpc codec: gob error encoding body:", err)
		return err
	}
	defer func() {
		_ = g.buf.Flush()
		if err != nil {
			_ = g.Close()
		}
	}()
	return nil
}
