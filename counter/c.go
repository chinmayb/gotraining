package counter

import "fmt"

var cnt int = 0 //it has to be var, not :=

var COUNT int //exported globally

func Init() {
	cnt = 1
}

func Next() (int) {
	result := cnt
	cnt ++ 
	return result
}

func Show() {
	fmt.Print("\n" , COUNT)
}