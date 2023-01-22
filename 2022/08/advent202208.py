file1 = open('input202208.txt','r')
lines = list(map(str.strip, file1.readlines()))

grid = [[0]*len(lines) for i in range(len(lines))]

for row in grid:
  row[0] = 1
  row[-1] = 1

grid[0] = [1] * len(lines)
grid[-1] = [1] * len(lines)

for [i, row] in enumerate(lines[1:-1], start = 1):
  ht = row[0]
  for [j, tree] in enumerate(row.strip()[1: -1], start = 1):
    if tree > ht:
      grid[i][j] = 1
      ht = tree
  ht = row[-1]
  for j in range(len(lines) - 2, 0, -1):
    if lines[i][j] > ht:
      grid[i][j] = 1
      ht = lines[i][j]

for j in range(1, len(lines)):
  ht = lines[0][j]
  for i in range(1, len(lines) - 1):
    if lines[i][j] > ht:
      grid[i][j] = 1
      ht = lines[i][j]
  ht = lines[-1][j]
  for i in range(len(lines) - 2, 0, -1):
    if lines[i][j] > ht:
      grid[i][j] = 1
      ht = lines[i][j]

print(sum(map(sum, grid)))

maxScore = 0

for i in range(1, len(lines) - 1):
  for j in range(1, len(lines) - 1):
    ht = lines[i][j]
    up = 0
    for k in range(i - 1, -1, -1):
      up += 1
      if lines[k][j] >= ht:
        break
    down = 0
    for k in lines[(i + 1):]:
      down += 1
      if k[j] >= ht:
        break
    left = 0
    for k in reversed(lines[i][:j]):
      left += 1
      if k >= ht:
        break
    right = 0
    for k in lines[i][(j + 1):]:
      right += 1
      if k >= ht:
        break
    score = up * down * left * right
    maxScore = max(maxScore, score)

print(maxScore)