package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, file := range files {
			//read file contents. Open the file and then pass
			//the file pointer to the counter
			f, err := os.Open(file)
			if err != nil {
				fmt.Println("Cannot open file ", file, " continuing on to the next...", err)
				continue
			}
			countLines(f, counts)
			//explicitly calling close. a better way would have been to defer the closure so
			//that this behavior is similar to finally block in Java
			f.Close()
		}
	}
	//now that the counts are available, loop through the map and print the counts of the lines
	for line, n := range counts {
		fmt.Printf("Line : %s, Count :%d\n", line, n)
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
