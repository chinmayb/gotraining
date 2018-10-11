package main


type Node struct {
	data int
	lLink, rLink *Node
}

type DLinkedList struct {
	head *Node
}

func NewDLinkedList() *DLinkedList{
	return &DLinkedList{&Node{}}
}

func newNode(data int, llink, rlink *Node) *Node{
	return &Node{data, llink, rlink}
}

func (d *DLinkedList) Add(data int) {
	cur := d.head
	if (cur.data == 0){
		cur.data = data
		cur.lLink = nil
		cur.rLink = nil
		return
	}
	if cur.rLink == nil && cur.lLink == nil {
		cur.rLink =  newNode(data, cur, nil)
		return
	}
	for cur.rLink != nil  {
		cur = cur.rLink
	}
	cur.rLink = newNode(data, cur, nil)
}

func (d *DLinkedList) Delete(data int){

}