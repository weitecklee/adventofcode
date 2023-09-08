from collections import defaultdict
from typing import Tuple

def part1(seats: dict[Tuple[int, int], bool]) -> int:
  changed = True
  while changed:
    changed = False
    occupied = countOccupied(seats)
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
    occupied = countOccupied2(seats)
    for seat, occ in seats.items():
      if occ and occupied[seat] >= 5:
        seats[seat] = False
        changed = True
      elif not occ and occupied[seat] == 0:
        seats[seat] = True
        changed = True
  return sum(seats.values())

def countOccupied(seats: dict[Tuple[int, int], bool]) -> defaultdict[Tuple[int, int], int]:
  occupied: defaultdict[Tuple[int, int], int] = defaultdict(lambda: 0)
  for seat, occ in seats.items():
    if occ:
      for check in toCheck:
        checkSeat = (seat[0] + check[0], seat[1] + check[1])
        occupied[checkSeat] += 1
  return occupied

def countOccupied2(seats: dict[Tuple[int, int], bool]) -> defaultdict[Tuple[int, int], int]:
  occupied: defaultdict[Tuple[int, int], int] = defaultdict(lambda: 0)
  for seat, occ in seats.items():
    if occ:
      for check in toCheck:
        k = 1
        while True:
          checkSeat = (seat[0] + k * check[0], seat[1] + k * check[1])
          if 0 <= checkSeat[0] <= r and 0 <= checkSeat[1] <= c:
            if checkSeat in seats:
              occupied[checkSeat] += 1
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
  file1 = open('input.txt','r')
  grid = [line.strip() for line in file1.readlines()]
  seats = parse(grid)
  seats2 = dict.copy(seats)
  c = len(grid) - 1
  r = len(grid[0]) - 1
  toCheck: list[Tuple[int, int]] = [(-1, -1), (-1, 0), (-1, 1), (0, -1), (0, 1), (1, -1), (1, 0), (1, 1)]
  print(part1(seats))
  print(part2(seats2))
