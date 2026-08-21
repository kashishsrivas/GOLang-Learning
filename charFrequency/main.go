package main

import "fmt"

//character frequency

func main() {
	str := "banana"
	//find how many time "a" appear
	count := 0
	for i := 0; i < len(str); i++ {
		if str[i] == 'a' {
			count++
		}
	}
	fmt.Println(count)
}
