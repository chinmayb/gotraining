package main

import "fmt"

/*
Find the lowest positive integer that does not exist in the array. The array can contain duplicates and negative
numbers as well.

For example, the input [3, 4, -1, 1] should give 2. The input [1, 2, 0] should give 3.
 */

 func segregatePositiveNumbers(a []int) []int{
 	var x []int
 		for _, ele := range a {
 			if ele > 0{
 				x = append(x, ele)
			}
		}
		return x
 }

 // Using hash table with value as key is one way.
func sf(a []int)  int{
  	l := len(a) -1
  	fmt.Println(a)
  	for _,elem := range a {
  		if elem  < l {
			a[elem -1] = -a[elem -1]
		}
	}
	fmt.Println(a)
	for i,elem := range a {
		if elem > 0 {
			return i+1
		}
	}
	return 0
}

func main() {

	//a := []int{3, 4, -1, 1}
	a := []int{2, 4, -1, 1, 3}

	a = segregatePositiveNumbers(a)

	fmt.Println(sf(a))
	

}