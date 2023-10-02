import re
from typing import List, Set, Tuple

def parse(puzzle_input: List[str]) -> List[Tuple[int, int, int]]:
  parsed_input: List[Tuple[int, int, int]] = []
  pattern = r'(\d+)\s+(\d+)\s+(\d+)'
  for line in puzzle_input:
    match = re.search(pattern, line)
    if match:
      parsed_input.append((int(match.group(1)), int(match.group(2)), int(match.group(3))))
  return parsed_input

def is_possible_triangle( a: int = 0, b: int = 0, c: int = 0, triangle: Tuple[int, int, int] | None = None) -> bool:
  if triangle is not None:
    a, b, c = triangle
  return a + b > c and a + c > b and b + c > a

def part1(triangles: List[Tuple[int, int, int]]) -> int:
  return sum([is_possible_triangle(triangle = triangle) for triangle in triangles])

def part2(lines: List[Tuple[int, int, int]]) -> int:
  count = 0
  for col in range(3):
    for row in range(0, len(lines), 3):
      if is_possible_triangle(lines[row][col], lines[row + 1][col], lines[row + 2][col]):
        count += 1
  return count

if __name__ == '__main__':
  with open('input.txt', 'r') as file:
    puzzle_input = [line.strip() for line in file]
  parsed_input = parse(puzzle_input)
  print(part1(parsed_input))
  print(part2(parsed_input))

