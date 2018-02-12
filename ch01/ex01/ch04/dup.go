package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]

	for _, arg := range files {
		countLines(arg, counts)
	}

	for line, files := range counts {
		if len(files) > 1 {
			fmt.Printf("%s\t: ", line)
			for filepath, _ := range files {
				fmt.Printf("%s\t", filepath)
			}
			fmt.Println()
		}
	}
}

func countLines(filepath string, counts map[string]map[string]int) {
	f, err := os.Open(filepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "dup: %v\n", err)
		f.Close()
		return
	}

	input := bufio.NewScanner(f)
	for input.Scan() {
		mm, nil := counts[input.Text()]
		if !nil {
			mm = make(map[string]int)
			counts[input.Text()] = mm
		}
		counts[input.Text()][filepath] = 1
	}

	f.Close()
}
