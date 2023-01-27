import re

file1 = open('input202113.txt','r')
lines = file1.readlines()
paper = set()
folds = []

for line in lines:
  nums = list(map(int, re.findall('\d+', line)))
  if len(nums) == 2:
    paper.add(tuple(nums))
  elif len(nums) == 1:
    if 'x' in line:
      folds.append(['x', nums[0]])
    else:
      folds.append(['y', nums[0]])

for fold in folds[:1]:
  paper2 = set()
  if fold[0] == 'x':
    for point in paper:
      if point[0] > fold[1]:
        paper2.add((2 * fold[1] - point[0], point[1]))
      else:
        paper2.add(point)
  else:
    for point in paper:
      if point[1] > fold[1]:
        paper2.add((point[0], 2 * fold[1] - point[1]))
      else:
        paper2.add(point)
  paper = paper2

print(len(paper))

for fold in folds[1:]:
  paper2 = set()
  if fold[0] == 'x':
    for point in paper:
      if point[0] > fold[1]:
        paper2.add((2 * fold[1] - point[0], point[1]))
      else:
        paper2.add(point)
  else:
    for point in paper:
      if point[1] > fold[1]:
        paper2.add((point[0], 2 * fold[1] - point[1]))
      else:
        paper2.add(point)
  paper = paper2

xMax = max([x for (x, y) in paper])
yMax = max([y for (x, y) in paper])

code = [['.'] * (xMax + 1) for i in range((yMax + 1))]

for (x, y) in paper:
  code[y][x] = '#'

for line in code:
  print(''.join(line))