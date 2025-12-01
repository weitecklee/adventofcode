use std::fs;

fn main() {
    let puzzle_input = fs::read_to_string("../input.txt").expect("Error reading input.txt");
    let mut elves = parse_input(puzzle_input);
    elves.sort();
    println!("{}", part1(&elves));
    println!("{}", part2(&elves));
}

fn parse_input(puzzle_input: String) -> Vec<i32> {
    let mut elves = Vec::new();
    let mut curr = 0;
    for line in puzzle_input.lines() {
        if line.is_empty() {
            elves.push(curr);
            curr = 0;
        } else {
            curr += line.parse::<i32>().unwrap();
        }
    }
    elves.push(curr);
    elves
}

fn part1(elves: &[i32]) -> i32 {
    // *elves.iter().max().unwrap()
    *elves.last().unwrap()
}

fn part2(elves: &[i32]) -> i32 {
    elves[elves.len() - 1] + elves[elves.len() - 2] + elves[elves.len() - 3]
}
