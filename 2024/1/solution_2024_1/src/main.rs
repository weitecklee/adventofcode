use std::{collections::HashMap, fs};

fn main() {
    let puzzle_input = fs::read_to_string("../input.txt").expect("Error reading input.txt");
    let (mut list1, mut list2) = parse_input(puzzle_input);
    list1.sort();
    list2.sort();
    println!("{}", part1(&list1, &list2));
    println!("{}", part2(&list1, &list2));
}

fn parse_input(puzzle_input: String) -> (Vec<i32>, Vec<i32>) {
    let lines: Vec<&str> = puzzle_input.lines().collect();
    let mut list1 = Vec::with_capacity(lines.len());
    let mut list2 = Vec::with_capacity(lines.len());
    for line in lines {
        let parts: Vec<&str> = line.split_ascii_whitespace().collect();
        let n1 = parts[0].parse::<i32>().unwrap();
        let n2 = parts[1].parse::<i32>().unwrap();
        list1.push(n1);
        list2.push(n2);
    }
    (list1, list2)
}

fn part1(list1: &[i32], list2: &[i32]) -> i32 {
    let mut res = 0;
    for (n1, n2) in list1.iter().zip(list2) {
        res += (n1 - n2).abs();
    }
    res
}

fn part2(list1: &[i32], list2: &[i32]) -> i32 {
    let mut counts1: HashMap<i32, i32> = HashMap::new();
    let mut counts2: HashMap<i32, i32> = HashMap::new();
    for (n1, n2) in list1.iter().zip(list2) {
        *counts1.entry(*n1).or_insert(0) += 1;
        *counts2.entry(*n2).or_insert(0) += 1;
    }
    let mut res = 0;
    for (k, v) in counts1 {
        res += k * v * counts2.get(&k).unwrap_or(&0);
    }
    res
}
