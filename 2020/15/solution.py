import os
from typing import Dict, List

def playGame(numbers: List[int], turns: int) -> int:
  history: Dict[int, int] = {}
  for i, n in enumerate(numbers):
    history[n] = i + 1
  nextNumber = 0
  for i in range(len(numbers), turns - 1):
    if nextNumber in history:
      temp = i - history[nextNumber] + 1
    else:
      temp = 0
    history[nextNumber] = i + 1
    nextNumber = temp
  return nextNumber

if __name__ == "__main__":
  with open(os.path.join(os.path.dirname(os.path.abspath(__file__)), 'input.txt'),'r') as file:
    line = file.readline()
  numbers = [int(n) for n in line.split(',')]
  print(playGame(numbers, 2020))
  print(playGame(numbers, 30000000))
