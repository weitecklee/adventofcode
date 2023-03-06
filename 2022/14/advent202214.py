import re

file1 = open('input.txt','r')
lines = file1.readlines()

rocks = set()
maxY = 0

for line in lines:
  nums = list(map(int, re.findall('\d+', line)))
  for i in range(0, len(nums) - 2, 2):
    x0 = nums[i]
    y0 = nums[i + 1]
    x1 = nums[i + 2]
    y1 = nums[i + 3]
    maxY = max(maxY, y0, y1)
    if x0 < x1:
      for i in range(x0, x1 + 1):
        rocks.add((i, y0))
    elif x1 < x0:
      for i in range(x1, x0 + 1):
        rocks.add((i, y0))
    elif y0 < y1:
      for i in range(y0, y1 + 1):
        rocks.add((x0, i))
    elif y1 < y0:
      for i in range(y1, y0 + 1):
        rocks.add((x0, i))

trail = [(500, 0)]

rest = False

def trailer():
  global trail
  x, y = trail[-1]
  rest = False
  while not rest and y < maxY + 1:
    y += 1
    if (x, y) not in rocks:
      trail.append((x, y))
    elif (x - 1, y) not in rocks:
      x -= 1
      trail.append((x, y))
    elif (x + 1, y) not in rocks:
      x += 1
      trail.append((x, y))
    else:
      rest = True
  return rest

sand = 0

while trailer():
  rocks.add(trail.pop())
  sand += 1

print(sand)

while len(trail):
  trailer()
  rocks.add(trail.pop())
  sand += 1

print(sand)





