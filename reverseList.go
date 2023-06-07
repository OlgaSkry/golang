package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func main() {
	head := &Node{Value: 4}
	node1 := &Node{Value: 3}
	node2 := &Node{Value: 2}
	node3 := &Node{Value: 1}

	head.Next = node1
	node1.Next = node2
	node2.Next = node3

	printList(head)
	fmt.Println()
	printList(reverse(head))

}

func reverse(head *Node) *Node {
	current := head
	var previous *Node
	next := head.Next

	for current != nil {
		next = current.Next
		current.Next = previous
		previous = current
		current = next
	}
	return previous

}

func printList(l *Node) {
	for l != nil {
		fmt.Println(l.Value)
		l = l.Next
	}
}
