package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Output: Because nums[0] + nums[1] == 9, we return [0, 1].
*/

func twoSum(nums []int, target int) []int {
	keymap := make(map[int]int)
	for i, k := range nums {
		keymap[k] = i
	}
	for i, n := range nums {
		if index, ok := keymap[target-n]; ok && index != i {
			return []int{i, index}
		}
	}
	return []int{}
}

func main() {
	var a string
	_, err := fmt.Scanln(&a)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
	var target int
	if _, err := fmt.Scanln(&target); err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
	numberArr := strings.Split(a, ",")
	var num []int
	for _, i := range numberArr {
		n, err := strconv.Atoi(i)
		if err != nil {
			continue
		}
		num = append(num, n)
	}
	fmt.Print(twoSum(num, target))
}
