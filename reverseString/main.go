package main

import "fmt"

func main() {
	str := "hello"
	reverseStr := ""

	for i := len(str) - 1; i >= 0; i-- {
		reverseStr = reverseStr + string(str[i]) //str[i] is byte and thats why we use string since reversestr is in string
	}
	fmt.Println(reverseStr)
}
