package main

import (
	"fmt"
	"time"
)

func access(ch chan int) {
	time.Sleep(time.Second)
	fmt.Println("Accessing channel")

	for i := range ch {
		fmt.Println("Channel received ", i)
		time.Sleep(time.Second)
	}
}

func main() {
	ch := make(chan int, 4)
	defer close(ch)

	go access(ch)
	for i := 0; i < 9; i++ {
		ch <- i
		fmt.Println("Pushed to channel ", i)
	}

	time.Sleep(time.Second * 3)
	fmt.Println("Done")
}
