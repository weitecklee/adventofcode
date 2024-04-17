#!/bin/bash

part1() {
  local pos=0
  local depth=0

  while IFS=' ' read -r action val || [ -n "$action" ]; do
    if [ "$action" == "forward" ]; then
      (( pos += val ))
    elif [ "$action" == "down" ]; then
      (( depth += val ))
    else
      (( depth -= val ))
    fi
  done < "$1"

  echo $(( pos * depth ))
}

part2() {
  local pos=0
  local depth=0
  local aim=0

  while IFS=' ' read -r action val || [ -n "$action" ]; do
    if [ "$action" == "forward" ]; then
      (( pos += val ))
      (( depth += aim * val ))
    elif [ "$action" == "down" ]; then
      (( aim += val ))
    else
      (( aim -= val ))
    fi
  done < "$1"

  echo $(( pos * depth ))
}

main() {
  puzzle_input="input.txt"

  if [ -f "$puzzle_input" ]; then
    echo "$(part1 "$puzzle_input")"
    echo "$(part2 "$puzzle_input")"
  else
    echo "Input file not found: $puzzle_input"
    exit 1
  fi
}

main
