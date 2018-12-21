package main

import (
	"fmt"
)

func max(a, b int) int{
	if a < b{
		return b
	}else{
		return a
	}
}
// O(nk) time | O(nk) space
func maxProfileWIthKTrasact(prices []int, k int) {
	profits := make([][]int, k+1)
	profits[0] = make([]int, len(prices))

	for t := 1; t<= k; t++ {
		profits[t] = make([]int, len(prices))
		maxThusFar := -1111111 // large negetive number
		for d:=1 ; d< len(prices); d++ {
			maxThusFar = max(maxThusFar, profits[t-1][d-1] - prices[d-1])
			profits[t][d] = max(maxThusFar + prices[d], profits[t][d-1])
		}
	}

	fmt.Println(profits)
}

func main(){
	prices := []int{5,11,3,50,60,90}
	k := 2
	maxProfileWIthKTrasact(prices, k)
}
