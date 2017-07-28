package first

import "fmt"

func IntFloatTypes() {
	/*
		multi line comments
	*/
	var a int // declaring a variable, initialized to empty value
	a = 1
	a = 3
	b := 2 // declare and assign together

	//var a float64 = 2.0
	c := 3.0 // when youre defining a variable
	//a=4 // update with new value, but not :=
	//a++ // Should be in one line, not in the middle of an expression
	fmt.Printf("Hello %d %d %f", a, b, c)

	var result float32 = 5.0 / 2
	fmt.Printf("[-] Gloat %f", result)
	x := 2
	y := 5.0
	fmt.Printf("[-] Gloat2 %f", float64(x)/y)
	//fmt.Printf("%s", string(a))
}
