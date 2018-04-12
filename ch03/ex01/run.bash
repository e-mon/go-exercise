#!/bin/sh

echo "--- not skipping ---"
go run ./surface.go > output.html
nan=$(cat output.html | grep -c NaN)
echo "NaN : ${nan}";
inf=$(cat output.html | grep -c Inf)
echo "inf: ${inf}";

echo "--- skipping ---"
go run ./surface_skip_infnan.go > output.html
nan=$(cat output.html | grep -c NaN)
echo "NaN : ${nan}";
inf=$(cat output.html | grep -c Inf)
echo "inf: ${inf}";
