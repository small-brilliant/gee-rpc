package main

import (
	"fmt"
	"reflect"
)

func chanDir(v interface{}) {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Chan {
		switch t.ChanDir() {
		case reflect.RecvDir:
			fmt.Println("receive")
		case reflect.SendDir:
			fmt.Println("send")
		case reflect.BothDir:
			fmt.Println("receive and send")

		}
	}
}
func main() {
	var a <-chan int
	var b chan<- int
	var c chan int
	chanDir(a)
	chanDir(b)
	chanDir(c)
}
