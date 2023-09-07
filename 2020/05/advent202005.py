class Seat:
  def __init__(self, seat: str) -> None:
    row = 0
    for i in range(7):
      if seat[i] == 'F':
        row <<= 1
      else:
        row = (row << 1) | 1
    col = 0
    for i in range(7, 10):
      if seat[i] == 'L':
        col <<= 1
      else:
        col = (col << 1) | 1
    self.seat = seat
    self.row = row
    self.col = col
    self.seatID = row * 8 + col

def part1(seatIDs: list[int]) -> int:
  return seatIDs[-1]

def part2(seatIDs: list[int]) -> int:
  for i in range(len(seatIDs) - 1):
    if seatIDs[i] == seatIDs[i + 1] - 2:
      return seatIDs[i] + 1
  return -1

if __name__ == "__main__":
  file1 = open('input.txt','r')
  lines = [line.strip() for line in file1.readlines()]

  seats: list[Seat] = [Seat(line) for line in lines]
  seatIDs: list[int] = [seat.seatID for seat in seats]
  seatIDs.sort()
  print(part1(seatIDs))
  print(part2(seatIDs))

