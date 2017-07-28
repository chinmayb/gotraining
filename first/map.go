//pointers slices and maps are not allocated, they point at allocated memory
package first

import "fmt"

var v1 map[int] string

func MadMaps() {
	v2 := make(map[int]string)
	v2[1] = "ásdAAÄÄ"

	fmt.Print(v2)

	v3 := map[string]int {
		"k1": 21,
	}
}

//write me a lib to matrices
//create, takes dimension, 
// slice of slices. slices of ints
// add , takees two matrices
//