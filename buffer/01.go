package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type Record struct {
	err  error
	data int
}

func main() {
	c := make(chan Record)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	go getData1(c, ctx)
	go getData2(c, ctx)

	total := 0

	for i := 0; i < 2; i++ {
		d := <-c
		if d.err != nil {
			cancel()
			break
		}
	}

	fmt.Println(total)

}

func getData1(c chan Record, ctx context.Context) {
	lc := make(chan int, 1)
	go func() {
		time.Sleep(time.Second)
		lc <- 100
	}()
	select {
	case a := <-lc:
		c <- Record{nil, a}
	case <-ctx.Done():
		c <- Record{errors.New("sorry"), 0}
	}
}

func getData2(c chan Record, ctx context.Context) {
	lc := make(chan int, 1)
	go func() {
		time.Sleep(time.Second)
		lc <- 100
	}()
	select {
	case a := <-lc:
		c <- Record{nil, a}
	case <-ctx.Done():
		c <- Record{errors.New("sorry"), 0}
	}
}
