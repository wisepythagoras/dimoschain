#!/bin/bash

total=0

while read file; do
    res=`wc -l $file;`
    echo $res

    lines=$(echo $res | awk '{ print $1 }')
    lines=$(($lines + $total))
    total=$lines
done < <(find . -name "*.go")

echo -e "\n$lines total lines of code"
