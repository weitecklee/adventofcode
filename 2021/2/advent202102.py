from typing import List

def part1(lines: List[str]) -> int:
  pos, depth = 0, 0
  for line in lines:
    action, val = line.split(' ')
    if action == 'forward':
      pos += int(val)
    elif action == 'down':
      depth += int(val)
    else:
      depth -= int(val)
  return pos * depth

def part2(lines: List[str]) -> int:
  pos, depth, aim = 0, 0, 0
  for line in lines:
    action, val = line.split(' ')
    if action == 'forward':
      pos += int(val)
      depth += aim * int(val)
    elif action == 'down':
      aim += int(val)
    else:
      aim -= int(val)
  return pos * depth

if __name__ == "__main__":
  with open('input.txt','r') as file:
    lines = [line.strip() for line in file]
  print(part1(lines))
  print(part2(lines))
