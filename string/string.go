package main

import "fmt"

func main() {
	var str string = "中国"
	for i := 0; i < len(str); i++ {
		fmt.Println(string(str[i]))
	}
	for index, value := range str {
		fmt.Println(index, string(value))
	}
}
