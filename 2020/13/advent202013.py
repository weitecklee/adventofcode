import re
from typing import Tuple

def part1(lines: list[str]) -> int:
  ptn = r'\d+'
  timestamp = int(lines[0])
  matches = re.findall(ptn, lines[1])
  minWait: int = 1000
  res: int = 0
  if matches is None:
    raise Exception('Regex error', lines[1])
  for match in matches:
    bus = int(match)
    wait = bus - timestamp % bus
    if wait < minWait:
      res = bus * wait
      minWait = wait
  return res

def part2(lines: list[str]) -> int:
  nums = lines[1].split(',')
  ptn = r'\d+'
  schedule: list[Tuple[int, int]] = []
  for i, num in enumerate(nums):
    match = re.match(ptn, num)
    if match is not None:
      n = int(match.group())
      schedule.append((n, i))
  timestamp = schedule[0][0]
  period = schedule[0][0]
  for sched in schedule[1:]:
    depart, i = sched
    while (timestamp + i) % depart > 0:
      timestamp += period
    period *= depart
  return timestamp

if __name__ == "__main__":
  with open('input.txt','r') as file:
    lines = [line.strip() for line in file]
  print(part1(lines))
  print(part2(lines))
