package main

import (
	"fmt"
	"strconv"
)

/*
Given a 32-bit signed integer, reverse digits of an integer.

Note:
Assume we are dealing with an environment that could only store integers within the 32-bit signed integer range: [−2 pow 31,  2 pow (31 − 1)]. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.


Example 1:

Input: x = 123
Output: 321

Input: x = 120
Output: 21


*/

func reverse(x int) int {
	const minInt32 = -2147483648
	const maxInt32 = 2147483647
	var revNum string
	var initIndex int
	strnum := strconv.Itoa(x)
	if x < 0 {
		initIndex = 1
	} else if x == 0 {
		return x
	}
	for i := len(strnum) - 1; i >= initIndex; i-- {
		n1 := strnum[i]
		revNum = revNum + fmt.Sprintf("%c", n1)
	}

	if initIndex == 1 {
		revNum = "-" + revNum
	}
	num, _ := strconv.Atoi(revNum)

	fmt.Println(num)
	fmt.Println(minInt32)

	// check on 32 bit it overflows or not
	if num > maxInt32 || num < minInt32 {
		return 0
	}
	return num
}

func main() {
	x := -2147483412
	fmt.Print(reverse(x))
}
