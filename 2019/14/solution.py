from collections import defaultdict
import math
import os

def parse_input(input_lines):
    reactions = {}
    for line in input_lines:
        inputs, output = line.split(" => ")
        output_qty, output_chem = output.split(" ")
        inputs = [(int(qty), chem) for qty, chem in (component.split(" ") for component in inputs.split(", "))]
        reactions[output_chem] = (int(output_qty), inputs)
    return reactions

def ore_required(reactions, target_chemical, target_amount, surplus):
    if target_chemical == 'ORE':
        return target_amount

    if target_chemical in surplus:
        available = min(surplus[target_chemical], target_amount)
        surplus[target_chemical] -= available
        target_amount -= available

    if target_amount == 0:
        return 0

    output_qty, inputs = reactions[target_chemical]
    batches = math.ceil(target_amount / output_qty)
    surplus[target_chemical] += batches * output_qty - target_amount

    ore_needed = 0
    for input_qty, input_chem in inputs:
        ore_needed += ore_required(reactions, input_chem, batches * input_qty, surplus)

    return ore_needed

def part1(reactions):
    surplus = defaultdict(int)
    return ore_required(reactions, 'FUEL', 1, surplus)

def ore_required_for_fuel(reactions, fuel_amount, surplus):
    return ore_required(reactions, 'FUEL', fuel_amount, surplus)

def part2(reactions, total_ore):
    low = 0
    high = total_ore

    while low < high:
        mid = (low + high + 1) // 2
        surplus = defaultdict(int)
        required_ore = ore_required_for_fuel(reactions, mid, surplus)

        if required_ore <= total_ore:
            low = mid
        else:
            high = mid - 1

    return low

if __name__ == "__main__":
  with open(os.path.join(os.path.dirname(os.path.abspath(__file__)), 'input.txt'),'r') as file:
    input_lines = [line.strip() for line in file]
  reactions = parse_input(input_lines)
  print(part1(reactions))
  print(part2(reactions, 1000000000000))
