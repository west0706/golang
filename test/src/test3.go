package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	temp := 0

	for {
		if temp > 100000 {
			break
		}
		fmt.Println(temp)
		temp += 1

	}
}
