package main

import "fmt"

type node struct {
	value int
	left, right *node
}

func main() {
	values := []int { 5, 3, 4, 1, 2, 6, 7}
	Sort(values)
}
func Sort(values []int) {
	var root *node
	for _, v := range  values {
		root = add(root, v)
	}
	values = appendValues(values[:0], root)
	for _, v := range values {
		fmt.Println(v)
	}
}

func add(n *node, v int) *node {

	if n == nil {
		n = new(node)
		n.value = v
		return n
	}
	if v < n.value {
		n.left = add(n.left, v)
	} else {
		n.right = add(n.right, v)
	}

	return n
}

func appendValues(values []int , root *node) []int  {
	if root != nil {
		values = appendValues(values, root.left)
		values = append(values, root.value)
		values = appendValues(values, root.right)
	}
	return values
}