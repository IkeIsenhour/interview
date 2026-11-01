package listnode

type ListNode struct {
	Value int
	Next  *ListNode
}

func NewListNode(v int, n *ListNode) *ListNode {
	return &ListNode{
		Value: v,
		Next:  n,
	}
}
