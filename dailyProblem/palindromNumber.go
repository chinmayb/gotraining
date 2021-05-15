package main

import "fmt"

/*
123

3
without string comparison
*/
func isPalindrome(x int) bool {
	const minInt32 = -2147483648
	const maxInt32 = 2147483647
	var temp int
	orig := x
	if x < 0 || x != 0 && x%10 == 0 {
		return false
	}

	for x != 0 {
		remainder := x % 10
		x = x / 10

		if temp > maxInt32 || temp < minInt32 {
			fmt.Print("mot suuported")
			return false
		}
		temp = temp*10 + remainder
	}
	if temp == orig {
		return true
	}
	return false
}

func main() {
	x := 101123123123
	fmt.Print(isPalindrome(x))
}
