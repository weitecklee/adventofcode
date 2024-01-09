import re
import os

file1 = open(os.path.join(os.path.dirname(os.path.abspath(__file__)), 'input.txt'),'r')
lines = file1.readlines()
n = 0

for line in lines:
    nums = list(map(int, re.findall('\d+', line)))
    if (nums[0] >= nums[2] and nums[1] <= nums[3]) or (nums[0] <= nums[2] and nums[1] >= nums[3]) :
        n += 1

print(n)

n2 = 0

for line in lines:
    nums = list(map(int, re.findall('\d+', line)))
    if nums[0] <= nums[2] <= nums[1] or nums[0] <= nums[3] <= nums[1] or nums[2] <= nums[0] <= nums[3] or nums[2] <= nums[1] <= nums[3]:
        n2 += 1

print(n2)