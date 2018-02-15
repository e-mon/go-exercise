package main

import (
	"fmt"
	"strings"
	"time"
)

const (
	num int = 100000
)

func main() {

	// naive
	s := ""
	start := time.Now()
	for i := 0; i < num; i++ {
		s += "hoge"
	}
	end := time.Now()

	fmt.Println(fmt.Sprintf("naive : %0.5f sec", end.Sub(start).Seconds()))

	// strings.join
	ss := make([]string, num)

	start = time.Now()
	for i := 0; i < int(num); i++ {
		ss[i] = "hoge"
	}
	strings.Join(ss, "")
	end = time.Now()

	fmt.Println(fmt.Sprintf("strings.Join : %0.5f sec", end.Sub(start).Seconds()))
}
