package stack_linklist

// program implementasi stack dengan linkedlist
// terdapat 2 struct
// - linkedlist sebagai private class
// - stack sebagai public class

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

// konstruktor node
func newNode(data int) *node {
	n := new(node)
	n.data = data
	n.next = nil
	return n
}

func (n *node) printList() {
	cn := n

	for cn != nil {
		fmt.Println(cn, "->")
		cn = cn.next
	}
}

func (n *node) insert(newnode *node, post uint) *node {
	if post == 0 {
		post++
	}
	if post == 1 {
		newnode.next = n
		return newnode
	}
	cn := n //create temporary node and pointing to head node
	for i := uint(2); i < post; i++ {
		if cn.next == nil {
			break
		}
		cn = cn.next
	}
	// pointing next node of newnode as nrext of current node
	newnode.next = cn.next
	cn.next = newnode //pointing next of current node to newnode
	return n
}

func (n *node) delete(target *node) *node {
	if n == target {
		return n.next
	}
	cn := n //current node as temporary
	for cn.next != nil && cn.next != target {
		cn = cn.next
	}
	// pointing next current node to next of next current node
	// next node of current node is node target to delete
	cn.next = cn.next.next
	return n
}

func (n *node) getNode(index uint) *node {
	if index == 0 {
		index++
	}
	if index == 1 {
		return n
	}
	temp := n
	i := uint(1)
	for temp.next != nil {
		if i == index {
			break
		}
		temp = temp.next
		i++
	}
	return temp
}

func (n *node) size() uint {
	i := uint(0)
	temp := n
	for temp != nil {

		temp = temp.next
		i++
	}
	return i
}

type Stack struct {
	Data *node
	size uint
}

func NewStack(data int) *Stack {
	node := newNode(data)
	size := node.size()
	return &Stack{
		Data: node,
		size: size,
	}
}

func (s *Stack) Push(data int) {
	newnode := newNode(data)
	s.Data = s.Data.insert(newnode, s.size+1)
	s.size = s.Data.size()
}

func (s *Stack) GetSize() uint {
	return s.size
}

func (s *Stack) GetLayer() int {
	data := s.Data.getNode(s.size)
	return data.data
}

func (s *Stack) Pop() int {
	pop := s.GetLayer()
	size := s.size
	target := s.Data.getNode(size) //create delete target node
	s.Data = s.Data.delete(target) //reset node after deletion
	s.size = s.Data.size()         //reset size value after deletion
	return pop
}
