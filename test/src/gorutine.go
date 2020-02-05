package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(8)
	fmt.Println(runtime.GOMAXPROCS(0))

	defer fmt.Println("Main END")
	var wait sync.WaitGroup

	wait.Add(1)
	temp := "Hello "

	go func() { //스레드 생성 (go만 붙이면 펑션이 스레드로 동작)
		defer wait.Done()
		fmt.Println(temp + "First Thread")
	}()

	fmt.Println("Main Thread")
	wait.Wait() // add 1 한 뒤 0일때까지 기다림
}
