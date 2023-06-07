package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func main() {
	head := &Node{Value: 1} //initialize the linked list with values
	node2 := &Node{Value: 2}
	node3 := &Node{Value: 3}
	node4 := &Node{Value: 4}

	head.Next = node2 //point to the elements using linked list
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2

	var has_circle bool
	var loopNode *Node

	fmt.Println("Does this linked list has loop?")
	has_circle, loopNode = has_loop(head)
	fmt.Println(has_circle) // Output: true

	fmt.Println("Where it is?")
	fmt.Println(getLoopNode(head, loopNode)) // Output: true
}

func has_loop(head *Node) (bool, *Node) {
	high := head
	low := head

	for high != nil && high.Next != nil {
		low = low.Next
		high = high.Next.Next
		if low == high {
			return true, low
		}
	}
	return false, nil
}

func getLoopNode(head *Node, meetingPoint *Node) *Node {
	high := meetingPoint
	low := head

	for high != nil && high.Next != nil {
		low = low.Next
		high = high.Next
		if low == high {
			return low
		}
	}
	return nil
}
