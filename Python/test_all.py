import os
import subprocess
import sys
from pathlib import Path

from rich import print

work_dir = Path(__file__).parent
os.environ["PYTHONPATH"] = str(work_dir)

test_files = []
for file in work_dir.joinpath("code").glob("*.py"):
    if file.name != "__init__.py":
        test_files.append(file)
test_files.sort()

pass_count = 0
fail_count = 0
for file in test_files:
    print(f"Testing [cyan]{file.name}[/]... ", end="")
    try:
        subprocess.run(
            ["python", str(file)],
            check=True,
            stdout=subprocess.DEVNULL,
            stderr=subprocess.DEVNULL,
        )
    except subprocess.CalledProcessError:
        print("[red bold]Failed[/]")
        fail_count += 1
    else:
        print("[green bold]Passed[/]")
        pass_count += 1

print(
    f"Test finished, "
    f"[green]{pass_count}[/] passed, "
    f" [red]{fail_count}[/] failed"
)
if fail_count > 0:
    sys.exit(1)
