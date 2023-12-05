import argparse, os, re


def extract_digits_p1(line: str):
    first = None
    lastIdx = -1
    for i, el in enumerate(line):
        if el.isnumeric():
            if first == None:
                first = el
            else:
                lastIdx = i
    if first is None:
        return ""
    if lastIdx == -1:
        return first + first
    return first + line[lastIdx]


def extract_digits_p2(line: str):
    numMap = {
        "one": 1,
        "two": 2,
        "three": 3,
        "four": 4,
        "five": 5,
        "six": 6,
        "seven": 7,
        "eight": 8,
        "nine": 9,
    }
    pattern = r"(?=(one|two|three|four|five|six|seven|eight|nine|[1-9]))"
    matches = [match.group(1) for match in re.finditer(pattern, line)]
    first, second = matches[0], matches[-1]
    if first in numMap:
        first = numMap[first]
    if second in numMap:
        second = numMap[second]
    return f"{first}{second}"


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


def run_p1(lines: list[str]) -> int:
    sum = 0
    for line in lines:
        digits = extract_digits_p1(line)
        if digits != "":
            sum = sum + int(digits)
    return sum


def run_p2(lines: list[str]) -> int:
    sum = 0
    for line in lines:
        digits = extract_digits_p2(line)
        if digits != "":
            sum = sum + int(digits)
    return sum


if __name__ == "__main__":
    input = parse_args()
    with open(input) as file:
        lines = file.readlines()
        print(f"part 1: {run_p1(lines)}")
        print(f"part 1: {run_p2(lines)}")
