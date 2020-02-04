package main

import "fmt"

func main() {
	var a map[string]string
	a = map[string]string{"a": "aaa", "b": "bbb", "c": "ccc"}

	fmt.Println(a, a["a"], len(a))
}
