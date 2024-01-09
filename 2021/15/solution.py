import os
file1 = open(os.path.join(os.path.dirname(os.path.abspath(__file__)), 'input.txt'),'r')
lines = file1.readlines()

import heapq
import math
from collections import defaultdict

risks = [list(map(int,list(line.strip()))) for line in lines]

def findMinRisk(grid):
  visited = set()

  h = len(grid)
  w = len(grid[0])

  start = (0, 0)
  dest = (h - 1, w - 1)

  minRisk = defaultdict(lambda: math.inf)
  minRisk[start] = 0

  q = [(0, start)]

  while len(q):
    risk, pos = heapq.heappop(q)
    if pos == dest:
      return risk
    if pos in visited:
      continue
    visited.add(pos)
    r, c = pos
    if r > 0 and (r - 1, c) not in visited:
      newRisk = risk + grid[r - 1][c]
      if newRisk < minRisk[(r - 1, c)]:
        minRisk[(r - 1, c)] = newRisk
        heapq.heappush(q, (newRisk, (r - 1, c)))
    if c > 0 and (r, c - 1) not in visited:
      newRisk = risk + grid[r][c - 1]
      if newRisk < minRisk[(r, c - 1)]:
        minRisk[(r, c - 1)] = newRisk
        heapq.heappush(q, (newRisk, (r, c - 1)))
    if r < h - 1 and (r + 1, c) not in visited:
      newRisk = risk + grid[r + 1][c]
      if newRisk < minRisk[(r + 1, c)]:
        minRisk[(r + 1, c)] = newRisk
        heapq.heappush(q, (newRisk, (r + 1, c)))
    if c < w - 1 and (r, c + 1) not in visited:
      newRisk = risk + grid[r][c + 1]
      if newRisk < minRisk[(r, c + 1)]:
        minRisk[(r, c + 1)] = newRisk
        heapq.heappush(q, (newRisk, (r, c + 1)))

print(findMinRisk(risks))

h = len(risks)
w = len(risks[0])
risks2 = [[0] * 5 * w for _ in range(5 * h)]

for r in range(5):
  for c in range(5):
    for i in range(h):
      for j in range(w):
        tmp = risks[i][j] + r + c
        if tmp > 9:
          tmp -= 9
        risks2[h * r + i][w * c + j] = tmp

print(findMinRisk(risks2))
