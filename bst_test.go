package bst

import (
	"testing"
)

var Tree Node

func TestSetup(t *testing.T) {

	lts := []int{5, 4, 2, 5, 7, 3, 7, 1, 10, 11}
	for _, i := range lts {
		Tree.Insert(i)
	}

}

func TestSort(t *testing.T) {

	sort := Tree.Sort()
	expected := []int{1, 2, 3, 4, 5, 7, 10, 11}

	if len(sort) != len(expected) {
		t.Errorf("SORT DOES NOT RETURNED THE SAME NUMBER OF ELEMENTS")
	}

	for index, ele := range expected {
		if sort[index] != ele {
			t.Errorf("ELEMENTS ARE NOT EQUAL")
		}
	}
}

func TestSortInv(t *testing.T) {
	sort := Tree.SortInv()
	expected := []int{11, 10, 7, 5, 4, 3, 2, 1}

	if len(sort) != len(expected) {
		t.Errorf("SORT INV DOES NOT RETURNED THE SAME NUMBER OF ELEMENTS")
	}

	for index, ele := range expected {
		if sort[index] != ele {
			t.Errorf("ELEMENTS ARE NOT EQUAL")
		}
	}
}

func TestCheck(t *testing.T) {
	sort := Tree.Sort()
	if !Tree.Check(sort[0]) {
		t.Errorf("Error at checking")
	}
	if Tree.Check(100) {
		t.Errorf("Error at checking")
	}
}

func TestGetMin(t *testing.T) {
	sort := Tree.Sort()
	if Tree.GetMin() != sort[0] {
		t.Errorf("Error at GetMin")
	}
}

func TestGetMax(t *testing.T) {
	if Tree.GetMax() != 11 {
		t.Errorf("Error at GetMax")
	}
}

// Remove Test
func TestRemove(t *testing.T) {
	sort_before := Tree.Sort()
	Tree.Remove(11)
	sort_after := Tree.Sort()

	if len(sort_before) == len(sort_after) {
		t.Errorf("ERROR AT Remove, BOTH LIST HAVE THE SAME AMOUNT OF ELEMENTS")
	}

	for _, i := range sort_after {
		if i == 11 {
			t.Errorf("ELEMENTS STILL IN THE Tree AFTER Remove")
		}
	}
}
