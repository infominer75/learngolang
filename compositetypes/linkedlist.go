package main

import "fmt"

type listnode struct {
	v int
	next *listnode
}

func main() {
	values := []int{1, 3, 5, 6, 4}
	var current *listnode
	var head *listnode

	for _, v := range values {
		current = Add(current, v)
		if head == nil {
			head = current
		}
	}

	PrintList(head)
}

func Add(current *listnode, value int) *listnode {
	if current ==  nil {
		current = new (listnode)
		current.v = value
		return current
	}
	current.next = Add(current.next, value)
	return current.next
}

func PrintList(start *listnode) {
	for start != nil {
		fmt.Println(start.v)
		start = start.next
	}
}
