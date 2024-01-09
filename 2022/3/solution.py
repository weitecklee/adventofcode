import os
file1 = open(os.path.join(os.path.dirname(os.path.abspath(__file__)), 'input.txt'),'r')
lines = file1.readlines()
sum = 0
prio = 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ'
set1 = set()
set2 = set()

for line in lines:
    set1.clear()
    set2.clear()
    l = int(len(line) / 2)
    for i in range(0, l):
        set1.add(line[i])
        set2.add(line[i + l])
    same = set1.intersection(set2)
    for x in same:
        sum += prio.index(x) + 1

print(sum)

set1.clear()
set2.clear()
set3 = set()
sum2 = 0
n = 0

for line in lines:
    if n == 0:
        for c in line.strip():
            set1.add(c)
        n += 1
    elif n == 1:
        for c in line.strip():
            set2.add(c)
        n += 1
    elif n == 2:
        for c in line.strip():
            set3.add(c)
        same = set1.intersection(set2, set3)
        for x in same:
            sum2 += prio.index(x) + 1
        set1.clear()
        set2.clear()
        set3.clear()
        n = 0

print(sum2)