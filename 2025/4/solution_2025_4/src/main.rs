use std::{collections::HashMap, fs};

fn main() {
    let puzzle_input = fs::read_to_string("../input.txt").expect("Error reading input.txt");
    let mut paper_map = parse_input(&puzzle_input);
    println!("{}", part1(&paper_map));
    println!("{}", part2(&mut paper_map));
}

fn parse_input(puzzle_input: &str) -> HashMap<[isize; 2], i8> {
    let mut map = HashMap::new();

    for (r, line) in puzzle_input.lines().enumerate() {
        for (c, ch) in line.chars().enumerate() {
            if ch == '@' {
                map.insert([r as isize, c as isize], 0);
            }
        }
    }

    let coords: Vec<[isize; 2]> = map.keys().copied().collect();

    for [row, col] in coords {
        for r in row - 1..=row + 1 {
            for c in col - 1..=col + 1 {
                if r == row && c == col {
                    continue;
                }
                if let Some(n) = map.get(&[r, c]) {
                    map.insert([r, c], n + 1);
                }
            }
        }
    }

    map
}

fn part1(paper_map: &HashMap<[isize; 2], i8>) -> usize {
    paper_map.values().filter(|&&v| v < 4).count()
}

fn part2(paper_map: &mut HashMap<[isize; 2], i8>) -> isize {
    let mut res = 0;
    loop {
        let coords: Vec<[isize; 2]> = paper_map.keys().copied().collect();
        let mut removed = 0;
        for [row, col] in coords {
            if let Some(v) = paper_map.get(&[row, col])
                && *v >= 4
            {
                continue;
            }
            for r in row - 1..=row + 1 {
                for c in col - 1..=col + 1 {
                    if r == row && c == col {
                        continue;
                    }
                    if let Some(n) = paper_map.get(&[r, c]) {
                        paper_map.insert([r, c], n - 1);
                    }
                }
            }
            paper_map.remove(&[row, col]);
            removed += 1;
        }
        if removed == 0 {
            break;
        }
        res += removed;
    }
    res
}
