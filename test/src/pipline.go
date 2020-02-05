package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	// wait := sync.WaitGroup{}

	s1 := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(s1)
	inValue := make(chan int)
	out := make(chan map[int]int)
	sumResult := make(chan int)

	myMap := make(map[int]int)

	stopFlag := false
	mutex := sync.RWMutex{}

	go func() {
		defer func() {
			mutex.Lock()
			// wait.Done()
			fmt.Println("In Funtion")
			mutex.Unlock()
		}()

		for {
			if stopFlag {
				// close(in)
				break
			}
			number := rand.Intn(10)
			inValue <- number
			fmt.Println("in function", number)
			// time.Sleep(time.Second)
		}
	}()

	go func() {
		defer func() {
			mutex.Lock()
			// wait.Done()
			fmt.Println("Check Funtion")
			mutex.Unlock()

		}()
		var temp int
		for {

			temp = <-inValue
			val, flag := myMap[temp]
			if flag {
				stopFlag = true
				out <- myMap
				fmt.Println("check function", myMap)
				break
			} else {
				myMap[temp] = val
				// fmt.Println(temp, myMap[temp])
			}

		}

	}()

	go func() {
		defer func() {
			mutex.Lock()
			// wait.Done()
			fmt.Println("Result Funtion")
			mutex.Unlock()

		}()
		var temp map[int]int

		temp = <-out
		// fmt.Println("result function", temp)
		var sumKey int
		sumKey = 0

		for key := range temp {
			sumKey += key
		}
		sumResult <- sumKey
		fmt.Println("result function", sumKey)

	}()
	// wait.Wait()

	var pipResult int
	pipResult = <-sumResult
	fmt.Println(pipResult)

}
