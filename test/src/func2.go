package main

import "fmt"

func test() func(temp string) string {
	result := "CloserTest"
	return func(temp string) string {
		return temp + result
	}
}

func main() {

	myFunc := test()
	funcResult := myFunc("param")
	fmt.Println(funcResult)
	fmt.Println(funcResult)
}
