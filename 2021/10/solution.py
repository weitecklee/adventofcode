import os

file1 = open(os.path.join(os.path.dirname(os.path.abspath(__file__)), 'input.txt'),'r')
lines = file1.readlines()

brackets = {'{': '}', '(': ')', '<': '>', '[': ']'}
corruptedScores = {')': 3, ']': 57, '}': 1197, '>': 25137}
corruptedScore = 0
incompleteScores = {')': 1, ']': 2, '}': 3, '>': 4}
scores = []

for line in lines:
  queue = []
  incomplete = True
  for bracket in line.strip():
    if bracket in brackets:
      queue.append(bracket)
    elif len(queue) == 0 or brackets[queue.pop()] != bracket:
      corruptedScore += corruptedScores[bracket]
      incomplete = False
      break
  if incomplete:
    incompleteScore = 0
    while len(queue):
      incompleteScore *= 5
      incompleteScore += incompleteScores[brackets[queue.pop()]]
    scores.append(incompleteScore)

scores.sort()

print(corruptedScore)
print(scores[len(scores) // 2])

