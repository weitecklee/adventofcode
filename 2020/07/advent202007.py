import re
from typing import Tuple

def parse(lines: list[str]) -> Tuple[dict[str, dict[str, int]], dict[str, set[str]]]:
  bags1: dict[str, dict[str, int]] = {}
  bags2: dict[str, set[str]] = {}
  pattern1 = r'\w+ \w+ bag'
  pattern2 = r'(\d+) (\w+ \w+ bag)'
  for line in lines:
    outerBag = re.search(pattern1, line).group()
    innerBags = re.finditer(pattern2, line)
    bags1[outerBag] = {}
    for bagMatch in innerBags:
      bag = bagMatch.group(2)
      n = int(bagMatch.group(1))
      bags1[outerBag][bag] = n
      if bag not in bags2:
        bags2[bag] = set()
      bags2[bag].add(outerBag)
  return bags1, bags2

def part1(bags: dict[str, set[str]]) -> int:
  checked: set[str] = set()
  toCheck: list[str] = ['shiny gold bag']
  i = 0
  while i < len(toCheck):
    checkBag = toCheck[i]
    if checkBag in bags:
      for bag in bags[checkBag]:
        if bag not in checked:
          toCheck.append(bag)
          checked.add(bag)
    i += 1
  return len(checked)

def recur(bags: dict[str, dict[str, int]], bag: str) -> int:
  count = 1
  for innerBag, n in bags[bag].items():
    count += n * recur(bags, innerBag)
  return count

def part2(bags: dict[str, dict[str, int]]) -> int:
  return recur(bags, 'shiny gold bag') - 1

if __name__ == "__main__":
  file1 = open('input.txt','r')
  lines = [line.strip() for line in file1.readlines()]
  bags1, bags2 = parse(lines)
  print(part1(bags2))
  print(part2(bags1))
