use std::fs;

fn main() {
    let puzzle_input = fs::read_to_string("../input.txt").expect("Error reading input.txt");
    let batteries = parse_input(&puzzle_input);
    println!("{}", part1(&batteries));
    println!("{}", part2(&batteries));
}

fn parse_input(puzzle_input: &str) -> Vec<Vec<i64>> {
    puzzle_input
        .lines()
        .map(|line| {
            line.chars()
                .map(|d| d.to_digit(10).unwrap() as i64)
                .collect()
        })
        .collect()
}

fn part1(batteries: &[Vec<i64>]) -> i64 {
    batteries.iter().map(|r| find_largest_joltage(r, 2)).sum()
}

fn part2(batteries: &[Vec<i64>]) -> i64 {
    batteries.iter().map(|r| find_largest_joltage(r, 12)).sum()
}

fn find_largest_joltage(digits: &[i64], window_len: usize) -> i64 {
    let mut res = 0;
    let mut idx = 0;
    for i in 0..window_len {
        let (n, idx2) = find_largest_in_window(&digits[idx..digits.len() - window_len + i + 1]);
        idx += idx2 + 1;
        res = res * 10 + n;
    }
    res
}

fn find_largest_in_window(window: &[i64]) -> (i64, usize) {
    let mut max = 0;
    let mut idx = 0;
    for (i, n) in window.iter().enumerate() {
        if *n > max {
            max = *n;
            idx = i;
        }
        if max == 9 {
            break;
        }
    }
    (max, idx)
}
