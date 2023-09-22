from typing import List, Tuple
from functools import reduce

def checker(lines: List[str], gradient: Tuple[int, int]) -> int:
  x, y = 0, 0
  count = 0
  h, w = len(lines), len(lines[0])
  while y < h:
    if lines[y][x] == "#":
      count += 1
    x = (x + gradient[0]) % w
    y += gradient[1]
  return count

if __name__ == "__main__":
  with open('input.txt','r') as file:
    lines = [line.strip() for line in file]

  part1 = checker(lines, (3, 1))
  print(part1)

  slopes: List[Tuple[int, int]] = [
    (1, 1),
    (3, 1),
    (5, 1),
    (7, 1),
    (1, 2),
  ]

  part2 = reduce(lambda x, y: x * y, [checker(lines, slope) for slope in slopes])
  print(part2)
