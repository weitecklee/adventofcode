#!/bin/bash

filename="input.txt"

if [ ! -f "$filename" ]; then
  echo "File '$filename' not found!"
  exit 1
fi

max=0
curr=0
elves=()

while IFS= read -r line; do

  if [ -z "$line" ]; then
    elves+=("$curr")
    if [ "$curr" -gt "$max" ]; then
      max="$curr"
    fi
    curr=0
  else
    curr=$((curr + line))
  fi

done < "$filename"

echo $max

IFS=$'\n' sorted=($(sort -nr <<<"${elves[*]}"))
unset IFS

part2=0
for ((i=0; i<3; i++)); do
    part2=$((part2 + sorted[i]))
done

echo $part2