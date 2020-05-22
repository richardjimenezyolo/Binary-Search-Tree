# Binary Search Tree Package for the Golang programing language
A simple implementation of a binary search tree in go

This package is just a simple way to use binaries search tries.

## Create a Tree
```golang
var Tree bst.Node;
```

## Insert a number into the Tree
```golang
Tree.Insert(2)
```

## Get the smallest number
```golang
SmallestNumber := Tree.GetMin()
```

## Get the biggest number
```golang
BigNumber := Tree.GetMax()
```

## Get the tree sorted
```golang
sort := Tree.Sort() // Return an array
```
