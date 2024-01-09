import os
from collections import defaultdict
from typing import Tuple

def part1(seats: dict[Tuple[int, int], bool]) -> int:
  changed = True
  while changed:
    changed = False
    occupied = count_occupied(seats)
    for seat, occ in seats.items():
      if occ and occupied[seat] >= 4:
        seats[seat] = False
        changed = True
      elif not occ and occupied[seat] == 0:
        seats[seat] = True
        changed = True
  return sum(seats.values())

def part2(seats: dict[Tuple[int, int], bool]) -> int:
  changed = True
  while changed:
    changed = False
    occupied = count_occupied2(seats)
    for seat, occ in seats.items():
      if occ and occupied[seat] >= 5:
        seats[seat] = False
        changed = True
      elif not occ and occupied[seat] == 0:
        seats[seat] = True
        changed = True
  return sum(seats.values())

def count_occupied(seats: dict[Tuple[int, int], bool]) -> defaultdict[Tuple[int, int], int]:
  occupied: defaultdict[Tuple[int, int], int] = defaultdict(lambda: 0)
  for seat, occ in seats.items():
    if occ:
      for check in to_check:
        check_seat = (seat[0] + check[0], seat[1] + check[1])
        occupied[check_seat] += 1
  return occupied

def count_occupied2(seats: dict[Tuple[int, int], bool]) -> defaultdict[Tuple[int, int], int]:
  occupied: defaultdict[Tuple[int, int], int] = defaultdict(lambda: 0)
  for seat, occ in seats.items():
    if occ:
      for check in to_check:
        k = 1
        while True:
          check_seat = (seat[0] + k * check[0], seat[1] + k * check[1])
          if 0 <= check_seat[0] <= r and 0 <= check_seat[1] <= c:
            if check_seat in seats:
              occupied[check_seat] += 1
              break
          else:
            break
          k += 1
  return occupied

def parse(grid: list[str]) -> dict[Tuple[int, int], bool]:
  seats: dict[Tuple[int, int], bool] = dict()
  for j in range(len(grid)):
    for i in range(len(grid[j])):
      if grid[j][i] == 'L':
        seats[(i, j)] = False
  return seats

if __name__ == "__main__":
  with open(os.path.join(os.path.dirname(os.path.abspath(__file__)), 'input.txt'),'r') as file:
    grid = [line.strip() for line in file]
  seats = parse(grid)
  seats2 = dict.copy(seats)
  c = len(grid) - 1
  r = len(grid[0]) - 1
  to_check: list[Tuple[int, int]] = [(-1, -1), (-1, 0), (-1, 1), (0, -1), (0, 1), (1, -1), (1, 0), (1, 1)]
  print(part1(seats))
  print(part2(seats2))
