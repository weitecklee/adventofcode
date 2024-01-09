import os
from collections import defaultdict
from typing import DefaultDict, List, Set, Tuple

def parse(lines: List[str]) -> Set[Tuple[int, int, int, int]]:
  cubes: Set[Tuple[int, int, int, int]] = set()
  for i, line in enumerate(lines):
    for j, c in enumerate(line):
      if c == '#':
        cubes.add((i, j, 0, 0))
  return cubes

def simulate(cubes: Set[Tuple[int, int, int, int]], part2: bool = False) -> Set[Tuple[int, int, int, int]]:
  cubes2: Set[Tuple[int, int, int, int]] = set()
  neighbors: DefaultDict[Tuple[int, int, int, int], int] = defaultdict(int)
  for cube in cubes:
    x, y, z, w = cube
    for a in range(-1, 2):
      for b in range(-1, 2):
        for c in range(-1, 2):
          if part2:
            for d in range(-1, 2):
              if a == b == c == d == 0:
                continue
              coord = (x + a, y + b, z + c, w + d)
              neighbors[coord] += 1
          else:
            if a == b == c == 0:
              continue
            coord = (x + a, y + b, z + c, w)
            neighbors[coord] += 1
  for cube in cubes:
    if 2 <= neighbors[cube] <= 3:
      cubes2.add(cube)
  for cube, n in neighbors.items():
    if n == 3 and cube not in cubes:
      cubes2.add(cube)
  return cubes2

def part1(cubes: Set[Tuple[int, int, int, int]], cycles: int) -> int:
  for _ in range(cycles):
    cubes = simulate(cubes)
  return len(cubes)

def part2(cubes: Set[Tuple[int, int, int, int]], cycles: int) -> int:
  for _ in range(cycles):
    cubes = simulate(cubes, True)
  return len(cubes)


if __name__ == "__main__":
  with open(os.path.join(os.path.dirname(os.path.abspath(__file__)), 'input.txt'),'r') as file:
    lines = [line.strip() for line in file]
  cubes = parse(lines)
  print(part1(cubes, 6))
  print(part2(cubes, 6))
