package main

import "fmt"

/*
def cons(a, b):
	def pair(f):
		return f(a, b)
return pair
*/

func cons(a func(int) int, b int) func(i int) bool{
	return func(i int) bool{
		if a(i) > b {
			if i >= 0 {
				return true
			}
			return false
		}
		return true
	}
}

func main(){
	a := func(i int) int{
		return i
	}
	f := cons(a ,2)
	fmt.Println(f(1))
}
