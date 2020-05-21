package bst

import (
	"fmt"
)

// Node Structure

type Node struct {
	Left  *Node
	val   int
	Right *Node
}

// Insert number into the Tree
func (t *Node) Insert(v int) error {
	// If the Node is Empty
	if t.val == 0 {
		t.val = v
	}
	// Left
	if t.val > v {
		if t.Left == nil {
			t.Left = &Node{val: v}
			return nil
		}

		return t.Left.Insert(v)
	}
	// Right
	if t.val < v {
		if t.Right == nil {
			t.Right = &Node{val: v}
			return nil
		}
		return t.Right.Insert(v)
	}
	return nil
}

// Get The Smallest Number
func (t *Node) GetMin() int {
	if t.Left == nil {
		return t.val
	}
	return t.Left.GetMin()
}

// Get The Biggest Number
func (t *Node) GetMax() int {
	if t.Right == nil {
		return t.val
	}

	return t.Right.GetMax()
}

// Get A Certain Node
func (t Node) GetNode(n int) int {
	if t.val == n {
		return n
	}
	if n > t.val {
		fmt.Println("Right")
		return t.Right.GetNode(n)
	}
	if n < t.val {
		fmt.Println("Left")
		return t.Left.GetNode(n)
	}
	return t.GetNode(n)
}

// Order The Tree
var res []int

func OrderFunc(this *Node) {
	if this == nil {
		return
	}
	OrderFunc(this.Left)
	res = append(res, this.val)
	OrderFunc(this.Right)
}

func (this *Node) Sort() []int {
	OrderFunc(this)
	return res
}
