package operations

import (
	"math/rand"
	//"strconv"
	"fmt"
)

type Matrix struct {
	Id     int
	matrix [][]int
}
type Matrices []Matrix

var id int = 1

func generate_random_num() int {
	return rand.Intn(100)
}
func (m Matrix) getCords() (int, int) {
	return len(m.matrix), len(m.matrix[0])
}

func (m Matrix) Show() {
	fmt.Println("\n=================================")
	fmt.Println("id      :  ", m.Id)
	fmt.Print("Matrix  :  ")
	for i := range m.matrix {
		fmt.Println("")
		for j := range m.matrix[i] {
			fmt.Print("\t")
			fmt.Print(m.matrix[i][j], " ")
		}
	}
}

func List() {

}
func increment() {
	id++
}

func New(r int, c int) Matrix {
	fmt.Print("Creating....")
	ss := make([][]int, r)
	for i := range ss {
		ss[i] = make([]int, c)
	}
	defer increment()
	return Matrix{Id: id, matrix: ss}
}

func (m Matrix) Insert() {
	//Creates a matrix
	fmt.Print("inserting....")
	for i := range m.matrix {
		for j := range m.matrix[i] {
			m.matrix[i][j] = generate_random_num()
		}
	}
}

func isValid() bool {
	return true
}

func Subtract(m1 Matrix, m2 Matrix) Matrix {
	m1_row_len, m1_col_len := m1.getCords()
	if isValid() != false {

	}
	m := New(m1_row_len, m1_col_len)

	for i := range m.matrix {
		for j := range m.matrix[i] {
			m.matrix[i][j] = m1.matrix[i][j] - m2.matrix[i][j]
		}
	}
	return m
	// res_ss := make([][]int, m1_row_len)
	// for i:=0; i< m1_row_len; i++ {
	// 	res_ss[i] = make([] int, m1_col_len)
	// 	for j:=0; j < m1_col_len; j++ {
	// 		res_ss[i][j] = m1.matrix[i][j] - m2.matrix[i][j]
	// 	}
	// }
	// defer increment()
	//return Matrix{Id: id, matrix: res_ss}
}

func Add(m1 Matrix, m2 Matrix) Matrix {
	m1_row_len, m1_col_len := m1.getCords()
	if isValid() != false {

	}
	m := New(m1_row_len, m1_col_len)

	for i := range m.matrix {
		for j := range m.matrix[i] {
			m.matrix[i][j] = m1.matrix[i][j] + m2.matrix[i][j]
		}
	}
	return m
}

func iterate(m Matrix, opfunc func(int, int)) {
	for i := range m.matrix {
		for j := range m.matrix[i] {
			opfunc(i, j)
		}
	}
}

func (m Matrix) ToString() (result string) {
	// Use this , since client can write the outout
	// to a file.
	//result := ""
	iterate(m, func(i int, j int) {
		delim := "\t"
		if j == len(m.matrix[0])-1 {
			delim = "\n"
		}
		result += fmt.Sprintf("%d%s", m.matrix[i][j], delim)
	})
	result += "\n"
	return
}
