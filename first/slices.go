package first

import "fmt"

func SliceOperations(){
	arr1 := []int{1,2,2,2} // [...]int{1,2,,} will be an array
	//arr2 := make([]int, 10) // slice of capaciy 10
	// arr2 is a pointer basically
	// you can declare slice from an array
	arr1[0] = 2
	arr1 = append(arr1, 3) //variadic funcs

	arro := [...]int{12,123,1234}
	arr3 := arro[0:3]
	fmt.Printf("%d", arr3)
	//arr1 = arr1[2:3]
	//len() , cap()
	arr1 = arr1[3:]
	for i,v := range arr1 {
		fmt.Printf("\n%d --- %d", i, v)
	}

	test()
}

// slice always depends on the array, it points to that.
func test() {
	var ar1 = [...]string{"A", "B", "C", "D"}
	var sl1 []string
	sl1 = ar1[2:3] //slicing an array returns a slice
	fmt.Printf("'%s ", sl1)
	sl1[0] = "E"
	fmt.Printf("'%s ", sl1)

	_ = append(sl1, "Z")
	fmt.Printf("slice %s  arr : %s", sl1, ar1)
	sl1 = append(sl1, "Y")

	fmt.Printf("slice %s  arr : %s", sl1, ar1) //ar1 remains the same, sl1 expands

}
