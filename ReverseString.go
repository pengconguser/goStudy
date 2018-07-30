package main

import (
	"fmt"
)

func main() {

	str := "hello go"

	result := ReverseString(str)

	fmt.Printf("%v", result)
}

func ReverseString(str string) string {
	str_len := len(str)
	str_rune := []string(str)

	result_str := make([]rune, 0)

	for i := 0; i < str_len; i++ {
		result_str = append(result_str, str_rune[str_len-i-1])
	}

	// fmt.Printf("%v \n", right)
	return result_str
}
