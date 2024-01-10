# AoC Template Python file
import argparse, os
from typing import List


def parse_args() -> str:
    parser = argparse.ArgumentParser()
    parser.add_argument(
        "input",
        help="input file (.txt) to be read",
        type=str,
        nargs="?",
        default="input.txt",
    )
    args = parser.parse_args()
    input_dir = os.path.dirname(os.path.realpath(__file__))
    path = os.path.join(input_dir, args.input)
    if os.path.isfile(path):
        return path
    else:
        raise FileNotFoundError(f"Input file {path} not found!")


def run_p1(lines: List[str]) -> int:
    pass


def run_p2(lines: List[str]) -> int:
    pass


if __name__ == "__main__":
    input = parse_args()
    with open(input) as file:
        lines = file.readlines()
        print(f"part 1: {run_p1(lines)}")
        print(f"part 2: {run_p2(lines)}")
