import re
import os
from typing import Tuple

def part1(steps: list[Tuple[str, int]]) -> int:
  pos: list[int] = [0, 0]
  direc: list[int] = [1, 0]
  for step in steps:
    tmp_direc: list[int] = [0, 0]
    if step[0] == 'N':
      tmp_direc[1] = 1
    elif step[0] == 'E':
      tmp_direc[0] = 1
    elif step[0] == 'W':
      tmp_direc[0] = -1
    elif step[0] == 'S':
      tmp_direc[1] = -1
    elif step[0] == 'F':
      tmp_direc = direc
    elif step[0] == 'L':
      if step[1] == 90:
        direc[0], direc[1] = -direc[1], direc[0]
      elif step[1] == 180:
        direc[0], direc[1] = -direc[0], -direc[1]
      elif step[1] == 270:
        direc[0], direc[1] = direc[1], -direc[0]
      else:
        raise Exception('Unexpected step: ', step)
    elif step[0] == 'R':
      if step[1] == 90:
        direc[0], direc[1] = direc[1], -direc[0]
      elif step[1] == 180:
        direc[0], direc[1] = -direc[0], -direc[1]
      elif step[1] == 270:
        direc[0], direc[1] = -direc[1], direc[0]
      else:
        raise Exception('Unexpected step: ', step)
    else:
      raise Exception('Unexpected step: ', step)
    pos[0] += tmp_direc[0] * step[1]
    pos[1] += tmp_direc[1] * step[1]
  return abs(pos[0]) + abs(pos[1])

def part2(steps: list[Tuple[str, int]]) -> int:
  pos: list[int] = [0, 0]
  wayp: list[int] = [10, 1]
  for step in steps:
    if step[0] == 'F':
      pos[0] += wayp[0] * step[1]
      pos[1] += wayp[1] * step[1]
    else:
      tmp_direc: list[int] = [0, 0]
      if step[0] == 'N':
        tmp_direc[1] = 1
      elif step[0] == 'E':
        tmp_direc[0] = 1
      elif step[0] == 'W':
        tmp_direc[0] = -1
      elif step[0] == 'S':
        tmp_direc[1] = -1
      elif step[0] == 'L':
        if step[1] == 90:
          wayp[0], wayp[1] = -wayp[1], wayp[0]
        elif step[1] == 180:
          wayp[0], wayp[1] = -wayp[0], -wayp[1]
        elif step[1] == 270:
          wayp[0], wayp[1] = wayp[1], -wayp[0]
        else:
          raise Exception('Unexpected step: ', step)
      elif step[0] == 'R':
        if step[1] == 90:
          wayp[0], wayp[1] = wayp[1], -wayp[0]
        elif step[1] == 180:
          wayp[0], wayp[1] = -wayp[0], -wayp[1]
        elif step[1] == 270:
          wayp[0], wayp[1] = -wayp[1], wayp[0]
        else:
          raise Exception('Unexpected step: ', step)
      else:
        raise Exception('Unexpected step: ', step)
      wayp[0] += tmp_direc[0] * step[1]
      wayp[1] += tmp_direc[1] * step[1]
  return abs(pos[0]) + abs(pos[1])

def parse(lines: list[str]) -> list[Tuple[str, int]]:
  steps: list[Tuple[str, int]] = list()
  pattern = r'(\w)(\d+)'
  for line in lines:
    m = re.match(pattern, line)
    if m is None:
      raise Exception("Line does not match regex: ", line)
    steps.append((m.group(1), int(m.group(2))))
  return steps

if __name__ == "__main__":
  with open(os.path.join(os.path.dirname(os.path.abspath(__file__)), 'input.txt'),'r') as file:
    lines = [line.strip() for line in file]
  steps = parse(lines)
  print(part1(steps))
  print(part2(steps))
