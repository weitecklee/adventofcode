def part1(numbers: list[int]) -> int:
  return sum(numbers)

def part2(numbers: list[int]) -> int:
  numberSet = set()
  res = 0
  while True:
    for number in numbers:
      res += number
      if res in numberSet:
        return res
      numberSet.add(res)
  return 0

if __name__ == "__main__":
  file1 = open('input.txt','r')
  lines = file1.readlines()

  numbers = [int(num) for num in lines]
  print(part1(numbers))
  print(part2(numbers))