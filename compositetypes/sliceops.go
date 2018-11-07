package main

import (
	"fmt"
)

func main() {
	numbers := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//the call "reverse(numbers)" is a compilation error - since what is expected is a slice.
	reverse(numbers[:]) //pass the whole underlying array as a slice to the reverse function
	fmt.Println("Reversed numbers : ", numbers)
}

//the below function takes a slice as a parameter
//this means that the slices are always passed by references
//hence all operations happening within the below function on
//would actually be modifying the underlying array.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		temp := s[i]
		s[i] = s[j]
		s[j] = temp
	}
}
