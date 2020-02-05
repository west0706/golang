package main

import "fmt"

func main() {
	x := 123
	y := 456.1
	z := "hint"

	printIT(x, y, z)
}

func printIT(v ...interface{}) {
	buff := v
	for _, value := range buff {
		fmt.Println(value)
	}

}
