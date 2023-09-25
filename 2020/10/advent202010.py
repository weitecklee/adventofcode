from collections import defaultdict

def part1(lines: list[int]) -> int:
  count1 = 0
  count3 = 0
  for i in range(1, len(lines)):
    diff = lines[i] - lines[i - 1]
    if diff == 1:
      count1 += 1
    elif diff == 3:
      count3 += 1
  return count1 * count3

def part2(lines: list[int]) -> int:
  ways = [0] * (lines[-1] + 1)
  ways[lines[-1]] = 1
  for i in reversed(lines[:-1]):
    ways[i] = ways[i + 1] + ways[i + 2] + ways[i + 3]
  return ways[0]

if __name__ == "__main__":
  with open('input.txt','r') as file:
    lines = [int(line.strip()) for line in file]
  lines.append(0)
  lines.sort()
  lines.append(lines[-1] + 3)
  print(part1(lines))
  print(part2(lines))
