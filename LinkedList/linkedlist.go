package main

import ops "./linkedlistops"
import "fmt"

func main() {
	//ll := ops.LinkedList{}
	data1 := "a"
	data2 := "b"
	data3 := 1
	node1 := ops.NewNode(data1, nil)
	node2 := ops.NewNode(data2, node1)
	node3 := ops.NewNode(data3, node2)
	fmt.Println(ops.ShowNode(node3))
	//ll.Append(node1)
	//fmt.Printf("%v", ll)
}
