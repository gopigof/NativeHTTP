package main

import (
	"fmt"
	"testing"
)

func BenchmarkFibMultiple(b *testing.B) {
	ch := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
		}
		quit <- 0
	}()
	fibonacciMultiple(ch, quit)
}
