import subprocess
import os

for year in range(2015, 2024):
  for root, dirs, files in os.walk(os.path.join('..', f'{year}')):
    for file in files:
      if 'solution' in file:
        solution_file = os.path.join(root, file)
        try:
          if 'js' in file:
            print(f'Running {solution_file}')
            result = subprocess.run(['node', solution_file], check=True, text=True, capture_output=True)
          # elif 'go' in file:
          #   result = subprocess.run(['go run', solution_file], check=True, text=True, capture_output=True)
          # elif 'py' in file:
          #   print(f'Running {solution_file}')
          #   result = subprocess.run(['py', solution_file], check=True, text=True, capture_output=True)
          else:
            continue

          stdout_output = result.stdout
          stderr_output = result.stderr

          if stdout_output:
              print('Standard Output:')
              print(stdout_output)

          if stderr_output:
              print('Standard Error:')
              print(stderr_output)

        except subprocess.CalledProcessError as e:
          print(f'An error occurred: {e}')