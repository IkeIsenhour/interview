package problem

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	visitedHash := map[*Node]*Node{}
	var clonedNode func(node *Node) *Node

	clonedNode = func(node *Node) *Node {
		if node == nil {
			return nil
		}

		if _, ok := visitedHash[node]; ok {
			return visitedHash[node]
		}

		newNode := &Node{
			Val: node.Val,
		}
		visitedHash[node] = newNode

		newNode.Next = clonedNode(node.Next)
		newNode.Random = clonedNode(node.Random)
		return newNode
	}

	return clonedNode(head)
}
