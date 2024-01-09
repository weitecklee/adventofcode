import os
from typing import Callable, List, Tuple

def solve(crabs: List[int], calculate: Callable[[List[int], int], int]) -> int:
  candidates: List[Tuple[int, int]] = []
  left = min(crabs)
  right = max(crabs)
  candidates.append((calculate(crabs, left), left))
  candidates.append((calculate(crabs, right), right))
  mid = left + (right - left) // 2
  candidates.append((calculate(crabs, mid), mid))
  candidates.sort()
  while abs(candidates[0][1] - candidates[1][1]) > 1:
    mid = (candidates[0][1] + candidates[1][1]) // 2
    candidates.append((calculate(crabs, mid), mid))
    candidates.sort()
  return candidates[0][0]

def calculate_fuel1(crabs: List[int], target: int) -> int:
  return sum(abs(crab - target) for crab in crabs)

def calculate_fuel2(crabs: List[int], target: int) -> int:
  return sum(abs(crab - target) * (abs(crab - target) + 1) // 2 for crab in crabs)

if __name__ == "__main__":
  with open(os.path.join(os.path.dirname(os.path.abspath(__file__)), 'input.txt'),'r') as file:
    puzzle_input = [int(n) for n in file.readline().split(',')]
  print(solve(puzzle_input, calculate_fuel1))
  print(solve(puzzle_input, calculate_fuel2))