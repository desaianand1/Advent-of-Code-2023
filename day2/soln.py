# This python solution was generated by translating the Go solution through GPT-3.5
# This is purely used as a benchmark against the Go solution
import argparse, os
from enum import Enum
import sys


class Color(str, Enum):
    Red = "red"
    Green = "green"
    Blue = "blue"


class Game:
    def __init__(self, _id, sets):
        self.id = _id
        self.sets = sets


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
    try:
        with open(path, "r") as input_file:
            file_lines = input_file.read().splitlines()
    except FileNotFoundError as e:
        print(e)
        sys.exit(1)

    return file_lines


def parse_games(lines):
    games = []
    for line in lines:
        sections = line.split(":")
        id_str = sections[0].split(" ")[1]
        game_id = int(id_str)

        cube_sets = []
        sets = sections[1].split(";")
        for _set in sets:
            cubes = _set.split(",")
            cube_map = {}
            for cube in cubes:
                cube = cube.strip()
                pair = cube.split(" ")
                color = Color(pair[1])
                count = int(pair[0])
                cube_map[color] = count
            cube_sets.append(cube_map)
        games.append(Game(game_id, cube_sets))
    return games


def run_p1(games):
    possible_cubes = {Color.Red: 12, Color.Green: 13, Color.Blue: 14}
    _sum = 0
    for game in games:
        is_possible = all(
            possible_cubes.get(color, 0) >= count
            for cube in game.sets
            for color, count in cube.items()
        )
        if is_possible:
            _sum += game.id
    return _sum


def run_p2(games):
    _sum = 0
    for game in games:
        least_cubes_required = {Color.Red: 0, Color.Green: 0, Color.Blue: 0}
        for cube in game.sets:
            for color, count in cube.items():
                if least_cubes_required.get(color, 0) < count:
                    least_cubes_required[color] = count
        power = 1
        for count in least_cubes_required.values():
            power *= count
        _sum += power
    return _sum


if __name__ == "__main__":
    lines = parse_args()
    games = parse_games(lines)
    print(f"part 1: {run_p1(games)}")
    print(f"part 2: {run_p2(games)}")