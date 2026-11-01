package linkedlist

import (
	"bytes"
	"testing"

	"github.com/IkeIsenhour/leetcode-go/internal/structure/listnode"
	utilityGoTest "github.com/IkeIsenhour/utility-go/pkg/test"
)

func TestPrint(t *testing.T) {
	c := listnode.NewListNode(3, nil)
	b := listnode.NewListNode(2, c)
	a := listnode.NewListNode(1, b)
	ll := NewLinkedList(a)

	t.Run("print output", func(t *testing.T) {
		var output bytes.Buffer

		ll.Print(&output)
		got := output.String()
		want := "1-->2-->3-->"

		utilityGoTest.AssertEquality(t, got, want)
	})
}

func TestSearch(t *testing.T) {
	c := listnode.NewListNode(3, nil)
	b := listnode.NewListNode(2, c)
	a := listnode.NewListNode(1, b)
	ll := NewLinkedList(a)

	t.Run("search success", func(t *testing.T) {
		node := ll.Search(3)

		got := node.Value
		want := 3

		utilityGoTest.AssertEquality(t, got, want)
	})
}

func TestInsert(t *testing.T) {
	d := listnode.NewListNode(4, nil)
	c := listnode.NewListNode(3, d)
	b := listnode.NewListNode(2, c)
	a := listnode.NewListNode(1, b)
	ll := NewLinkedList(a)

	t.Run("insert middle success", func(t *testing.T) {
		ll.Insert(7, 3)

		var output bytes.Buffer

		ll.Print(&output)
		got := output.String()
		want := "1-->2-->7-->3-->4-->"

		utilityGoTest.AssertEquality(t, got, want)
	})

	t.Run("insert first success", func(t *testing.T) {
		ll.Insert(10, 1)

		var output bytes.Buffer

		ll.Print(&output)
		got := output.String()
		want := "10-->1-->2-->7-->3-->4-->"

		utilityGoTest.AssertEquality(t, got, want)
	})

	t.Run("insert past length success", func(t *testing.T) {
		ll.Insert(13, 27)

		var output bytes.Buffer

		ll.Print(&output)
		got := output.String()
		want := "10-->1-->2-->7-->3-->4-->13-->"

		utilityGoTest.AssertEquality(t, got, want)
	})
}
