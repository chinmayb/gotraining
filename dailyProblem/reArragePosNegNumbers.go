package main

import "fmt"

//Rearrange positive and negative numbers
//in O(n) time and O(1) extra space

//[-1, 2, -3, 4, 5, 6, -7, 8, 9], then the output should be [9, -7, 8, -3, 5, -1, 2, 4, 6]

func swap(a , b *int) {
	a, b = b, a
	//*a, *b = *b, *a
	}

func main(){

	var arr = []int{-1, 2, -3, 4, 5, 6, -7, -8, 9}
	j := len(arr) - 1
	for i:=0;i < len(arr);i++ {
		for arr[i] < 0 && i < j{
			i++
		}
		for arr[j] > 0 && j > i{
			j--
		}
		swap(&arr[i], &arr[j])
	}
	fmt.Println(arr)
}