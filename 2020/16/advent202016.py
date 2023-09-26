import re
from typing import List, Set, Tuple
from functools import reduce

class Field:
  def __init__(self, name: str, lo1: int, hi1: int, lo2: int, hi2: int):
    self.name: str = name
    self.lo1: int = lo1
    self.hi1: int = hi1
    self.lo2: int = lo2
    self.hi2: int = hi2
    self.possibilities: Set[int] = set(range(20))
    self.position: int = -1

  def remove_possibility(self, n: int):
    self.possibilities.remove(n)

  def is_possible(self, n: int) -> bool:
    return self.lo1 <= n <= self.hi1 or self.lo2 <= n <= self.hi2

  def set_position(self):
    self.position = self.possibilities.pop()

def parse(lines: List[str]) -> Tuple[List[Field], List[int], List[List[int]]]:
  pattern = r'(.*): (\d+)-(\d+) or (\d+)-(\d+)'
  i = 0
  fields: List[Field] = []
  while lines[i]:
    match = re.match(pattern, lines[i])
    if match:
      fields.append(Field(match.group(1), int(match.group(2)), int(match.group(3)), int(match.group(4)), int(match.group(5))))
    i += 1
  i += 2
  my_ticket = [int(num) for num in lines[i].split(',')]
  i += 3
  nearby_tickets: List[List[int]] = []
  for j in range(i, len(lines)):
    nearby_tickets.append([int(num) for num in lines[j].split(',')])
  return fields, my_ticket, nearby_tickets

def part1(fields: List[Field], nearby_tickets: List[List[int]]) -> Tuple[int, List[List[int]]]:
  valid_tickets: List[List[int]] = []
  total = 0
  for ticket in nearby_tickets:
    ticket_is_valid = True
    for n in ticket:
      field_is_valid = False
      for field in fields:
        if field.is_possible(n):
          field_is_valid = True
          break
      if not field_is_valid:
        ticket_is_valid = False
        total += n
    if ticket_is_valid:
      valid_tickets.append(ticket)
  return total, valid_tickets

def part2(fields: List[Field], my_ticket: List[int], valid_tickets: List[List[int]]) -> int:
  for ticket in valid_tickets:
    for i, n in enumerate(ticket):
      for field in fields:
        if not field.is_possible(n):
          field.remove_possibility(i)
  departure_positions: List[int] = []
  while fields:
    determined_field = None
    for i, field in enumerate(fields):
      if len(field.possibilities) == 1:
        determined_field = fields.pop(i)
        break
    if determined_field:
      determined_field.set_position()
      determined_position = determined_field.position
      for field in fields:
        field.remove_possibility(determined_position)
      if determined_field.name.startswith('departure'):
        departure_positions.append(determined_position)
    else:
      raise Exception('No determined_field found')
  return reduce(lambda x, y: x * y, [my_ticket[pos] for pos in departure_positions])

if __name__ == "__main__":
  with open('input.txt','r') as file:
    lines = [line.strip() for line in file]
  fields, my_ticket, nearby_tickets = parse(lines)
  total, valid_tickets = part1(fields, nearby_tickets)
  print(total)
  print(part2(fields, my_ticket, valid_tickets))