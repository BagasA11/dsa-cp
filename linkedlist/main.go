package main

import (
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

func main() {
	n1 := NewNode(1)
	n2 := NewNode(2)
	n3 := NewNode(3)
	n4 := NewNode(-1)
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n1.PrintNode()
	// n3.Destroy()
	Traverse(n1)
	fmt.Println("lowest: ", findLowestVal(n2))

	fmt.Println("")
	n1 = DeleteSpecificNode(n1, n3)
	Traverse(n1)
	// fmt.Println(n3)
	n5 := NewNode(5)
	fmt.Println("insert at: ", 4)
	n1 = InsertNode(n1, 4, n5)
	Traverse(n1)
	fmt.Println("mencari nilai x=3: ", search(n1, 3))

}

func NewNode(data int) *Node {
	node := new(Node)
	node.Data = data
	node.Next = nil
	return node
}

func (n *Node) PrintNode() {
	// fmt.Println(*n)
	nn := n
	for nn != nil {
		fmt.Println(nn)
		nn = nn.Next
	}
}

func Traverse(head *Node) {
	currNode := head

	for currNode != nil {

		fmt.Print(currNode, "->")
		currNode = currNode.Next
	}
	fmt.Println("")
}

func findLowestVal(head *Node) int {
	low := head.Data
	node := head
	for node != nil {
		if node.Data < low {
			low = node.Data
		}
		node = node.Next
	}
	return low
}

func DeleteSpecificNode(head *Node, target *Node) *Node {
	if head == target {
		return head.Next
	}
	currentNode := head
	for currentNode.Next != nil && currentNode.Next != target {
		currentNode = currentNode.Next
	}
	if currentNode.Next == nil {
		return head
	}
	currentNode.Next = currentNode.Next.Next
	return head
}

func InsertNode(head *Node, post uint, newnode *Node) *Node {
	// starting point at 1
	if post < 1 {
		post = 1
	}

	// if new node inserted to most left
	if post == 1 {
		newnode.Next = head
		return newnode // return new node as head
	}

	curentNode := head
	// traverse into position
	for i := uint(1); i < post; i++ {

		if curentNode.Next == nil {
			break
		}
		curentNode = curentNode.Next
	}

	newnode.Next = curentNode.Next
	curentNode.Next = newnode

	return head //return head, but with all new node

}

func search(head *Node, key int) bool {
	curentNode := head
	for curentNode != nil {
		if curentNode.Data == key {
			return true
		}
		curentNode = curentNode.Next
	}
	return false
}
