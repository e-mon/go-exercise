#!/bin/bash

go run fetchall.go http://www.worldslongestwebsite.com/
go run fetchall.go http://www.worldslongestwebsite.com/

echo "------ results ------"
cat ./result*

