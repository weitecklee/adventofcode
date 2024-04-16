import os
from typing import List

def parse(puzzle_input: List[str]) -> List[int]:
  elves: List[int] = []
  curr = 0
  for line in puzzle_input:
    if line:
      curr += int(line)
    else:
      elves.append(curr)
      curr = 0
  elves.append(curr)
  return elves

if __name__ == '__main__':
  with open(os.path.join(os.path.dirname(os.path.abspath(__file__)), 'input.txt'),'r') as file:
    puzzle_input = [line.strip() for line in file]
  elves = parse(puzzle_input)
  elves.sort()
  print(elves[-1])
  print(sum(elves[-3:]))