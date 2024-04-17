#!/bin/bash

filename="input.txt"

if [ ! -f "$filename" ]; then
  echo "File '$filename' not found!"
  exit 1
fi

part1() {
  local numbers=("$@")
  local count=0
  for ((i=0; i<${#numbers[@]}-1; i++)); do
    if [ "${numbers[i+1]}" -gt "${numbers[i]}" ]; then
      ((count++))
    fi
  done
  echo "$count"
}

part2() {
  local numbers=("$@")
  local sums=()
  for ((i=2; i<${#numbers[@]}; i++)); do
    sums+=("$(( ${numbers[i-2]} + ${numbers[i-1]} + ${numbers[i]} ))")
  done
  echo "$(part1 "${sums[@]}")"
}

main() {
  local puzzle_input=()
  while IFS= read -r num || [ -n "$num" ]; do
    puzzle_input+=("$num")
  done < "$filename"

  echo "$(part1 "${puzzle_input[@]}")"
  echo "$(part2 "${puzzle_input[@]}")"
}

main
