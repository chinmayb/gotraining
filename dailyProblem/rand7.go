package main

import (
	"fmt"
	"math/rand"
)

/*
Using a function rand7() that returns an integer from 1 to 7 (inclusive) with uniform probability

1/8
*/

func main() {
	for i := 1; i <= 7; i++ {
		fmt.Print(rand.Intn(7))
	}

}
