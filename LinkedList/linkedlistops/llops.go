package linkedlistops

import "fmt"

// type LinkedListOperators interface {
// 	Remove()
// 	Append()
// 	Prepend()
// 	Show()
// }

type Node struct {
	Data interface{}
	link *Node
}

// func (n *Node) Next() *Node {
// 	return n.Link
// }

type LinkedList struct {
	head *Node
}

func ShowNode(n *Node) (result string) {
	temp := n
	for temp != nil {
		result += fmt.Sprintf("%v -> ", temp.Data)
		temp = temp.link
	}
	return
}

func (l LinkedList) Show() string {
	if l.head == nil {
		return ""
	}
	result := ""
	temp := l.head
	for temp.link != nil {
		result += fmt.Sprintf("%v ->", temp.Data)
		temp = temp.link
	}
	return result
}

func (l *LinkedList) Append(node *Node) {
	if l.head == nil {
		l.head = node
	} else {
		temp := l.head
		for temp.link != nil {
			temp = temp.link
		}
		temp.link = node
	}
}

func NewNode(data interface{}, l *Node) *Node {
	return &Node{Data: data, link: l}
}
