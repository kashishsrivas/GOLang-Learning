package main

import "fmt"

//Find the frequency of every character in a string.

func main() {
	str := "hello"
	frequency := make(map[byte]int) // map: "h" : 1 (key value pair)

	for i := 0; i < len(str); i++ {
		frequency[str[i]]++
	}

	//ch = key, count=value -> h= key, 1=count
	for ch, count := range frequency {
		fmt.Printf("%c: %d\n", ch, count)
	}
}
