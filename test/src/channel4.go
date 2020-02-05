package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		ch <- 123
	}()

	var i int
	i = <-ch
	fmt.Println(i)
}
