import re
from collections import defaultdict
from typing import DefaultDict, List, Tuple

class Line:
  def __init__(self, x1: int, y1: int, x2: int, y2: int) -> None:
    self.x1 = x1
    self.y1 = y1
    self.x2 = x2
    self.y2 = y2
    self.horizontal = y1 == y2
    self.vertical = x1 == x2

  def mark_on_grid(self, grid: DefaultDict[Tuple[int, int], int]) -> None:
    x_inc = (self.x1 < self.x2) - (self.x1 > self.x2)
    y_inc = (self.y1 < self.y2) - (self.y1 > self.y2)
    x = self.x1
    y = self.y1
    while x != self.x2 + x_inc or y != self.y2 + y_inc:
      grid[(x, y)] += 1
      x += x_inc
      y += y_inc

def parse(puzzle_input: List[str]) -> List[Line]:
  lines: List[Line] = []
  pattern = r'(\d+),(\d+) -> (\d+),(\d+)'
  for row in puzzle_input:
    match = re.match(pattern, row)
    if match:
      x1, y1, x2, y2 = match.groups()
      lines.append(Line(int(x1), int(y1), int(x2), int(y2)))
  return lines

def part1(lines: List[Line]) -> Tuple[DefaultDict[Tuple[int, int], int], List[Line]]:
  grid: DefaultDict[Tuple[int, int], int] = defaultdict(int)
  diagonal_lines: List[Line] = []
  for line in lines:
    if line.horizontal or line.vertical:
      line.mark_on_grid(grid)
    else:
      diagonal_lines.append(line)
  return grid, diagonal_lines

def part2(grid: DefaultDict[Tuple[int, int], int], lines: List[Line]) -> int:
  for line in lines:
    line.mark_on_grid(grid)
  return len([a for a in grid if grid[a] >= 2])

if __name__ == "__main__":
  with open('input.txt','r') as file:
    puzzle_input = [(line.strip()) for line in file]
  lines = parse(puzzle_input)
  grid, diagonal_lines = part1(lines)
  print(len([a for a in grid if grid[a] >= 2]))
  print(part2(grid, diagonal_lines))