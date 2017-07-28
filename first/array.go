package first
import "fmt"
var arr1 [5]int //array of ints

func sumarray(arr2 [5]int) int{
	var sum int
	for _,v := range arr2 {
		sum += v
	}
	return sum
}

func Arrayoperations(){
	arr2 := [5]int {1,2,2,2}
	arr1[2] = 2
	arr1[3] = 3
	for i,v := range arr1 {
		fmt.Printf("\n%d --- %d", i, v)
	}
	fmt.Println("sumarray", sumarray(arr2))
}