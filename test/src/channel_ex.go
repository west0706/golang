package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 1 ~ 100까지 난수를 계속 생성
// 한번이라도 같은 난수가 나오면 기존에 있던 난수를 모두 더해 리턴

func main() {

	defer func() {
		fmt.Println("Main go routine ended!")
	}()

	var wait sync.WaitGroup
	var rand_list []int

	wait.Add(1)
	go func() {
		defer wait.Done()

		//random number
		for {
			s1 := rand.NewSource(time.Now().UnixNano())
			rand := rand.New(s1)
			randomNumber := rand.Intn(100)
			// fmt.Println(randomNumber)

			fmt.Println(checkValueInArray(randomNumber, rand_list))

			if checkValueInArray(randomNumber, rand_list) {
				sum_rand := sumAllRandNum(rand_list)
				fmt.Println("Sum of random numbers is ", sum_rand)
				break
			}

			rand_list = append(rand_list, randomNumber)
			// fmt.Println(rand_list)

			time.Sleep(time.Second)
		}
	}()

	wait.Wait()
}

func checkValueInArray(rand_num int, rand_list []int) bool {
	fmt.Println(rand_num)
	fmt.Println(rand_list)

	return_bool := false

	for _, v := range rand_list {
		// fmt.Println(v)

		if rand_num == v {
			return_bool = true
		}
	}
	return return_bool
}

func sumAllRandNum(rand_list []int) int {
	var sum_return int

	sum_return = 0

	for _, v := range rand_list {
		sum_return += v
	}
	return sum_return
}
