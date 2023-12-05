# This python solution was generated by translating the Go solution through GPT-3.5
# This is purely used as a benchmark against the Go solution
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


def is_digit(char: str) -> bool:
    return char.isdigit()


def is_symbol(char: str) -> bool:
    return not (is_digit(char) or char == ".")


def find_complete_numbers_keyed(lines: list[str]) -> dict[str, str]:
    idx_num_map = {}

    for i, line in enumerate(lines):
        is_constructing_num = False
        constructed_num = ""
        num_idxs = []

        for j, ch in enumerate(line):
            is_num = is_digit(ch)

            if is_constructing_num:
                if is_num:
                    num_idxs.append(f"{i},{j}")
                    constructed_num += ch
                else:
                    keyed_num = f"{constructed_num};key{i},{j}"
                    idx_num_map.update({idx: keyed_num for idx in num_idxs})
                    is_constructing_num = False
                    num_idxs.clear()
                    constructed_num = ""
            else:
                if is_num:
                    num_idxs.append(f"{i},{j}")
                    constructed_num += ch
                    is_constructing_num = True

            if j == len(line) - 1 and is_constructing_num:
                keyed_num = f"{constructed_num};key{i},{j}"
                idx_num_map.update({idx: keyed_num for idx in num_idxs})
                is_constructing_num = False
                constructed_num = ""

    return idx_num_map


def unlock_complete_number(keyed_num: str) -> int:
    """
    Extract and return the unlocked number from a keyed number.
    """
    parts = keyed_num.split(";key")
    if len(parts) != 2 or not parts[0].isnumeric():
        raise ValueError(f"Invalid keyed number format: {keyed_num}")
    return int(parts[0])


def delete_all_keys_for_keyed_number(num_idx_map: dict[str, str], keyed_num: str):
    """
    Delete all keys associated with the given keyed number from the map.
    """
    keys_to_del = [key for key, val in num_idx_map.items() if keyed_num == val]
    for k in keys_to_del:
        del num_idx_map[k]


def run_p1(lines: list[str]) -> int:
    """
    Run part 1 of the program.
    """
    return run(lines, 1)


def run_p2(lines: list[str]) -> int:
    """
    Run part 2 of the program.
    """
    return run(lines, 2)


def run(lines: list[str], part: int) -> int:
    """
    Run the specified part of the program.
    """
    num_idx_map = find_complete_numbers_keyed(lines)
    total_sum = 0

    for i, line in enumerate(lines):
        for j, ch in enumerate(line.strip("\n")):
            if is_symbol(ch):
                neighbors = [
                    f"{i},{j-1}",
                    f"{i},{j+1}",
                    f"{i-1},{j}",
                    f"{i+1},{j}",
                    f"{i-1},{j-1}",
                    f"{i-1},{j+1}",
                    f"{i+1},{j-1}",
                    f"{i+1},{j+1}",
                ]
                keyed_num_set = set()

                for neighbor in neighbors:
                    keyed_num = num_idx_map.get(neighbor, None)

                    if keyed_num:
                        keyed_num_set.add(keyed_num)
                        delete_all_keys_for_keyed_number(num_idx_map, keyed_num)

                if part == 1:
                    total_sum += sum(unlock_complete_number(x) for x in keyed_num_set)
                elif part == 2 and len(keyed_num_set) == 2:
                    p1, p2 = map(unlock_complete_number, keyed_num_set)
                    total_sum += p1 * p2

    return total_sum


if __name__ == "__main__":
    input_file = parse_args()
    with open(input_file) as file:
        lines = file.readlines()
        print(f"part 1: {run_p1(lines)}")
        print(f"part 2: {run_p2(lines)}")