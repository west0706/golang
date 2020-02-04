package main

import "fmt"

func main() {
	var temp string
	temp = "HelloWorld"
	fmt.Println(temp)

	var tvt [10]byte
	tt := []byte(temp)

	for index, val := range tt {
		fmt.Println(index, val)
		tvt[index] = val
	}
}
