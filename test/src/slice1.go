package main

import "fmt"

func main() {
	var a []int = make([]int, 5)
	var b = make([]int, 5, 6)
	c := make([]int, 5, 10)

	fmt.Println(a, b, c)
	fmt.Println(a, len(a), cap(a))
	fmt.Println(b, len(b), cap(b))
	fmt.Println(c, len(c), cap(c))

	a = append(a, 1, 2, 3)
	b = append(b, 1, 2, 3)
	c = append(c, 1, 2, 3, 4, 5, 6)
	c = append(c, 7, 8, 9, 10)

	fmt.Println(a, b, c)
	fmt.Println(a, len(a), cap(a))
	fmt.Println(b, len(b), cap(b))
	fmt.Println(c, len(c), cap(c))

	fmt.Println(len(a), len(b), len(c))
	fmt.Println(cap(a), cap(b), cap(c))
}
