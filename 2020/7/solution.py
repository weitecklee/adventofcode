import re
import os
from collections import defaultdict
from typing import DefaultDict, Dict, List, Set, Tuple

def parse(lines: List[str]) -> Tuple[DefaultDict[str, Dict[str, int]], DefaultDict[str, Set[str]]]:
  bags1: DefaultDict[str, Dict[str, int]] = defaultdict(dict)
  bags2: DefaultDict[str, Set[str]] = defaultdict(set)
  pattern1 = r'\w+ \w+ bag'
  pattern2 = r'(\d+) (\w+ \w+ bag)'
  for line in lines:
    outer_bag_match = re.match(pattern1, line)
    if outer_bag_match is None:
      raise Exception("outer_bag_match is None")
    outer_bag = outer_bag_match.group()
    inner_bags = re.finditer(pattern2, line)
    for inner_bag in inner_bags:
      bag = inner_bag.group(2)
      n = int(inner_bag.group(1))
      bags1[outer_bag][bag] = n
      bags2[bag].add(outer_bag)
  return bags1, bags2

def part1(bags: DefaultDict[str, Set[str]], target: str) -> int:
  checked: Set[str] = set()
  to_check = [target]
  while to_check:
    check_bag = to_check.pop()
    if check_bag in bags:
      for bag in bags[check_bag]:
        if bag not in checked:
          to_check.append(bag)
          checked.add(bag)
  return len(checked)

def recur(bags: DefaultDict[str, Dict[str, int]], bag: str) -> int:
  count = 1
  for inner_bag, n in bags[bag].items():
    count += n * recur(bags, inner_bag)
  return count

def part2(bags: DefaultDict[str, Dict[str, int]], target: str) -> int:
  return recur(bags, target) - 1

if __name__ == "__main__":
  with open(os.path.join(os.path.dirname(os.path.abspath(__file__)), 'input.txt'),'r') as file:
    lines = [line.strip() for line in file]
  bags1, bags2 = parse(lines)
  target = 'shiny gold bag'
  print(part1(bags2, target))
  print(part2(bags1, target))
