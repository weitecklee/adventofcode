import os, glob, re

def main():
  for file in glob.glob("**/input*", recursive = True):
    os.rename(file, re.sub('input\d+', 'input', file))

if __name__ == "__main__":
  main()