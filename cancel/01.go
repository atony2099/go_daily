package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int)

	go watch(c)

	time.Sleep(time.Second * 4)

	c <- 0

}

func watch(stop chan int) {

	for {
		select {

		case <-time.After(time.Second):
			fmt.Println("hello")
		case <-stop:
			fmt.Println("stop")

		}
	}

}
