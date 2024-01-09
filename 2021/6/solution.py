import os
from typing import List

def simulate(puzzle_input: List[int], days: int) -> int:
  fish = [0] * 9
  for i in puzzle_input:
    fish[i] += 1
  for _ in range(days):
    fish2 = [0] * 9
    for i in range(8):
      fish2[i] = fish[i + 1]
    fish2[6] += fish[0]
    fish2[8] = fish[0]
    fish = fish2
  return sum(fish)

if __name__ == "__main__":
  with open(os.path.join(os.path.dirname(os.path.abspath(__file__)), 'input.txt'),'r') as file:
    puzzle_input = [int(n) for n in file.readline().split(',')]
  print(simulate(puzzle_input, 80))
  print(simulate(puzzle_input, 256))