import re

class Match:
  def __init__(self, lo: int, hi: int, char: str, password: str):
    self.lo = lo
    self.hi = hi
    self.char = char
    self.password = password

  def valid1(self) -> bool:
    n = self.password.count(self.char)
    return self.lo <= n <= self.hi

  def valid2(self) -> bool:
    if self.password[self.lo - 1] == self.char and self.password[self.hi - 1] == self.char:
      return False
    return self.password[self.lo - 1] == self.char or self.password[self.hi - 1] == self.char

if __name__ == "__main__":
  file1 = open('input.txt','r')
  lines = file1.readlines()

  pattern = r'(\d+)-(\d+) (\w): (\w+)'
  matches = [re.search(pattern, line) for line in lines]
  matches2 = []
  for match in matches:
    matches2.append(Match(int(match[1]), int(match[2]), match[3], match[4]))

  part1 = [match.valid1() for match in matches2]
  part2 = [match.valid2() for match in matches2]

  print(sum(part1))
  print(sum(part2))