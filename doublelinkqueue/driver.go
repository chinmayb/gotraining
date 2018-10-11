package main

import "fmt"


func (d *DLinkedList) Print() {
	cur := d.head
	for cur != nil  {
		fmt.Printf("<-> %d" , cur.data)
		cur = cur.rLink
	}
}

func main() {
	dList := NewDLinkedList()
	dList.Add(1)
	dList.Add(2)
	dList.Print()
}
