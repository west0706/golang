package main

import "fmt"

func main() {
	var a = make([][]string, 3)

	for i := range a {
		a[i] = make([]string, 3)
	}
	fmt.Println(a, len(a), cap(a))
}
