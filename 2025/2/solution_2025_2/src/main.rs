use rayon::prelude::*;
use std::{fs, time::Instant};

fn main() {
    let puzzle_input = fs::read_to_string("../input.txt").expect("Error reading input.txt");
    let ranges = parse_input(puzzle_input);
    println!("{}", part1(&ranges));
    let now = Instant::now();
    println!("{}", part2(&ranges));
    println!("Elapsed: {}", now.elapsed().as_secs_f32())
}

fn parse_input(puzzle_input: String) -> Vec<[i64; 2]> {
    puzzle_input
        .split(",")
        .map(|s| s.split("-").collect())
        .map(|ss: Vec<&str>| [ss[0].parse::<i64>().unwrap(), ss[1].parse::<i64>().unwrap()])
        .collect()
}

fn part1(ranges: &[[i64; 2]]) -> i64 {
    ranges
        .par_iter()
        .map(|pair| pair[0]..=pair[1])
        .map(|r| r.filter(|n| is_invalid_id(*n)))
        .map(|r| r.sum::<i64>())
        .sum()
}

fn part2(ranges: &[[i64; 2]]) -> i64 {
    ranges
        .par_iter()
        .map(|pair| pair[0]..=pair[1])
        .map(|r| r.filter(|n| is_invalid_id2(*n)))
        .map(|r| r.sum::<i64>())
        .sum()
}

fn is_invalid_id(n: i64) -> bool {
    let s = n.to_string();
    s[0..s.len() / 2] == s[s.len() / 2..]
}

fn is_invalid_id2(n: i64) -> bool {
    let s = n.to_string();
    let l = s.len();
    for i in 1..=l / 2 {
        if !l.is_multiple_of(i) {
            continue;
        }
        let mut parts = Vec::with_capacity(l / i);
        for j in 0..l / i {
            parts.push(&s[i * j..i * (j + 1)])
        }
        if are_all_the_same_string(parts) {
            return true;
        }
    }
    false
}

fn are_all_the_same_string(ss: Vec<&str>) -> bool {
    for s in &ss[1..] {
        if ss[0] != *s {
            return false;
        }
    }
    true
}
