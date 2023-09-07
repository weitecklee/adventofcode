def part1(lines: list[str]) -> int:
  tmp = ''
  count = 0
  for line in lines:
    if line == '':
      yeses = set(list(tmp))
      count += len(yeses)
      tmp = ''
    else:
      tmp += line
  yeses = set(list(tmp))
  return count + len(yeses)

def part2(lines: list[str]) -> int:
  count = 0
  yeses: dict[str, int] = {}
  members = 0
  for line in lines:
    if line == '':
      for n in yeses.values():
        if n == members:
          count += 1
      yeses: dict[str, int] = {}
      members = 0
    else:
      members += 1
      for c in line:
        if c in yeses:
          yeses[c] += 1
        else:
          yeses[c] = 1
  for n in yeses.values():
    if n == members:
      count += 1
  return count


if __name__ == "__main__":
  file1 = open('input.txt','r')
  lines = [line.strip() for line in file1.readlines()]
  print(part1(lines))
  print(part2(lines))
