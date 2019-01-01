package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wordfreqmap := make(map[string]int)
	if len(os.Args) < 2 {
		fmt.Println("Usage: wordfreq <path to input>")
		os.Exit(1)
	}
	f, err := os.Open(os.Args[1])
	if (err != nil) {
		fmt.Fprintf(os.Stdout, "Error when reading from file :%v\n", err)
		os.Exit(2)
	}
	input := bufio.NewScanner(f)
	//breaks the input into a series of words. To break the input by line use
	//bufio.ScanLines
	input.Split(bufio.ScanWords)

	for input.Scan() {
		word := input.Text()
		wordfreqmap[word] = wordfreqmap[word] + 1
	}
	for word, count := range wordfreqmap {
		fmt.Fprintf(os.Stdout, "%s appears %d times\n", word, count)
	}
	fmt.Println("Done")
}
