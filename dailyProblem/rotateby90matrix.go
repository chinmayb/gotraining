package main

import "fmt"

/*
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [[7,4,1],[8,5,2],[9,6,3]]

1 2       3 1
3 4 ----> 4 2
*/

func rotate(matrix [][]int) {
	for i, row := range matrix {
		for j, _ := range row {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println("")
	}

}

func main() {
	a := [][]int{
		{1, 2, 3},
		{
			4, 5, 6,
		},
		{
			7, 8, 9,
		},
	}
	rotate(a)
}
