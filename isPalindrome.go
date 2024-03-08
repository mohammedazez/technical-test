package main

import "strings"

func IsPalindrome(str string) bool {
	lowerCase := strings.ToLower(str)
	for i := 0; i < len(string(lowerCase))/2; i++ {
		if string(lowerCase[i]) != string(lowerCase[len(lowerCase)-1-i]) {
			return false
		}
	}
	return true
}
