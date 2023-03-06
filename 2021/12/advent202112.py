file1 = open('input.txt','r')
lines = file1.readlines()

from collections import defaultdict, deque

connections = defaultdict(list)

for line in lines:
  a, b = line.strip().split('-')
  connections[a].append(b)
  connections[b].append(a)

q = deque([['start', set(['start'])]])
res = 0

while len(q):
  cave, visited = q.popleft()
  for connection in connections[cave]:
    if connection == 'end':
      res += 1
      continue
    if connection not in visited:
      newVisited = set(visited)
      if connection.islower():
        newVisited.add(connection)
      q.append([connection, newVisited])

print(res)

q = deque([['start', set(['start']), '']])
res = 0

while len(q):
  cave, visited, twice = q.popleft()
  for connection in connections[cave]:
    if connection == 'end':
      res += 1
      continue
    if connection not in visited:
      newVisited = set(visited)
      if connection.islower():
        newVisited.add(connection)
      q.append([connection, newVisited, twice])
    elif connection != 'start' and twice == '':
      q.append([connection, visited, connection])

print(res)