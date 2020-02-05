package main

import (
	"fmt"
	"time"
)

func main() {

	defer fmt.Println("main goroutine end!")
	myChannel := make(chan int)

	go func() {
		time.Sleep(time.Second)
		myChannel <- 1
		close(myChannel)
	}()

	go func() {
		val, flag := <-myChannel
		fmt.Println(val, flag)
		val, flag = <-myChannel
		fmt.Println(val, flag)
	}()

	val, flag := <-myChannel
	fmt.Println(val, flag)

}
