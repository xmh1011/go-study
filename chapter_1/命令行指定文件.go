package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int) //make(map[string]int) is a map literal
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg) // Open opens the named file for reading.
			if err != nil {
				fmt.Fprintf(os.Stderr, "demo2: %v\n", err)
			}
			countLines(f, counts) //countLines reads from stdin and counts the lines
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Println(line)
		}
	}
}

func countLines(stdin *os.File, counts map[string]int) {
	input := bufio.NewScanner(stdin) //NewScanner returns a new Scanner to read from r.
	for input.Scan() {
		counts[input.Text()]++
	}
}
