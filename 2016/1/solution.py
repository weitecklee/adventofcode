import os
from typing import List, Set, Tuple

def parse(puzzle_input: str) -> List[Tuple[str, int]]:
  instructions: List[Tuple[str, int]] = []
  for line in puzzle_input.split(', '):
    instructions.append((line[0], int(line[1:])))
  return instructions

def part1(instructions: List[Tuple[str, int]]) -> int:
  pos = [0, 0]
  direc = [0, 1]
  for instruction in instructions:
    if instruction[0] == 'L':
      direc[0], direc[1] = -direc[1], direc[0]
    else:
      direc[0], direc[1] = direc[1], -direc[0]
    pos[0] += instruction[1] * direc[0]
    pos[1] += instruction[1] * direc[1]
  return abs(pos[0]) + abs(pos[1])

def part2(instructions: List[Tuple[str, int]]) -> int:
  pos = [0, 0]
  direc = [0, 1]
  visited: Set[Tuple[int, ...]] = set()
  for instruction in instructions:
    if instruction[0] == 'L':
      direc[0], direc[1] = -direc[1], direc[0]
    else:
      direc[0], direc[1] = direc[1], -direc[0]
    for _ in range(instruction[1]):
      pos[0] += direc[0]
      pos[1] += direc[1]
      if tuple(pos) in visited:
        return abs(pos[0]) + abs(pos[1])
      visited.add(tuple(pos))
  return -1

if __name__ == '__main__':
  with open(os.path.join(os.path.dirname(os.path.abspath(__file__)), 'input.txt'),'r') as file:
    puzzle_input = file.readline().strip()
  instructions = parse(puzzle_input)
  print(part1(instructions))
  print(part2(instructions))

