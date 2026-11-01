package linkedlist

import (
	"fmt"
	"io"

	"github.com/IkeIsenhour/leetcode-go/internal/structure/listnode"
)

type LinkedList struct {
	Head *listnode.ListNode
}

func NewLinkedList(head *listnode.ListNode) *LinkedList {
	return &LinkedList{
		Head: head,
	}
}

func (l *LinkedList) Print(w io.Writer) {
	head := l.Head

	for head != nil {
		fmt.Fprint(w, head.Value, "-->")
		head = head.Next
	}
}
func (l *LinkedList) Search(target int) *listnode.ListNode {
	head := l.Head

	for head != nil {
		if head.Value == target {
			return head
		}
		head = head.Next
	}

	return nil
}

func (l *LinkedList) Insert(val int, position int) *listnode.ListNode {
	if position == 1 {
		newNode := listnode.NewListNode(val, l.Head)
		l.Head = newNode
		return newNode
	}

	temp := l.Head
	for i := 1; i < position-1; i++ {
		if temp.Next == nil {
			newNode := listnode.NewListNode(val, nil)
			temp.Next = newNode
			return newNode
		}
		temp = temp.Next
	}

	newNode := listnode.NewListNode(val, temp.Next)
	temp.Next = newNode

	return newNode
}
