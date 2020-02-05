package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start!")

	defer fmt.Println("main goroutine end!")
	myChannel := make(chan int)

	go func() {
		time.Sleep(time.Second)
		myChannel <- 1
	}()

	fmt.Println("Receive!")
	temp := <-myChannel
	fmt.Println(temp)

}
