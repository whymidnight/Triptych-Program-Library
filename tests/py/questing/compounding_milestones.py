"""
This program asserts the following principles:
    Such Base Input is not indicative of the real world.
    Such simulation assumes a fuzzed input.
    Such Output is computed for interpretation of correctness.
"""


class Milestone:
    tick: int
    modifier: float

    def __init__(self, tick: int, modifier: float) -> None:
        self.tick = tick
        self.modifier = modifier

        pass


if __name__ == "__main__":
    import sys

    until = int(sys.argv[1])
    base = float(sys.argv[2])
    base_initial = base

    milestones = [
        Milestone(7, 2.5),
        Milestone(15, 5.0),
        Milestone(30, 10.0),
        Milestone(60, 20.0),
        Milestone(180, 2.5),
        Milestone(365, 2.5),
    ]

    for milestone in milestones:
        if milestone.tick >= until:
            break

        unit = float(base) * (100 - milestone.modifier) / 100
        rounded = round(unit, 1)
        print(
            f"Milestone Day: {milestone.tick} Milestone Halving: {milestone.modifier}% Royalty: {rounded}%"
        )
        base = rounded

    # print("Final Royalty", int(base * pow(10, 2)))
    print(
        f"Initial Royalty: {base_initial}% Final Royalty: {base}% Debug: {base * pow(10, 2)}"
    )

