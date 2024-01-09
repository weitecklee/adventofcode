import os

def findMarker(line: str, distincts: int) -> int:
  i = 0
  n = distincts
  letters: dict[str, int] = {}
  while i < n:
    c = line[i]
    if c in letters:
      n = max(n, letters[c] + distincts + 1)
    letters[c] = i
    i += 1
  return i

if __name__ == "__main__":
  file1 = open(os.path.join(os.path.dirname(os.path.abspath(__file__)), 'input.txt'),'r')
  line = file1.read()

  print(findMarker(line, 4))
  print(findMarker(line, 14))