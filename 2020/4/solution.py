import re
import os

REQUIRED_FIELDS = {'byr', 'iyr', 'eyr', 'hgt', 'hcl', 'ecl', 'pid'}
VALID_ECL = ('amb', 'blu', 'brn', 'gry', 'grn', 'hzl', 'oth')

class Passport:
  def __init__(self) -> None:
    self.fields = {}

  def add_field(self, field: str) -> None:
    k, v = field.split(':')
    self.fields[k] = v

  def validate(self, part2 = False) -> bool:
    if not REQUIRED_FIELDS.issubset(self.fields.keys()):
      return False
    if part2:
      if not re.match(r'^\d{4}$', self.fields['byr']):
        return False
      if not 1920 <= int(self.fields['byr']) <= 2002:
        return False
      if not re.match(r'^\d{4}$', self.fields['iyr']):
        return False
      if not 2010 <= int(self.fields['iyr']) <= 2020:
        return False
      if not re.match(r'^\d{4}$', self.fields['eyr']):
        return False
      if not 2020 <= int(self.fields['eyr']) <= 2030:
        return False
      hgt_match = re.match(r'^(\d+)(cm|in)$', self.fields['hgt'])
      if not hgt_match:
        return False
      hgt, unit = int(hgt_match.group(1)), hgt_match.group(2)
      if unit == 'cm' and not 150 <= hgt <= 193:
        return False
      if unit == 'in' and not 59 <= hgt <= 76:
        return False
      if not re.match(r'^#[0-9a-f]{6}$', self.fields['hcl']):
        return False
      if self.fields['ecl'] not in VALID_ECL:
        return False
      if not re.match(r'^\d{9}$', self.fields['pid']):
        return False
    return True

if __name__ == "__main__":
  with open(os.path.join(os.path.dirname(os.path.abspath(__file__)), 'input.txt'),'r') as file:
    lines = [line.strip() for line in file]

  pattern = r'(\w{3}:\S+)'

  passports: list[Passport] = []
  passport = Passport()
  for line in lines:
    if not line:
      passports.append(passport)
      passport = Passport()
    else:
      matches = re.findall(pattern, line)
      for match in matches:
        passport.add_field(match)

  part1 = sum([passport.validate() for passport in passports])
  print(part1)
  part2 = sum([passport.validate(True) for passport in passports])
  print(part2)