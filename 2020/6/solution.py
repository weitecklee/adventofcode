from collections import defaultdict
import os

def part1(lines: list[str]) -> int:
  tmp = ''
  count = 0
  for line in lines:
    if not line:
      yeses = set(list(tmp))
      count += len(yeses)
      tmp = ''
    else:
      tmp += line
  yeses = set(list(tmp))
  return count + len(yeses)

def part2(lines: list[str]) -> int:
  count = 0
  yeses = defaultdict(int)
  members = 0
  for line in lines:
    if not line:
      count += sum(1 for n in yeses.values() if n == members)
      yeses = defaultdict(int)
      members = 0
    else:
      members += 1
      for c in line:
        yeses[c] += 1
  count += sum(1 for n in yeses.values() if n == members)
  return count

if __name__ == "__main__":
  with open(os.path.join(os.path.dirname(os.path.abspath(__file__)), 'input.txt'),'r') as file:
    lines = [line.strip() for line in file]

  print(part1(lines))
  print(part2(lines))
