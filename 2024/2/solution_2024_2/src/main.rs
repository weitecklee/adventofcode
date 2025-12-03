use std::fs;

fn main() {
    let puzzle_input = fs::read_to_string("../input.txt").expect("Error reading input.txt");
    let levels = parse_input(puzzle_input);
    println!("{}", part1(&levels));
    println!("{}", part2(&levels));
}

fn parse_input(puzzle_input: String) -> Vec<Vec<i32>> {
    puzzle_input
        .lines()
        .map(|line| {
            line.split_ascii_whitespace()
                .map(|s| s.parse::<i32>().unwrap())
                .collect()
        })
        .collect()
}

fn part1(levels: &[Vec<i32>]) -> usize {
    levels.iter().filter(|l| is_safe(l)).count()
}

fn part2(levels: &[Vec<i32>]) -> usize {
    levels.iter().filter(|l| is_safe_with_tolerance(l)).count()
}

fn is_safe(level: &[i32]) -> bool {
    let sign = (level[1] - level[0]).signum();
    if sign == 0 {
        return false;
    }

    level.windows(2).all(|w| {
        let diff = w[1] - w[0];
        diff.signum() == sign && diff.abs() <= 3
    })
}

fn is_safe_with_tolerance(level: &[i32]) -> bool {
    if is_safe(level) {
        return true;
    }

    (0..level.len()).any(|i| {
        let mut level2 = Vec::with_capacity(level.len() - 1);
        level2.extend_from_slice(&level[..i]);
        level2.extend_from_slice(&level[i + 1..]);
        is_safe(&level2)
    })
}
