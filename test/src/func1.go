package main

import "fmt"

func main() {
	fmt.Println(test(1, 2, 3, 4, 5))

	func() {
		fmt.Println("FUNC")
	}()

	func(temp string) {
		fmt.Println("FUNC", temp)
	}("HI")

}

func test(temp ...int) (result int) {
	fmt.Println(temp)
	result = 100
	return
}
