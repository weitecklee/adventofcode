file1 = open('input.txt','r')
lines = file1.readlines()

horizontal = 0
depth = 0

for line in lines:
  direction, dist = line.strip().split(' ')
  if direction == 'forward':
    horizontal += int(dist)
  elif direction == 'down':
    depth += int(dist)
  elif direction == 'up':
    depth -= int(dist)

print(horizontal * depth)

horizontal = 0
depth = 0
aim = 0

for line in lines:
  direction, dist = line.strip().split(' ')
  if direction == 'forward':
    horizontal += int(dist)
    depth += int(dist) * aim
  elif direction == 'down':
    aim += int(dist)
  elif direction == 'up':
    aim -= int(dist)

print(horizontal * depth)