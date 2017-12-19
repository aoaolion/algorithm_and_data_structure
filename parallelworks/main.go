package main

import (
	"fmt"
	"sync"
	"time"
)

func dealWork(ch chan int, ch2 chan int) {
	time.Sleep(time.Second)
	select {
	case i := <-ch:
		ch2 <- i
	}
}

func prepre(id int, ch chan int) {
	time.Sleep(time.Millisecond * 20)
	ch <- id
}

func after(ch2 chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 50)
	select {
	case i := <-ch2:
		fmt.Println("finish:", i)
	}
}

func main() {
	s := time.Now()
	var wg sync.WaitGroup
	ch := make(chan int, 10)
	ch2 := make(chan int, 2)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		prepre(i, ch)
		go dealWork(ch, ch2)
		go after(ch2, &wg)
	}
	wg.Wait()
	fmt.Println(time.Since(s))
}
