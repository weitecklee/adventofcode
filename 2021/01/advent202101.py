file1 = open('input.txt','r')
lines = file1.readlines()
lines = [int(line) for line in lines]

count = 0
for i in range(1, len(lines)):
  if lines[i] > lines[i - 1]:
    count += 1

print(count)

count2 = 0
for i in range(3, len(lines)):
  if lines[i] > lines[i - 3]:
    count2 += 1

print(count2)