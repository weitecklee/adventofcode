file1 = open('input.txt','r')
lines = file1.readlines()

grid = [list(map(int, list(line.strip()))) for line in lines]

risks = 0
basins = []

for i, line in enumerate(grid):
  for j, ht in enumerate(line):
    risk = 0
    if i == 0 or ht < grid[i - 1][j]:
      risk += 1
    if j == 0 or ht < grid[i][j - 1]:
      risk += 1
    if i == len(grid) - 1 or ht < grid[i + 1][j]:
      risk += 1
    if j == len(line) - 1 or ht < grid[i][j + 1]:
      risk += 1
    if risk == 4:
      risks += ht + 1
      basins.append([i, j])

print(risks)

def mapper(point):
  i, j = point
  if grid[i][j] == 9 or grid[i][j] == 'X':
    return 0
  grid[i][j] = 'X'
  size = 1
  if i > 0:
    size += mapper([i - 1, j])
  if j > 0:
    size += mapper([i, j - 1])
  if i < len(grid) - 1:
    size += mapper([i + 1, j])
  if j < len(grid[i]) - 1:
    size += mapper([i, j + 1])
  return size

sizes = []

for basin in basins:
  sizes.append(mapper(basin))

sizes = sorted(sizes, reverse = True)

print(sizes[0] * sizes[1] * sizes[2])
