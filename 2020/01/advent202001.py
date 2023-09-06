def part1(numbers: list[int], target: int) -> int:
  n = len(numbers)
  left = 0
  right = n - 1
  while numbers[left] + numbers[right] != target:
    if numbers[left] + numbers[right] > target:
      right -= 1
    else:
      left += 1
  return numbers[left] * numbers[right]

def part2(numbers: list[int], target: int) -> int:
  n = len(numbers)
  for i in range(n - 2):
    left = i + 1
    right = n - 1
    while left < right:
      sum = numbers[i] + numbers[left] + numbers[right]
      if sum == target:
        return numbers[i] * numbers[left] * numbers[right]
      if sum > target:
        right -= 1
      else:
        left += 1
  return -1

if __name__ == "__main__":
  file1 = open('input.txt','r')
  lines = file1.readlines()

  numbers = [int(num) for num in lines]
  numbers.sort()
  target = 2020
  print(part1(numbers, target))
  print(part2(numbers, target))