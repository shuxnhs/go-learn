package DataStructure

// 链表结点定义
type Node struct {
	Data interface{}
	Next *Node
}

func NewLinkedNode() *Node {
	return &Node{}
}
