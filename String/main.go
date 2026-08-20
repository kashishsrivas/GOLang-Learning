package main

import "fmt"

func main() {
	//count how many vowels are present in a string

	str := "programming"
	count := 0

	for i := 0; i < len(str); i++ {
		//str[i] will access the indiviual character
		character := str[i]
		if character == 'a' || character == 'e' || character == 'i' || character == 'o' || character == 'u' {
			count++
		}
	}
	fmt.Println(count)
}
