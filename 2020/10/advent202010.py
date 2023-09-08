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
  ways: defaultdict[int, int] = defaultdict(lambda: 0)
  ways[lines[-1]] = 1
  for i in range(len(lines) - 2, -1, -1):
    j = lines[i]
    ways[j] = ways[j + 1] + ways[j + 2] + ways[j + 3]
  return ways[0]

if __name__ == "__main__":
  file1 = open('input.txt','r')
  lines = [int(line.strip()) for line in file1.readlines()]
  lines.append(0)
  lines.sort()
  lines.append(lines[-1] + 3)
  print(part1(lines))
  print(part2(lines))
