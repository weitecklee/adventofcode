import os
from typing import List, Dict

WIN = 6
DRAW = 3
LOSS = 0
ROCK = 1
PAPER = 2
SCISSORS = 3
HANDS: Dict[str, int] = {
  'A': ROCK,
  'B': PAPER,
  'C': SCISSORS,
  'X': ROCK,
  'Y': PAPER,
  'Z': SCISSORS
}
RESULTS: Dict[str, int] = {
  'X': LOSS,
  'Y': DRAW,
  'Z': WIN
}

def parse(puzzle_input: List[str]) -> List[List[str]]:
  rounds: List[List[str]] = []
  for line in puzzle_input:
    rounds.append(line.split(' '))
  return rounds

def determine_result(hand1: str, hand2: str) -> int:
  rps1 = HANDS[hand1]
  rps2 = HANDS[hand2]
  if rps1 == rps2:
    return DRAW
  if (rps1 == ROCK and rps2 == PAPER) or (rps1 == PAPER and rps2 == SCISSORS) or (rps1 == SCISSORS and rps2 == ROCK):
    return WIN
  return LOSS

def part1(rounds: List[List[str]]) -> int:
  score = 0
  for round in rounds:
    hand1, hand2 = round
    score += HANDS[hand2] + determine_result(hand1, hand2)
  return score

def part2(rounds: List[List[str]]) -> int:
  score = 0
  for round in rounds:
    rps1 = HANDS[round[0]]
    res = RESULTS[round[1]]
    score += res
    if res == LOSS:
      if rps1 == ROCK:
        score += SCISSORS
      elif rps1 == SCISSORS:
        score += PAPER
      else:
        score += ROCK
    elif res == DRAW:
      score += rps1
    else:
      if rps1 == ROCK:
        score += PAPER
      elif rps1 == SCISSORS:
        score += ROCK
      else:
        score += SCISSORS
  return score

if __name__ == '__main__':
  with open(os.path.join(os.path.dirname(os.path.abspath(__file__)), 'input.txt'),'r') as file:
    puzzle_input = [line.strip() for line in file]
  rounds = parse(puzzle_input)
  print(part1(rounds))
  print(part2(rounds))