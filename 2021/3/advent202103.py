from typing import List

def find_common_bit(lines: List[str], place: int) -> int:
  bit = 0
  for line in lines:
    bit += 1 if line[place] == '1' else 0
  return 1 if bit >= len(lines) - bit else 0

def part1(lines: List[str]) -> int:
  gamma: List[str] = []
  epsilon: List[str] = []
  for i in range(len(lines[0])):
    bit = find_common_bit(lines, i)
    gamma.append(str(bit))
    epsilon.append(str(1 - bit))
  return int(''.join(gamma), 2) * int(''.join(epsilon), 2)

def part2(lines: List[str]) -> int:
  oxygen = lines
  co2 = lines
  place = 0
  while len(oxygen) > 1:
    bit = find_common_bit(oxygen, place)
    oxygen = [line for line in oxygen if line[place] == str(bit)]
    place += 1
  place = 0
  while len(co2) > 1:
    bit = 1 - find_common_bit(co2, place)
    co2 = [line for line in co2 if line[place] == str(bit)]
    place += 1
  return int(''.join(oxygen.pop()), 2) * int(''.join(co2.pop()), 2)

if __name__ == "__main__":
  with open('input.txt','r') as file:
    lines = [(line.strip()) for line in file]
  print(part1(lines))
  print(part2(lines))
