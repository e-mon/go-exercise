package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	sep = "\n"
	for i := 0; i < len(os.Args); i++ {
		s += fmt.Sprintf("%d : %s%s", i, os.Args[i], sep)
	}
	fmt.Println(s)
}
