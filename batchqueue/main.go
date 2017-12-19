package main

import (
	"fmt"
	"time"
)

type Obj int

var (
	queue = make(chan Obj, 10)
)

func send() {
	for i := 0; i < 100000; i++ {
		queue <- Obj(i)
	}
	queue <- Obj(-1)
}

func receive() {
	rcvLoop()
}

func batch(buf []Obj) {
	fmt.Println("#", buf)
}

func rcvLoop() {
	buf := make([]Obj, 0)
	for {
		select {
		case obj := <-queue:
			if obj == -1 {
				return
			}
			buf = append(buf, obj)
		}
		if len(buf) >= 5 {
			batch(buf)
			buf = buf[:0]
		}
	}
}

func main() {
	ts := time.Now()
	go send()
	receive()
	fmt.Println(time.Since(ts))
}

