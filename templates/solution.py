import os

if __name__ == '__main__':
  with open(os.path.join(os.path.dirname(os.path.abspath(__file__)), 'input.txt'),'r') as file:
    puzzle_input = file.read().strip().split('\n')