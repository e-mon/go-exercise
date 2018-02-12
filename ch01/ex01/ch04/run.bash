#!/bin/sh

for i in {0..4}; do
	go run ./rand.go > $i.txt
done

go run ./dup.go *.txt
