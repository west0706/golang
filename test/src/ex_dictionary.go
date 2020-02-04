package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	f, err := ioutil.ReadFile("resource/test_word.txt")

	var dict = make(map[string]int)
	var str_temp string

	if err != nil {
		panic(err)
	} else {
		fmt.Println(len(f))
		str_temp = string(f)
	}

	re := regexp.MustCompile(`[^a-zA-Z ]+`)
	str_temp2 := re.ReplaceAllString(str_temp, "")

	str_temp3 := strings.ToLower(str_temp2)

	fmt.Println(str_temp3)

	for _, value := range strings.Split(str_temp3, " ") {

		val_tmp := strings.Trim(value, "\"")
		val_tmp = strings.Trim(val_tmp, ".")
		val_tmp = strings.Trim(val_tmp, ":")
		val_tmp = strings.Trim(val_tmp, ",")

		_, exists := dict[val_tmp]
		if exists {
			dict[val_tmp] += 1
		} else {
			dict[val_tmp] = 1
		}
	}

	fmt.Println(dict)

	for k, v := range dict {
		fmt.Println(k, v)
	}

}
