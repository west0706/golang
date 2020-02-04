package main

import "fmt"

func main() {

	// var a [5]int = [5]int{11, 22, 33, 44, 55}
	// a[2] = 7
	// fmt.Println(a)

	buff := make([]int, 5)
	buff[0] = 11
	buff[1] = 22
	buff[2] = 33
	fmt.Println(buff)

	buff_1 := make([][]int, 5, 5)
	buff_1[0] = buff
	buff_1[1] = []int{3, 4, 5, 6, 7}
	fmt.Println(buff_1)
	fmt.Println("================")

	for i, val := range buff_1 {
		fmt.Println(i, val)
	}
}
