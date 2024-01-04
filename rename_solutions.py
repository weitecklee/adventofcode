import os

for year in range(2015, 2024):
  for root, dirs, files in os.walk(f'{year}'):
    for file in files:
      if 'advent' in file:
        old_path = os.path.join(root, file)
        new_file = 'solution' + os.path.splitext(file)[1]
        new_path = os.path.join(root, new_file)
        os.rename(old_path, new_path)
        print(f'{old_path} renamed to {new_path}')