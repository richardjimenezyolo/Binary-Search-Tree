# Binary Search Tree Package for the Golang programing language
A simple implementation of a binary search tree in go

This package is just a simple way to use binaries search trees. For the moment this package can only work with integers 

## Create a Tree
```golang
var Tree bst.Node;
```

## Insert a number into the Tree
```golang
Tree.Insert(2) // Instead of 2 you can use any number you want
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
