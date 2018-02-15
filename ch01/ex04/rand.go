package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 32; i++ {
		fmt.Printf("%v\n", rand.Intn(128))
	}
}
