import re
from typing import List

def part1(lines: List[str]) -> int:
  memory: dict[int, int] = {}
  pattern1 = r'mask = (.*)'
  pattern2 = r'mem\[(.*)\] = (.*)'
  mask = ''
  address = 0
  val = 0
  for line in lines:
    if match1 := re.search(pattern1, line):
      mask = match1.group(1)
    elif match2 := re.search(pattern2, line):
      address = int(match2.group(1))
      val = format(int(match2.group(2)), f'0{len(mask)}b')
      masked_val = []
      for i, c in enumerate(mask):
        if c == 'X':
          masked_val.append(val[i])
        else:
          masked_val.append(c)
      memory[address] = int(''.join(masked_val), 2)
  return sum(memory.values())

def part2(lines: List[str]) -> int:
  memory: dict[int, int] = {}
  pattern1 = r'mask = (.*)'
  pattern2 = r'mem\[(.*)\] = (.*)'
  mask = ''
  address = 0
  val = 0
  for line in lines:
    if match1 := re.search(pattern1, line):
      mask = match1.group(1)
    elif match2 := re.search(pattern2, line):
      address = format(int(match2.group(1)), f'0{len(mask)}b')
      val = int(match2.group(2))
      masked_address = []
      for i, c in enumerate(mask):
        if c == '0':
          masked_address.append(address[i])
        else:
          masked_address.append(c)
      possible_addresses = generate_possible_addresses(masked_address)
      for address in possible_addresses:
        memory[address] = val
  return sum(memory.values())

def generate_possible_addresses(masked_address) -> List[int] :
  possible_addresses: List[List[str]] = [[]]
  for c in masked_address:
    if c == 'X':
      tmp = []
      tmp.extend([address + ['0'] for address in possible_addresses])
      tmp.extend([address + ['1'] for address in possible_addresses])
      possible_addresses = tmp
    else:
      for address in possible_addresses:
        address.append(c)
  return [int(''.join(address), 2) for address in possible_addresses]

if __name__ == "__main__":
  with open('input.txt','r') as file:
    lines = [line.strip() for line in file]
  print(part1(lines))
  print(part2(lines))
