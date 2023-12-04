import argparse, os


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


def run_p1():
    pass


def run_p2():
    pass


if __name__ == "__main__":
    raise NotImplementedError("Only a Go solution was used for this day")
