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

// Nomor 3
func PrintTriangle(n int) string {
	var result string
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			result += "*"
		}
		result += "\n"
	}
	return result
}

// Nomor 4
func Factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * Factorial(n-1)
}

func main() {
	// IsPalindrome
	fmt.Println("Number 1")
	fmt.Println(IsPalindrome("abba"))
	fmt.Println(IsPalindrome("kodok"))
	fmt.Println(IsPalindrome("radar"))
	fmt.Println(IsPalindrome("hello"))
	fmt.Println("----------------------------------")

	// FindMax
	fmt.Println("Number 2")
	input := []int{3, 5, 1, 9, 2}
	fmt.Println(FindMax(input))
	fmt.Println("----------------------------------")

	// PrintTriangle
	fmt.Println("Number 3")
	fmt.Println(PrintTriangle(5))
	fmt.Println("----------------------------------")

	// Factorial
	fmt.Println("Number 4")
	fmt.Println(Factorial(5))
	fmt.Println("----------------------------------")
}
