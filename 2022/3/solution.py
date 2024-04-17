import os
from typing import Dict, List, Set

ALPHABET = '_abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ'
PRIORITIES: Dict[str, int] = {c: i for i, c in enumerate(ALPHABET)}

def calculate_priorities(*args):
  return sum([PRIORITIES[c] for c in set.intersection(*args)])

def part1(puzzle_input: List[str]) -> int:
  res = 0
  for line in puzzle_input:
    n = len(line)
    comp1: Set[str] = set(line[:n//2])
    comp2: Set[str] = set(line[n//2:])
    res += calculate_priorities(comp1, comp2)
  return res

def part2(puzzle_input: List[str]) -> int:
  res = 0
  for i in range(0, len(puzzle_input), 3):
    ruck1: Set[str] = set(puzzle_input[i])
    ruck2: Set[str] = set(puzzle_input[i + 1])
    ruck3: Set[str] = set(puzzle_input[i + 2])
    res += calculate_priorities(ruck1, ruck2, ruck3)
  return res

if __name__ == '__main__':
  with open(os.path.join(os.path.dirname(os.path.abspath(__file__)), 'input.txt'),'r') as file:
    puzzle_input = file.read().strip().split('\n')
  print(part1(puzzle_input))
  print(part2(puzzle_input))