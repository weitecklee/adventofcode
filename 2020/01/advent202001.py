from typing import List

def part1(numbers: List[int], target: int) -> int:
  left, right = 0, len(numbers) - 1
  while numbers[left] + numbers[right] != target:
    if numbers[left] + numbers[right] > target:
      right -= 1
    else:
      left += 1
  return numbers[left] * numbers[right]

def part2(numbers: List[int], target: int) -> int:
  n = len(numbers)
  for i in range(n - 2):
    left, right = i + 1, n - 1
    while left < right:
      total = numbers[i] + numbers[left] + numbers[right]
      if total == target:
        return numbers[i] * numbers[left] * numbers[right]
      if total > target:
        right -= 1
      else:
        left += 1
  return -1

if __name__ == "__main__":
  with open('input.txt', 'r') as file:
    numbers = [int(num) for num in file]
  numbers.sort()
  target = 2020
  print(part1(numbers, target))
  print(part2(numbers, target))