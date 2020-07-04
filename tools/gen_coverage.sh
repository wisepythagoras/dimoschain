#!/bin/bash

current=`pwd`
output="$current/cover.out"

printf "mode: set\n" > $output

for dir in `find -name '*.go' -printf '%h\n' | sort -u`; do
    echo "Doing $dir"
    cd $dir
    go test -coverprofile /dev/stdout | awk 'NR>3 { print }' | head -n-1 >> $output
    cd $current
done

rm $output
go tool cover -html=$output -o coverage.html
