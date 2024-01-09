import os
from typing import List, Set, Tuple

class Instruction:
  def __init__(self, action: str, val: int):
    self.action = action
    self.val = val

  def execute(self, nLine: int, acc: int) -> Tuple[int, int]:
    if self.action == 'acc':
      acc += self.val
    elif self.action == 'jmp':
      nLine += self.val - 1
    nLine += 1
    return nLine, acc

  def execute2(self, nLine: int, acc: int) -> Tuple[int, int]:
    if self.action == 'acc':
      acc += self.val
    elif self.action == 'nop':
      nLine += self.val - 1
    nLine += 1
    return nLine, acc

def parse(lines: List[str]) -> List[Instruction]:
  instructions: List[Instruction] = []
  for line in lines:
    action, val = line.split(' ')
    instructions.append(Instruction(action, int(val)))
  return instructions

def runProgram(instructions: List[Instruction], lineToSkip: int) -> Tuple[int, List[int], bool]:
  nLine = 0
  acc = 0
  executed: Set[int] = set()
  linesToSkip: List[int] = []
  while nLine < len(instructions) and nLine not in executed:
    executed.add(nLine)
    if instructions[nLine].action in ['jmp', 'nop']:
      linesToSkip.append(nLine)
    if nLine == lineToSkip:
      nLine, acc = instructions[nLine].execute2(nLine, acc)
    else:
      nLine, acc = instructions[nLine].execute(nLine, acc)
  return acc, linesToSkip, nLine >= len(instructions)

def part2(instructions: List[Instruction], linesToSkip: List[int]) -> int:
  for lineToSkip in linesToSkip:
    acc, _, success = runProgram(instructions, lineToSkip)
    if success:
      return acc
  return -1

if __name__ == "__main__":
  with open(os.path.join(os.path.dirname(os.path.abspath(__file__)), 'input.txt'),'r') as file:
    lines = [line.strip() for line in file]
  instructions = parse(lines)
  part1, linesToSkip, _ = runProgram(instructions, -1)
  print(part1)
  print(part2(instructions, linesToSkip))
