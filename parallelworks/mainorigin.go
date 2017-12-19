package main

import (
	"fmt"
	"time"
)

func dealWork(id int) int {
	time.Sleep(time.Second)
	return id
}

func prepre(id int) int {
	time.Sleep(time.Millisecond * 20)
	return id
}

func after(id int) {
	time.Sleep(time.Millisecond * 50)
	fmt.Println("finish:", id)
}

func main() {
	s := time.Now()
	for i := 0; i < 100; i++ {
		p := prepre(i)
		d := dealWork(p)
		after(d)
	}

	fmt.Println(time.Since(s))
}
