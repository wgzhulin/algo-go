package list

type Node struct {
	Val  int
	Next *Node
}

type DoublyNode struct {
	Val  int
	Pre  *DoublyNode
	Next *DoublyNode
}
