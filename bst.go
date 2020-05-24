// A simple implementation of a binary search tree in go
// This package is just a simple way to use binaries search trees. For the moment this package can only work with integers
package bst

// Node Structure
type Node struct {
	Left  *Node
	Val   int
	Right *Node
}

// Insert number into the Tree
func (t *Node) Insert(n int) error {
	// If the Node is Empty
	if t.Val == 0 {
		t.Val = n
	}
	// Left
	if t.Val > n {
		if t.Left == nil {
			t.Left = &Node{Val: n}
			return nil
		}

		return t.Left.Insert(n)
	}
	// Right
	if t.Val < n {
		if t.Right == nil {
			t.Right = &Node{Val: n}
			return nil
		}
		return t.Right.Insert(n)
	}
	return nil
}

// Get The Smallest Number
func (t *Node) GetMin() int {
	if t.Left == nil {
		return t.Val
	}
	return t.Left.GetMin()
}

// Get The Biggest Number
func (t *Node) GetMax() int {
	if t.Right == nil {

		return t.Val
	}

	return t.Right.GetMax()
}

// Get A Certain Node
func (t Node) GetNode(n int) int {
	if t.Val == n {
		return n
	}
	if n > t.Val {
		return t.Right.GetNode(n)
	}
	if n < t.Val {
		return t.Left.GetNode(n)
	}
	return t.GetNode(n)
}

// Returns an array with the sorted Tree
func (this *Node) Sort() []int {
	var res []int
	this.srt(&res)
	return res
}

func (this *Node) srt(res *[]int) {
	if this == nil {
		return
	}
	this.Left.srt(res)
	*res = append(*res, this.Val)
	this.Right.srt(res)
}

// Returns an array with the INVERSE sorted Tree
func (this *Node) SortInv() []int {
	var res []int
	this.srtInv(&res)
	return res
}

func (this *Node) srtInv(res *[]int) {
	if this == nil {
		return
	}
	this.Right.srtInv(res)
	*res = append(*res, this.Val)
	this.Left.srtInv(res)
}

// Check if a node exist. It returns a bool
func (this Node) Check(n int) bool {
	if this.Val == n {
		return true
	}
	if n > this.Val {
		if this.Right != nil {
			return this.Right.Check(n)
		}
	}
	if n < this.Val {
		if this.Left != nil {
			return this.Left.Check(n)
		}
	}
	return false
}

// Remove an element from the Tree
func (this *Node) Remove(n int) {
	var res []int
	sort := this.Sort()
	for _, i := range sort {
		if i != n {
			res = append(res, i)
		}
	}
	*this = Node{}
	for _, i := range res {
		this.Insert(i)
	}
}