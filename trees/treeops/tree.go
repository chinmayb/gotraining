package treeops

import (
	"fmt"
	"math/rand"
)

type Tree struct {
	Left  *Tree
	Right *Tree
	Data  int
}

func New(n int) *Tree {
	var t *Tree
	for i := 0; i < n; i++ {
		data := rand.Intn(50)
		fmt.Println(data)
		t = Insert(t, data)
	}
	return t

}

func Insert(t *Tree, data int) *Tree {
	if t == nil {
		return &Tree{Data: data}
	}
	if data < t.Data {
		t.Left = Insert(t.Left, data)
		return t
	}
	t.Right = Insert(t.Right, data)
	return t

}

func PreOrder(t *Tree) {
	if t == nil {
		return
	}
	PreOrder(t.Left)
	fmt.Println(t.Data)
	PreOrder(t.Right)
}
