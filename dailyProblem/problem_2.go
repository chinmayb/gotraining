package main

import "fmt"

/*
For example, if our input was [1, 2, 3, 4, 5],
the expected output would be [120, 60, 40, 30, 24]. If our input was [3, 2, 1], the expected output would be [2, 3, 6].
*/


func getProduct(a []int) int{
	var prod = 1
	for _, elem := range a {
		prod = elem*prod
	}
	return prod
}

func main() {

	a := []int{1, 2, 3, 4, 5}
	newA := [5]int{}
	prod := getProduct(a)
	for i, _ := range a {
		newA[i] = prod / a[i]
	}
	fmt.Println(newA)
}
