import requests
import os
import browser_cookie3

cookies = browser_cookie3.chrome()
headers = {'User-Agent': 'github.com/weitecklee/adventofcode'}

for year in range(2015, 2024):
  for day in range(1,26):
    r = requests.get(f"https://adventofcode.com/{year}/day/{day}/input", cookies = cookies, headers = headers)
    if not os.path.exists(f"{year}/{day}"):
      os.mkdir(f"{year}/{day}")
    with open(f"{year}/{day}/input.txt", 'w') as file:
      file.write(r.text[:-1]) # omit final blank line