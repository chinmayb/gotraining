package main

import "fmt"

type stack struct {
	elem []int
}

func newStack() *stack {
	return &stack{}
}

func (s *stack) push(elem int) {
	s.elem = append(s.elem, elem)
}

func (s *stack) pop() int {
	l := len(s.elem)
	temp := s.elem[l-1]
	s.elem = s.elem[:l-1]
	return temp
}

func (s *stack) print() {
	fmt.Println("")
	for i := len(s.elem)-1; i >=0 ; i-- {
		fmt.Printf("|__%d__|\n", s.elem[i])
	}
}

func main() {
	s1 := newStack()
	s1.push(1)
	s1.push(2)
	s1.push(3)
	s1.push(4)
	s1.push(5)

	s2 := newStack()
	s2.push(s1.pop())
	s2.push(s1.pop())
	s2.push(s1.pop())
	s2.push(s1.pop())
	s2.push(s1.pop())
	s2.print()

	s2.pop(
	)
	s2.print()
}
