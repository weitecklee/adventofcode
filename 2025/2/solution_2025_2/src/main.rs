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
        .map(|s| {
            let (a, b) = s.split_once("-").unwrap();
            [a.parse().unwrap(), b.parse().unwrap()]
        })
        .collect()
}

fn part1(ranges: &[[i64; 2]]) -> i64 {
    ranges
        .par_iter()
        .map(|[a, b]| (*a..=*b).filter(|n| is_invalid_id(*n)).sum::<i64>())
        .sum()
}

fn part2(ranges: &[[i64; 2]]) -> i64 {
    ranges
        .par_iter()
        .map(|[a, b]| (*a..=*b).filter(|n| is_invalid_id2(*n)).sum::<i64>())
        .sum()
}

fn is_invalid_id(n: i64) -> bool {
    let s = n.to_string();
    let mid = s.len() / 2;
    s[0..mid] == s[mid..]
}

fn is_invalid_id2(n: i64) -> bool {
    let s = n.to_string();
    let l = s.len();
    for chunk_size in 1..=l / 2 {
        if !l.is_multiple_of(chunk_size) {
            continue;
        }

        let first = &&s.as_bytes()[..chunk_size];
        if s.as_bytes().chunks(chunk_size).all(|c| c == *first) {
            return true;
        }
    }
    false
}
