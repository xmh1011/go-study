package main

import (
	"fmt"
	"os"
	"strings"
)

// write a function that reads from specified file and prints the results to stdout
func main() {
	counts := make(map[string]int)
	for _, filenames := range os.Args[1:] {
		data, err := os.ReadFile(filenames) // ReadFile reads the named file and returns the contents.
		if err != nil {
			fmt.Fprintf(os.Stderr, "demo3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
		for line, n := range counts {
			if n > 1 {
				fmt.Println(line)
			}
		}
	}
}
