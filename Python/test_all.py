import os
import subprocess
import sys

from rich import print

base_dir = 'leetcode_python'

test_files = []
for file in os.listdir(base_dir):
    if file.endswith('.py') and file != '__init__.py':
        test_files.append(file)
test_files.sort()

pass_count = 0
fail_count = 0
for file in test_files:
    print(f'Testing [cyan]{file}[/cyan]... ', end='')
    try:
        real_path = os.path.join(base_dir, file)
        result = subprocess.run(['python', real_path],
                                check=True,
                                stdout=subprocess.DEVNULL,
                                stderr=subprocess.DEVNULL)
    except subprocess.CalledProcessError:
        print('[red bold]Failed[/red bold]')
        fail_count += 1
    else:
        print('[green bold]Passed[/green bold]')
        pass_count += 1

print(f'Test finished, '
      f'[green]{pass_count}[/green] passed,'
      f' [red]{fail_count}[/red] failed')
if fail_count > 0:
    sys.exit(1)
