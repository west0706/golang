package main

import (
	"fmt"
	"sync"
)

func myTest_1(wg *sync.WaitGroup, mychannel <-chan string) {
	defer wg.Done()
	buff := <-mychannel
	fmt.Println("myTest_1", buff)
}

func myTest_2(wg *sync.WaitGroup, mychannel <-chan string) {
	defer wg.Done()
	buff := <-mychannel
	fmt.Println("myTest_2", buff)
}

func myTest_3(wg *sync.WaitGroup, mychannel <-chan string) {
	defer wg.Done()
	buff := <-mychannel
	fmt.Println("myTest_3", buff)
}

func myTestIn(wg *sync.WaitGroup, mychannel chan<- string) {
	defer wg.Done()
	buff := "myTestIn"
	mychannel <- buff
}

func main() {
	var wait sync.WaitGroup
	defer func() {
		fmt.Println("main gorutine End!!")
	}()
	myChannel_1 := make(chan string)

	wait.Add(1)
	go myTestIn(&wait, myChannel_1)
	wait.Add(1)
	go myTestIn(&wait, myChannel_1)
	wait.Add(1)
	go myTest_1(&wait, myChannel_1)

	wait.Add(1)
	go myTest_2(&wait, myChannel_1)
	wait.Wait()
}
