file1 = open('input202212.txt','r')
lines = file1.readlines()
lines = list(map(str.strip, lines))

visited1 = []
visited2 = []
alphabet = 'abcdefghijklmnopqrstuvwxyz'

for i, line in enumerate(lines):
  line = list(line)
  for j, letter in enumerate(line):
    if letter == 'S':
      start = [i, j]
      line[j] = 0
    elif letter == 'E':
      end = [i, j]
      line[j] = 25
    else:
      line[j] = alphabet.index(letter)
  visited1.append(line.copy())
  visited2.append(line.copy())

q = [[start, 1, 0]]
visited1[start[0]][start[1]] = 'X'

for item in q:
  pos, ht, steps = item
  if pos[0] == end[0] and pos[1] == end[1]:
    print(steps)
    break
  if pos[0] > 0 and visited1[pos[0] - 1][pos[1]] != 'X' and visited1[pos[0] - 1][pos[1]] <= ht + 1:
    q.append([[pos[0] - 1, pos[1]], visited1[pos[0] - 1][pos[1]], steps + 1])
    visited1[pos[0] - 1][pos[1]] = 'X'
  if pos[1] > 0 and visited1[pos[0]][pos[1] - 1] != 'X' and visited1[pos[0]][pos[1] - 1] <= ht + 1:
    q.append([[pos[0], pos[1] - 1], visited1[pos[0]][pos[1] - 1], steps + 1])
    visited1[pos[0]][pos[1] - 1] = 'X'
  if pos[0] < len(visited1) - 1 and visited1[pos[0] + 1][pos[1]] != 'X' and visited1[pos[0] + 1][pos[1]] <= ht + 1:
    q.append([[pos[0] + 1, pos[1]], visited1[pos[0] + 1][pos[1]], steps + 1])
    visited1[pos[0] + 1][pos[1]] = 'X'
  if pos[1] < len(visited1[0]) - 1 and visited1[pos[0]][pos[1] + 1] != 'X' and visited1[pos[0]][pos[1] + 1] <= ht + 1:
    q.append([[pos[0], pos[1] + 1], visited1[pos[0]][pos[1] + 1], steps + 1])
    visited1[pos[0]][pos[1] + 1] = 'X'

q = [[end, 25, 0]]
visited2[end[0]][end[1]] = 'X'

for item in q:
  pos, ht, steps = item
  if ht == 0:
    print(steps)
    break
  if pos[0] > 0 and visited2[pos[0] - 1][pos[1]] != 'X' and visited2[pos[0] - 1][pos[1]] >= ht - 1:
    q.append([[pos[0] - 1, pos[1]], visited2[pos[0] - 1][pos[1]], steps + 1])
    visited2[pos[0] - 1][pos[1]] = 'X'
  if pos[1] > 0 and visited2[pos[0]][pos[1] - 1] != 'X' and visited2[pos[0]][pos[1] - 1] >= ht - 1:
    q.append([[pos[0], pos[1] - 1], visited2[pos[0]][pos[1] - 1], steps + 1])
    visited2[pos[0]][pos[1] - 1] = 'X'
  if pos[0] < len(visited2) - 1 and visited2[pos[0] + 1][pos[1]] != 'X' and visited2[pos[0] + 1][pos[1]] >= ht - 1:
    q.append([[pos[0] + 1, pos[1]], visited2[pos[0] + 1][pos[1]], steps + 1])
    visited2[pos[0] + 1][pos[1]] = 'X'
  if pos[1] < len(visited2[0]) - 1 and visited2[pos[0]][pos[1] + 1] != 'X' and visited2[pos[0]][pos[1] + 1] >= ht - 1:
    q.append([[pos[0], pos[1] + 1], visited2[pos[0]][pos[1] + 1], steps + 1])
    visited2[pos[0]][pos[1] + 1] = 'X'



