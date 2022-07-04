package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var set = make(map[int]bool, 0)

func printOnce(index int, num int) {

	if _, ok := set[num]; !ok {
		fmt.Println(index, num)
	}
	set[num] = true
}

func TestPrint(t *testing.T) {
	for i := 0; i < 10; i++ {
		go printOnce(i, 100)
	}
	time.Sleep(time.Second)
}

func TestPrint2(t *testing.T) {
	var num = 0
	var locker sync.Mutex

	for i := 0; i < 100000; i++ {
		go func(i int) {
			locker.Lock()
			defer locker.Unlock()

			num++
			fmt.Printf("goroutine %d: num = %d\n", i, num)
		}(i)
	}

	time.Sleep(time.Second)
}
