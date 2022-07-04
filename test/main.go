package main

import (
	"fmt"
	"time"
)

func add(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
}

// timeout problem recurrent
func test2() {
	ch := make(chan int, 10)
	go add(ch)
	for {
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("timeout")
			return
		case result := <-ch:
			fmt.Println(result) // if ch not empty, time.After will nerver exec
		}
	}
}
func main() {
	test2()
}
