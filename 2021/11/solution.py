import os

file1 = open(os.path.join(os.path.dirname(os.path.abspath(__file__)), 'input.txt'),'r')
lines = file1.readlines()
lines = list(map(str.strip, lines))
lines = [list(map(int, list(line))) for line in lines]

def flasher(a, b, q):
  if lines[a][b] > 0:
    lines[a][b] += 1
    if lines[a][b] == 10:
      lines[a][b] = 0
      q.append([a, b])

def flashout(q):
  for i, row in enumerate(lines):
    for j, octopus in enumerate(row):
      lines[i][j] += 1
      if lines[i][j] == 10:
        lines[i][j] = 0
        q.append([i, j])
  for i, j in q:
    if i > 0:
      flasher(i - 1, j, q)
    if j > 0:
      flasher(i, j - 1, q)
    if i < len(lines) - 1:
      flasher(i + 1, j, q)
    if j < len(lines[0]) - 1:
      flasher(i, j + 1, q)
    if i > 0 and j > 0:
      flasher(i - 1, j - 1, q)
    if i > 0 and j < len(lines[0]) - 1:
      flasher(i - 1, j + 1, q)
    if i < len(lines) - 1 and j > 0:
      flasher(i + 1, j - 1, q)
    if i < len(lines) - 1 and j < len(lines[0]) - 1:
      flasher(i + 1, j + 1, q)
  return len(q)

flashes = 0

for step in range(100):
  flashq = []
  flashes += flashout(flashq)

print(flashes)

for step in range(101, 1000):
  flashq = []
  if flashout(flashq) == len(lines) * len(lines[0]):
    print(step)
    break

