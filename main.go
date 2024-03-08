package main

import (
	"fmt"
	"strings"
)

// Nomor 1
func IsPalindrome(str string) bool {
	lowerCase := strings.ToLower(str)
	for i := 0; i < len(string(lowerCase))/2; i++ {
		if string(lowerCase[i]) != string(lowerCase[len(lowerCase)-1-i]) {
			return false
		}
	}
	return true
}

// Nomor 2
func FindMax(arr []int) int {
	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}
	return max
}

func main() {
	// IsPalindrome
	fmt.Println(IsPalindrome("abba"))
	fmt.Println(IsPalindrome("kodok"))
	fmt.Println(IsPalindrome("radar"))
	fmt.Println(IsPalindrome("hello"))

	// FindMax
	input := []int{3, 5, 1, 9, 2}
	fmt.Println(FindMax(input))

}
