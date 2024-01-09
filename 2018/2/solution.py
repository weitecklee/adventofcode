import os
from collections import Counter

def part1(lines: list[str]) -> int:
  count2 = 0
  count3 = 0
  for line in lines:
    char_count = Counter(line)
    if 2 in char_count.values(): count2 += 1
    if 3 in char_count.values(): count3 += 1
  return count2 * count3

def part2(lines: list[str]) -> str:
  for i, line in enumerate(lines):
    for line2 in lines[i+1:]:
      diff = 0
      for j, char in enumerate(line2):
        if char != line[j]:
          diff += 1
          if diff > 1:
            break
      if diff == 1:
        res: list[str] = []
        for j, char in enumerate(line2):
          if char == line[j]:
            res.append(char)
        return ''.join(res)
  return ''

if __name__ == "__main__":
  with open(os.path.join(os.path.dirname(os.path.abspath(__file__)), 'input.txt'),'r') as file:
    lines = [line.strip() for line in file]

  print(part1(lines))
  print(part2(lines))