package main

import "fmt"

func main() {
	//PalindromeCheck using two pointers

	str := "madam"
	left := 0
	right := len(str) - 1
	isPalindrome := true

	for left < right {
		if str[left] != str[right] {
			isPalindrome = false
			break
		}

		left++
		right--
	}

	fmt.Println(isPalindrome)
}
