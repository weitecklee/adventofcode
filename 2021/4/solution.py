import os
from typing import Dict, List, Tuple

class Board:
  def __init__(self, grid: List[List[int]]) -> None:
    self.grid: Dict[int, Tuple[int, int]] = {}
    for row, a in enumerate(grid):
      for col, b in enumerate(a):
          self.grid[b] = row, col
    self.rows: Dict[int, int] = {}
    self.cols: Dict[int, int] = {}
    for i in range(5):
      self.rows[i] = 0
      self.cols[i] = 0

  def mark(self, n: int) -> bool:
    if n in self.grid:
      row, col = self.grid[n]
      self.rows[row] += 1
      self.cols[col] += 1
      self.grid.pop(n)
      return self.rows[row] == 5 or self.cols[col] == 5
    return False

  def unmarked_sum(self) -> int:
    return sum(self.grid.keys())

def parse(lines: List[str]) -> Tuple[List[int], List[Board]]:
  boards: List[Board] = []
  numbers = [int(n) for n in lines[0].split(',')]
  for i in range(2, len(lines), 6):
    grid: List[List[int]] = []
    for j in range(5):
      row = [int(n) for n in lines[i + j].split()]
      grid.append(row)
    boards.append(Board(grid))
  return numbers, boards

def part1(numbers: List[int], boards: List[Board]) -> int:
  for n in numbers:
    for board in boards:
      if board.mark(n):
        return board.unmarked_sum() * n
  return -1

def part2(numbers: List[int], boards: List[Board]) -> int:
  for n in numbers:
    boards2: List[Board] = []
    for board in boards:
      bingo = board.mark(n)
      if bingo and len(boards) == 1:
        return board.unmarked_sum() * n
      elif not bingo:
        boards2.append(board)
    boards = boards2
  return -1

if __name__ == "__main__":
  with open(os.path.join(os.path.dirname(os.path.abspath(__file__)), 'input.txt'),'r') as file:
    lines = [(line.strip()) for line in file]
  numbers, boards = parse(lines)
  print(part1(numbers, boards))
  print(part2(numbers, boards))
