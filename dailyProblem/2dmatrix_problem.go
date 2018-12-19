package main

import (
	"math/rand"
	"fmt"
)

/*
[ 1 2 4 3]
[ 2 6 7 4]
[ 2 9 8 2]

output

[1  *  * ]
[*  *  *]
[*  *  *]

 */
type matrix [][]int

func newMatrix(m, n int) matrix{
	var m1 matrix
	for i:=0;i<m; i++{
		for  j:=0;j<n;j++{
			m1[i][j] = rand.Intn(10)
		}
	}
	return  m1
}
func (m *matrix)string() string{
	return ""

}

 func main(){
 	m1:= newMatrix(3,4)
 	fmt.Printf("%s", m1)

 }