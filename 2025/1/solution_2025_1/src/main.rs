use std::fs;

fn main() {
    let puzzle_input = fs::read_to_string("../input.txt").expect("Error reading input.txt");
    let turns = parse_input(puzzle_input);
    println!("{}", part1(&turns));
    println!("{}", part2(&turns));
}

fn parse_input(puzzle_input: String) -> Vec<i32> {
    puzzle_input
        .lines()
        .map(|line| {
            line.get(1..)
                .unwrap()
                .parse::<i32>()
                .expect("Failed to parse number")
                * (if line.starts_with('L') { -1 } else { 1 })
        })
        .collect()
}

fn part1(turns: &Vec<i32>) -> i32 {
    let mut dial = 50;
    let mut res = 0;
    for turn in turns {
        dial += turn;
        dial %= 100;
        if dial == 0 {
            res += 1;
        }
    }
    res
}

fn part2(turns: &Vec<i32>) -> i32 {
    let mut dial = 50;
    let mut res = 0;
    let mut prev = dial;
    for turn in turns {
        res += turn.abs() / 100;
        dial += turn % 100;
        if (prev < 0 && dial >= 0) || (prev > 0 && dial <= 0) {
            res += 1;
        }
        if prev < 100 && dial >= 100 {
            res += 1;
            dial -= 100;
        }
        if prev > -100 && dial <= -100 {
            res += 1;
            dial += 100;
        }
        prev = dial;
    }
    res
}
