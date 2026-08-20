package main

import "fmt"

func main() {
	//count vowels and consonants separately.

	str := "programming"
	countVowel := 0
	countConso := 0

	for i := 0; i < len(str); i++ {
		if str[i] == 'a' || str[i] == 'e' || str[i] == 'i' || str[i] == 'o' || str[i] == 'u' {
			countVowel++
		} else {
			countConso++
		}
	}
	fmt.Println("vowel: ", countVowel)
	fmt.Println("consonants: ", countConso)

}
