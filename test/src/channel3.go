package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wait sync.WaitGroup
	defer func() {
		fmt.Println("main goroutine end!")
	}()
	myChannel_1 := make(chan int)
	myChannel_2 := make(chan string)

	wait.Add(1)
	go func() {
		defer wait.Done()
		fmt.Println("Func1")
		time.Sleep(time.Second)
		myChannel_1 <- 1
		close(myChannel_1)
	}()
}
