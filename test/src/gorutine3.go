package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	wg := new(sync.WaitGroup) //대기그룹 생성

	fmt.Println(runtime.NumCPU())
	fmt.Println(wg)

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(n int) {
			fmt.Println(n)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("the end")

}
