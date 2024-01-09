import re
import os

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
    return (self.password[self.lo - 1] == self.char) != (self.password[self.hi - 1] == self.char)

if __name__ == "__main__":
  with open(os.path.join(os.path.dirname(os.path.abspath(__file__)), 'input.txt'),'r') as file:
    lines = [line.strip() for line in file]

  pattern = r'(\d+)-(\d+) (\w): (\w+)'
  matches = [re.match(pattern, line) for line in lines]
  matches2 = [Match(int(match.group(1)), int(match.group(2)), match.group(3), match.group(4)) for match in matches if match]

  part1 = [match.valid1() for match in matches2]
  part2 = [match.valid2() for match in matches2]

  print(sum(part1))
  print(sum(part2))