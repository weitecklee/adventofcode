import os

def part1(numbers: list[int]) -> int:
  return sum(numbers)

def part2(numbers: list[int]) -> int:
  number_set = set()
  res = 0
  while True:
    for number in numbers:
      res += number
      if res in number_set:
        return res
      number_set.add(res)
  return 0

if __name__ == "__main__":
  with open(os.path.join(os.path.dirname(os.path.abspath(__file__)), 'input.txt'),'r') as file:
    numbers = [int(num) for num in file]

  print(part1(numbers))
  print(part2(numbers))