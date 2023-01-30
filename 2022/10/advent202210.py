file1 = open('input202210.txt','r')
lines = file1.readlines()

register = 1
cycle = 0
cycles = [20, 60, 100, 140, 180, 220]
strength = 0
screen = [['.'] * 40 for _ in range(6)]

def tick():
  global cycle, strength, screen
  if (register - 1) <= (cycle % 40) <= (register + 1):
    screen[cycle // 40][cycle % 40] = '#'
  cycle += 1
  if cycle in cycles:
    strength += cycle * register

for line in lines:
  if line[0] == 'n':
    tick()
  else:
    parse = line.split(' ')
    tick()
    tick()
    register += int(parse[1])

print(strength)

for line in screen:
  print(''.join(line))