def part1(lines: list[int]) -> int:
  numbers: set[int] = set()
  numbers.update(lines[:25])
  for i in range(25, len(lines)):
    check = lines[i]
    fail = True
    for n in numbers:
      if check - n in numbers and n != check - n:
        fail = False
        break
    if fail:
      return check
    numbers.add(check)
    numbers.remove(lines[i - 25])
  return -1

def part2(lines: list[int], target: int) -> int:
  curr = 0
  start = 0
  for i, num in enumerate(lines):
    curr += num
    while curr > target and start < i:
      curr -= lines[start]
      start += 1
    if curr == target:
      return min(lines[start: i + 1]) + max(lines[start: i + 1])
  return -1

if __name__ == "__main__":
  file1 = open('input.txt','r')
  lines = [int(line.strip()) for line in file1.readlines()]
  target = part1(lines)
  print(target)
  print(part2(lines, target))
