package main

import (
	"os"

	"github.com/IkeIsenhour/leetcode-go/internal/structure/linkedlist"
	"github.com/IkeIsenhour/leetcode-go/internal/structure/listnode"
)

func main() {
	c := listnode.NewListNode(3, nil)
	b := listnode.NewListNode(2, c)
	a := listnode.NewListNode(1, b)
	ll := linkedlist.NewLinkedList(a)

	ll.Print(os.Stdout)
}
