package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- 10
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			c2 <- "Hello, World!"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			select {
			case i := <-c1:
				fmt.Println("c1 : ", i)
			case i := <-c2:
				fmt.Println("c2 : ", i)
			}
		}
	}()

	time.Sleep(10 * time.Second)
}
