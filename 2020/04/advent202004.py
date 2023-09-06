import re

class Passport:
  def __init__(self) -> None:
    self.fields = dict()

  def addField(self, field: str):
    a, b = field.split(':')
    self.fields[a] = b

  def valid1(self) -> bool:
    if 'byr' not in self.fields:
      return False
    if 'iyr' not in self.fields:
      return False
    if 'eyr' not in self.fields:
      return False
    if 'hgt' not in self.fields:
      return False
    if 'hcl' not in self.fields:
      return False
    if 'ecl' not in self.fields:
      return False
    if 'pid' not in self.fields:
      return False
    return True

  def valid2(self) -> bool:
    if 'byr' not in self.fields:
      return False
    if re.match(r'^\d{4}$', self.fields['byr']) == None:
      return False
    byr = int(self.fields['byr'])
    if not 1920 <= byr <= 2002:
      return False
    if 'iyr' not in self.fields:
      return False
    if re.match(r'^\d{4}$', self.fields['iyr']) == None:
      return False
    iyr = int(self.fields['iyr'])
    if not 2010 <= iyr <= 2020:
      return False
    if 'eyr' not in self.fields:
      return False
    if re.match(r'^\d{4}$', self.fields['eyr']) == None:
      return False
    eyr = int(self.fields['eyr'])
    if not 2020 <= eyr <= 2030:
      return False
    if 'hgt' not in self.fields:
      return False
    match = re.match(r'^(\d+)(cm|in)$', self.fields['hgt'])
    if match == None:
      return False
    hgt = int(match.group(1))
    if match.group(2) == 'cm' and not 150 <= hgt <= 193:
      return False
    if match.group(2) == 'in' and not 59 <= hgt <= 76:
      return False
    if 'hcl' not in self.fields:
      return False
    if re.match(r'^#[0-9a-f]{6}$', self.fields['hcl']) == None:
      return False
    if 'ecl' not in self.fields:
      return False
    if self.fields['ecl'] not in ecl:
      return False
    if 'pid' not in self.fields:
      return False
    if re.match(r'^\d{9}$', self.fields['pid']) == None:
      return False
    return True

if __name__ == "__main__":
  file1 = open('input.txt','r')
  lines = [line.strip() for line in file1.readlines()]

  pattern = r'(\w{3}:\S+)'

  passports: list[Passport] = []
  passport = Passport()
  for line in lines:
    if len(line) == 0:
      passports.append(passport)
      passport = Passport()
    else:
      matches = re.findall(pattern, line)
      for match in matches:
        passport.addField(match)

  ecl = ('amb', 'blu', 'brn', 'gry', 'grn', 'hzl', 'oth')

  part1 = sum([passport.valid1() for passport in passports])
  print(part1)
  part2 = sum([passport.valid2() for passport in passports])
  print(part2)