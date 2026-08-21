package main

import "fmt"

func main() {
	str := "madam"
	rev := ""

	for i := len(str) - 1; i >= 0; i-- {
		rev = rev + string(str[i])
	}
	if str == rev {
		fmt.Println("palindrome")
	} else {
		fmt.Println("not palindrome")
	}

	fmt.Println(rev)
}
