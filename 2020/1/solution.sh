#!/bin/bash

part1() {
  local numbers=("$@")
  local i=0
  local j=$(( ${#numbers[@]} - 1 ))

  local sum=$(( ${numbers[i]} + ${numbers[j]} ))
  while [ "$i" -lt "$j" ] && [ "${sum}" != "${target}" ]; do
    if [ "${sum}" -gt "${target}" ]; then
      ((j--))
    else
      ((i++))
    fi
    sum=$(( ${numbers[i]} + ${numbers[j]} ))
  done
  if [ "$i" -ge "$j" ]; then
    return
  fi
  echo $(( ${numbers[i]} * ${numbers[j]} ))
}

part2() {
  local numbers=("$@")
  for ((i=0; i<${#numbers[@]}-2; i++)); do
    local j=$(( ${i} + 1 ))
    local k=$(( ${#numbers[@]} - 1 ))
    local sum=$(( ${numbers[i]} + ${numbers[j]} + ${numbers[k]} ))
    while [ "$j" -lt "$k" ] && [ "${sum}" != "${target}" ]; do
      if [ "${sum}" -gt "${target}" ]; then
        ((k--))
      else
        ((j++))
      fi
      sum=$(( ${numbers[i]} + ${numbers[j]} + ${numbers[k]} ))
    done
    if [ "${sum}" == "${target}" ]; then
      echo $(( ${numbers[i]} * ${numbers[j]} * ${numbers[k]} ))
      return
    fi
  done
}

main() {
  local filename="input.txt"
  local target=2020

  if [ -f "$filename" ]; then
    local sorted=($(sort -n < "$filename"))
    echo "$(part1 "${sorted[@]}")"
    echo "$(part2 "${sorted[@]}")"
  else
    echo "Input file not found: $filename"
    exit 1
  fi
}

main
