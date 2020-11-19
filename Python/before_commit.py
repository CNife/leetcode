import os
import subprocess
import sys
from pathlib import Path
from typing import List

from rich import print


def collect_test_files(directory: Path) -> List[Path]:
    result = []
    for file in directory.glob("*.py"):
        if file.name != "__init__.py":
            result.append(file)
    result.sort()
    return result


def run_test(tests: List[Path]) -> None:
    pass_count = 0
    fail_count = 0
    for test in tests:
        print(f"Testing [cyan]{test.name}[/]... ", end="")
        try:
            subprocess.run(
                ["python", str(test)],
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


def run_format(files: List[Path]) -> None:
    try:
        command = ["python", "-m", "black"]
        command.extend(str(file) for file in files)
        subprocess.run(command, check=True)
    except subprocess.CalledProcessError:
        print("[red bold]Format failed[/]")
        sys.exit(1)
    else:
        print("[green bold]Format finished[/]")


if __name__ == "__main__":
    work_dir = Path(__file__).parent
    os.environ["PYTHONPATH"] = str(work_dir)

    test_files = collect_test_files(work_dir / "code")
    run_test(test_files)
    run_format(test_files)
