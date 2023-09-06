def checker(lines: list[str], gradient: tuple[int, int]) -> int:
  pos: list[int] = [0, 0]
  count: int = 0
  h, w = len(lines), len(lines[0])
  while pos[1] < h:
    if lines[pos[1]][pos[0]] == "#":
      count += 1
    pos[0] += gradient[0]
    pos[1] += gradient[1]
    pos[0] %= w
  return count

if __name__ == "__main__":
  file1 = open('input.txt','r')
  lines = [line.strip() for line in file1.readlines()]

  print(checker(lines, (3, 1)))

  slopes: list[tuple[int, int]] = [
    (1, 1),
    (3, 1),
    (5, 1),
    (7, 1),
    (1, 2),
  ]

  part2 = 1
  for slope in slopes:
    part2 *= checker(lines, slope)
  print(part2)
