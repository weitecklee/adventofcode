import re

file1 = open('input.txt','r')
lines = file1.readlines()
lines = list(map(str.strip, lines))
monkeys = []

class Monkey:
  def __init__(self, items, operation, testDiv, trueMonkey, falseMonkey):
    self.items = items
    self.operation = operation
    self.testDiv = testDiv
    self.trueMonkey = trueMonkey
    self.falseMonkey = falseMonkey
    self.inspections = 0
  def test(self, item):
    if item % self.testDiv:
      return self.falseMonkey
    else:
      return self.trueMonkey

reliever = 1

for i in range(0, len(lines), 7):
  items = list(map(int, re.findall('\d+', lines[i + 1])))
  line2 = lines[i + 2].split(' ')
  if line2[-1] == 'old':
    operation = lambda x: x * x
  else:
    num = int(line2[-1])
    if line2[-2] == '+':
      operation = lambda x, num = num: x + num
    else:
      operation = lambda x, num = num: x * num
  testDiv = int(lines[i + 3].split(' ')[-1])
  reliever *= testDiv
  trueMonkey = int(lines[i + 4].split(' ')[-1])
  falseMonkey = int(lines[i + 5].split(' ')[-1])
  monkey = Monkey(items, operation, testDiv, trueMonkey, falseMonkey)
  monkeys.append(monkey)

def relief(item):
  # return item // 3
  return item % reliever

# for round in range(20):
#   for i, monkey in enumerate(monkeys):
#     monkey.inspections += len(monkey.items)
#     for item in monkey.items:
#       item = relief(monkey.operation(item))
#       monkeys[monkey.test(item)].items.append(item)
#     monkey.items = []

# inspections = [monkey.inspections for monkey in monkeys]
# inspections.sort()
# print(inspections[-2] * inspections[-1])

for round in range(10000):
  for i, monkey in enumerate(monkeys):
    monkey.inspections += len(monkey.items)
    for item in monkey.items:
      item = relief(monkey.operation(item))
      monkeys[monkey.test(item)].items.append(item)
    monkey.items = []

inspections = [monkey.inspections for monkey in monkeys]
inspections.sort()
print(inspections[-2] * inspections[-1])
