use std::fs;

fn main() {
    let puzzle_input = fs::read_to_string("../input.txt").expect("Error reading input.txt");
    let (mut ingredient_ranges, ingredient_ids) = parse_input(&puzzle_input);
    ingredient_ranges = condense_ranges(ingredient_ranges);
    println!("{}", part1(&ingredient_ranges, &ingredient_ids));
    println!("{}", part2(&ingredient_ranges));
}

fn parse_input(puzzle_input: &str) -> (Vec<[u64; 2]>, Vec<u64>) {
    let (part1, part2) = puzzle_input.split_once("\n\n").unwrap();
    let ranges: Vec<[u64; 2]> = part1
        .lines()
        .map(|l| {
            let (start, end) = l.split_once("-").unwrap();
            [start.parse().unwrap(), end.parse().unwrap()]
        })
        .collect();
    let ids: Vec<u64> = part2.lines().map(|l| l.parse().unwrap()).collect();
    (ranges, ids)
}

fn condense_ranges(mut ranges: Vec<[u64; 2]>) -> Vec<[u64; 2]> {
    ranges.sort_by_key(|r| r[0]);
    let mut res = vec![ranges[0]];
    for range in ranges.into_iter().skip(1) {
        let last = res.last_mut().unwrap();
        if last[1] >= range[0] {
            last[1] = last[1].max(range[1]);
        } else {
            res.push(range);
        }
    }
    res
}

fn part1(ranges: &[[u64; 2]], ids: &[u64]) -> usize {
    ids.iter()
        .filter(|&&id| ranges.iter().any(|r| id >= r[0] && id <= r[1]))
        .count()
}

fn part2(ranges: &[[u64; 2]]) -> u64 {
    ranges
        .iter()
        .fold(0, |acc, curr| acc + curr[1] - curr[0] + 1)
}
