file1 = open('input202201.txt','r')
lines = file1.readlines()
count = 0
maxC = 0
arr = []

for line in lines:
    if line == '\n':
        arr.append(count)
        if count > maxC:
            maxC = count
        count = 0
    else:
        count += int(line)

print('Maximum: ' + str(maxC))

arr.sort(reverse = True)

sum = 0
for i in range(0, 3):
    sum += arr[i]

print(sum)