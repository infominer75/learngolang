package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var s, sep string
	fmt.Printf("The number of arguments on the command line : %d\n", len(os.Args))
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	//printing the contents of the string as a range or slice iteration
	for i, s := range os.Args[1:] {
		fmt.Printf("Index : %d, Value :%s%s", i, s, "\n")
	}

	//the same thing using Join from the strings package
	fmt.Println("Using strings.Join")
	//assign the value to a variable - we would want to print the string
	//using an iteration over it's individual characters
	jointString := strings.Join(os.Args[1:], " ")
	fmt.Println(jointString)
	//loop through the contents of the joint string index by index
	for i, c := range jointString {
		fmt.Printf("Index : %d, Character : %c\n", i, c)
	}

	//read the string word by word
	input := strings.Split(jointString, " ")
	fmt.Println("Length of split string components ", len(input))
	fmt.Printf("The entire string array : %v", input)
}
