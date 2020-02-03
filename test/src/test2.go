package main

import (
	"fmt"
	// "unsafe"
)

func myCal(a int, b int) (int, int) {
	return (a + b), (a - b)
}

func main() {
	// var x int = 21
	// var y int = 40
	// x := 77
	// y := 88

	// var b1 bool = true
	// var b2 bool = false
	// var b3, b4 bool
	// b3, b4 = false, false

	// // fmt.Printf("The value of int is %d %d", x, y)
	// fmt.Println("Size of x: %d", unsafe.Sizeof(x))
	// fmt.Println("Size of y: %d", unsafe.Sizeof(y))
	// fmt.Println("Bools are %d, %d", b1, b2)

	temp_1, temp_2 := myCal(10, 10)
	fmt.Println(temp_1, temp_2)
}
