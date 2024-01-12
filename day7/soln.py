# This python solution was generated by translating the Go solution through GPT-3.5
# This is purely used as a benchmark against the Go solution
import argparse
import os
from dataclasses import dataclass
from enum import Enum
from typing import Dict, List


class HandType(str, Enum):
    FIVE_OF_A_KIND = "FIVE_OF_A_KIND"
    FOUR_OF_A_KIND = "FOUR_OF_A_KIND"
    FULL_HOUSE = "FULL_HOUSE"
    THREE_OF_A_KIND = "THREE_OF_A_KIND"
    TWO_PAIR = "TWO_PAIR"
    ONE_PAIR = "ONE_PAIR"
    HIGH_CARD = "HIGH_CARD"


@dataclass
class CardHand:
    cards: str
    bid: int
    _type: HandType


card_rank_map = {
    "2": 2,
    "3": 3,
    "4": 4,
    "5": 5,
    "6": 6,
    "7": 7,
    "8": 8,
    "9": 9,
    "T": 10,
    "J": 11,
    "Q": 12,
    "K": 13,
    "A": 14,
}

card_hand_type_map: Dict[HandType, int] = {
    HandType.FIVE_OF_A_KIND: 7,
    HandType.FOUR_OF_A_KIND: 6,
    HandType.FULL_HOUSE: 5,
    HandType.THREE_OF_A_KIND: 4,
    HandType.TWO_PAIR: 3,
    HandType.ONE_PAIR: 2,
    HandType.HIGH_CARD: 1,
}


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
    path = os.path.join(os.path.dirname(os.path.realpath(__file__)), args.input)
    if os.path.isfile(path):
        return path
    else:
        raise FileNotFoundError(f"Input file {path} not found!")


def parse_int(value: str) -> int:
    try:
        return int(value)
    except ValueError:
        print(f"{value} is NOT a valid integer")
        exit(1)


def create_card_count_map(cards: str) -> Dict[str, int]:
    return {card: cards.count(card) for card in set(cards)}


def determine_hand_type(cards: str) -> Enum:
    card_map = create_card_count_map(cards)
    max_count = max(card_map.values(), default=0)
    FIVE_COUNT, FOUR_COUNT, THREE_COUNT, TWO_COUNT = 5, 4, 3, 2

    if max_count == FIVE_COUNT:
        return HandType.FIVE_OF_A_KIND
    elif max_count == FOUR_COUNT:
        return HandType.FOUR_OF_A_KIND
    elif max_count == THREE_COUNT:
        return (
            HandType.FULL_HOUSE
            if TWO_COUNT in card_map.values()
            else HandType.THREE_OF_A_KIND
        )
    elif max_count == TWO_COUNT:
        return (
            HandType.TWO_PAIR
            if list(card_map.values()).count(TWO_COUNT) == 2
            else HandType.ONE_PAIR
        )
    else:
        return HandType.HIGH_CARD


def determine_hand_type_p2(cards: str) -> Enum:
    card_map = create_card_count_map(cards)
    max_count = max(
        [count if card != "J" else 0 for card, count in card_map.items()], default=0
    )
    j_count = card_map.get("J", 0)
    high_count_card = max(card_map, key=card_map.get)

    FIVE_COUNT, FOUR_COUNT, THREE_COUNT, TWO_COUNT, HIGH = 5, 4, 3, 2, 1
    if max_count == FIVE_COUNT or j_count == FIVE_COUNT:
        return HandType.FIVE_OF_A_KIND
    elif max_count == FOUR_COUNT:
        if j_count == HIGH:
            return HandType.FIVE_OF_A_KIND
        return HandType.FOUR_OF_A_KIND
    elif max_count == THREE_COUNT:
        if j_count == TWO_COUNT:
            return HandType.FIVE_OF_A_KIND
        elif j_count == HIGH:
            return HandType.FOUR_OF_A_KIND
        for count in card_map.values():
            if count == TWO_COUNT:
                return HandType.FULL_HOUSE
        return HandType.THREE_OF_A_KIND
    elif max_count == TWO_COUNT:
        if j_count == THREE_COUNT:
            return HandType.FIVE_OF_A_KIND
        elif j_count == TWO_COUNT:
            return HandType.FOUR_OF_A_KIND
        elif j_count == HIGH:
            for card, count in card_map.items():
                if count == TWO_COUNT and high_count_card != card:
                    return HandType.FULL_HOUSE
            return HandType.THREE_OF_A_KIND
        for card, count in card_map.items():
            if count == TWO_COUNT and high_count_card != card:
                return HandType.TWO_PAIR
        return HandType.ONE_PAIR
    elif max_count == HIGH:
        if j_count == FOUR_COUNT:
            return HandType.FIVE_OF_A_KIND
        elif j_count == THREE_COUNT:
            return HandType.FOUR_OF_A_KIND
        elif j_count == TWO_COUNT:
            return HandType.THREE_OF_A_KIND
        elif j_count == HIGH:
            return HandType.ONE_PAIR
        return HandType.HIGH_CARD
    else:
        return HandType.HIGH_CARD


def parse_hands(lines: List[str], is_part_two: bool) -> List[CardHand]:
    return [
        CardHand(
            parts[0],
            parse_int(parts[1]),
            determine_hand_type_p2(parts[0])
            if is_part_two
            else determine_hand_type(parts[0]),
        )
        for line in lines
        if (parts := line.split())
    ]


def calculate_total_winnings(ranked_hands: List[CardHand]) -> int:
    return sum((i + 1) * hand.bid for i, hand in enumerate(ranked_hands))


def rank_hands(hands: List[CardHand]) -> List[CardHand]:
    return sorted(
        hands,
        key=lambda x: (
            card_hand_type_map[x._type],
            [card_rank_map[card] for card in x.cards],
        ),
    )


def run_p1(lines: List[str]) -> int:
    card_rank_map["J"] = 11
    return calculate_total_winnings(rank_hands(parse_hands(lines, False)))


def run_p2(lines: List[str]) -> int:
    card_rank_map["J"] = 1
    return calculate_total_winnings(rank_hands(parse_hands(lines, True)))


if __name__ == "__main__":
    input = parse_args()
    with open(input) as file:
        lines = file.readlines()
        print(f"part 1: {run_p1(lines)}")
        print(f"part 2: {run_p2(lines)}")
