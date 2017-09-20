package ds

type Node struct {
	prev  *Node
	data string
	next  *Node
}

var Head *Node
var Tail *Node

func (node *Node) IsEmpty() bool {
	return  Head == nil && Tail == nil
}

func (node *Node) Add(data string) {
	temp := new(Node)
	temp.data = data

	if node.IsEmpty() {
		Head = temp
		Tail = temp
	} else {
		Tail.next = temp
		temp.prev = Tail
		Tail = temp
	}
}