import os

def part1(numbers: list[int]) -> int:
  return sum([1 for a, b in zip(numbers, numbers[1:]) if b > a])

def part2(numbers: list[int]) -> int:
  return part1([sum(numbers[i-2:i+1]) for i in range(2, len(numbers))])

if __name__ == "__main__":
  with open(os.path.join(os.path.dirname(os.path.abspath(__file__)), 'input.txt'),'r') as file:
    numbers = [int(num) for num in file]

  print(part1(numbers))
  print(part2(numbers))