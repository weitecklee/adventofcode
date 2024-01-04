from typing import List, Tuple

class Seat:
  def __init__(self, seat: str) -> None:
    self.seat = seat
    self.row, self.col = self._calculate_seat()
    self.seatID = self.row * 8 + self.col

  def _calculate_seat(self) -> Tuple[int, int]:
    row = int(self.seat[:7].replace('F', '0').replace('B', '1'), 2)
    col = int(self.seat[7:].replace('L', '0').replace('R', '1'), 2)
    return row, col

def part1(seatIDs: List[int]) -> int:
  return seatIDs[-1]

def part2(seatIDs: List[int]) -> int:
  for i in range(len(seatIDs) - 1):
    if seatIDs[i] == seatIDs[i + 1] - 2:
      return seatIDs[i] + 1
  return -1

if __name__ == "__main__":
  with open('input.txt','r') as file:
    lines = [line.strip() for line in file]

  seats: List[Seat] = [Seat(line) for line in lines]
  seatIDs: List[int] = [seat.seatID for seat in seats]
  seatIDs.sort()
  print(part1(seatIDs))
  print(part2(seatIDs))

