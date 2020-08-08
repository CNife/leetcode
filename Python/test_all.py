import os
import subprocess
import sys
from pathlib import Path

from rich import print

work_dir = Path(__file__).parent
os.environ['PYTHONPATH'] = str(work_dir)

test_files = []
for file in work_dir.joinpath('leetcode_python').glob('*.py'):
    if file.name != '__init__.py':
        test_files.append(file)
test_files.sort()

pass_count = 0
fail_count = 0
for file in test_files:
    print(f'Testing [cyan]{file.name}[/cyan]... ', end='')
    try:
        subprocess.run(
            ['python', str(file)],
            check=True,
            stdout=subprocess.DEVNULL,
            stderr=subprocess.DEVNULL,
        )
    except subprocess.CalledProcessError:
        print('[red bold]Failed[/red bold]')
        fail_count += 1
    else:
        print('[green bold]Passed[/green bold]')
        pass_count += 1

print(f'Test finished, '
      f'[green]{pass_count}[/green] passed, '
      f' [red]{fail_count}[/red] failed')
if fail_count > 0:
    sys.exit(1)
