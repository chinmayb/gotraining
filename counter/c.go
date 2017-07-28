package counter

import "fmt"

var cnt int = 0 //it has to be var, not :=

var COUNT int //exported globally

func Init() {
	fmt.Println("\nResetting")
	cnt = 0
}


// func Next() (int) {
// 	result := cnt
// 	cnt ++ 
// 	return result
// }


//alternatively using defer
func Next() int {
	defer increment()
	//defer increment()
	return cnt
}

func Show() {
	fmt.Print("\n" , COUNT)
}

func increment() {
	cnt++
}
