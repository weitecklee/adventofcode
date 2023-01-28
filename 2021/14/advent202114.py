import re

file1 = open('input202114.txt','r')
lines = file1.readlines()

instructions = {}

for line in lines[2:]:
  letters = re.findall('\w+', line)
  instructions[letters[0]] = letters[1]

polymer = lines[0].strip()

letters = {}
for char in polymer:
  if char in letters:
    letters[char] += 1
  else:
    letters[char] = 1

pairs = {}
for i in range(len(polymer) - 1):
  pair = polymer[i:i + 2]
  if pair in pairs:
    pairs[pair] += 1
  else:
    pairs[pair] = 1

def polymerize(pairs):
  pairs2 = {}
  for pair, count in pairs.items():
    if pair in instructions:
      leftPair = pair[0] + instructions[pair]
      rightPair = instructions[pair] + pair[1]
      if leftPair in pairs2:
        pairs2[leftPair] += count
      else:
        pairs2[leftPair] = count
      if rightPair in pairs2:
        pairs2[rightPair] += count
      else:
        pairs2[rightPair] = count
      if instructions[pair] in letters:
        letters[instructions[pair]] += count
      else:
        letters[instructions[pair]] = count
    else:
      pairs2[pair] = count
  return pairs2

for i in range(10):
  pairs = polymerize(pairs)

print(max(letters.values()) - min(letters.values()))

for i in range(30):
  pairs = polymerize(pairs)

print(max(letters.values()) - min(letters.values()))