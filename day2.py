
with open("input/2/input.txt") as f:
    data = [tuple(l.split()) for l in f]

ROCK_THEM = "A"
PAPER_THEM = "B"
SCISSORS_THEM = "C"

ROCK_US = "X"
PAPER_US = "Y"
SCISSORS_US = "Z"
LOOSE = "X"
DRAW = "Y"
WIN = "Z"

OPTIONS_1 = {
    (ROCK_THEM, ROCK_US): 1 + 3,
    (ROCK_THEM, PAPER_US): 2 + 6,
    (ROCK_THEM, SCISSORS_US): 3 + 0,
    (PAPER_THEM, ROCK_US): 1 + 0,
    (PAPER_THEM, PAPER_US): 2 + 3,
    (PAPER_THEM, SCISSORS_US): 3 + 6,
    (SCISSORS_THEM, ROCK_US): 1 + 6,
    (SCISSORS_THEM, PAPER_US): 2 + 0,
    (SCISSORS_THEM, SCISSORS_US): 3 + 3,
}

OPTIONS_2 = {
    (ROCK_THEM, LOOSE): OPTIONS_1[(ROCK_THEM, SCISSORS_US)],
    (ROCK_THEM, DRAW): OPTIONS_1[((ROCK_THEM, ROCK_US))],
    (ROCK_THEM, WIN): OPTIONS_1[(ROCK_THEM, PAPER_US)],
    (PAPER_THEM, LOOSE): OPTIONS_1[(PAPER_THEM, ROCK_US)],
    (PAPER_THEM, DRAW): OPTIONS_1[(PAPER_THEM, PAPER_US)],
    (PAPER_THEM, WIN): OPTIONS_1[(PAPER_THEM, SCISSORS_US)],
    (SCISSORS_THEM, LOOSE): OPTIONS_1[(SCISSORS_THEM, PAPER_US)],
    (SCISSORS_THEM, DRAW): OPTIONS_1[(SCISSORS_THEM, SCISSORS_US)],
    (SCISSORS_THEM, WIN): OPTIONS_1[(SCISSORS_THEM, ROCK_US)],
}


def part1():
    print(sum(OPTIONS_1[o] for o in data))


def part2():
    print(sum(OPTIONS_2[o] for o in data))


part1()
part2()