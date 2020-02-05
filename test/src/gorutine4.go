package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	// runtime.GOMAXPROCS(runtime.NumCPU())
	// wg := sync.WaitGroup //대기그룹 생성

	// fmt.Println(runtime.NumCPU())
	// fmt.Println(wg)

	// GET STRING
	f, err := ioutil.ReadFile("resource/test_word_large.txt")

	var str_temp string

	if err != nil {
		panic(err)
	} else {
		fmt.Println(len(f))
		str_temp = string(f)
	}

	// fmt.Println(str_temp)
	str_line := strings.Split(str_temp, "\n")
	fmt.Println(len(str_line))
	step1_line_num := len(str_line) / 2
	fmt.Println(step1_line_num)

	// var dict1 = make(map[string]int)
	// var dict2 = make(map[string]int)

	list1 := make([]string, step1_line_num)
	list2 := make([]string, step1_line_num)

	for i, v := range str_line {
		if i < step1_line_num {
			list1 = append(list1, v)
		} else {
			list2 = append(list2, v)
		}
	}

	// fmt.Println(list1)
	// fmt.Println(list2)

	for i, v := range trimString(list1) {
		fmt.Println(i, v)
	}

	// for i := 0; i < 10000; i++ {
	// 	wg.Add(1)
	// 	go func(n int) {
	// 		fmt.Println(n)
	// 		wg.Done()
	// 	}(i)
	// }

	// wg.Wait()
	// fmt.Println("the end")
}

func trimString(str_list []string) map[string]int {
	// re := regexp.MustCompile(`[^a-zA-Z ]+`)
	// str_temp := re.ReplaceAllString(str_list, "")
	// str_temp2 := strings.ToLower(str_list)
	fmt.Println(str_list)

	var temp_map map[string]int

	for _, v := range str_list {
		temp_map[v] = 1
	}

	return temp_map
}
