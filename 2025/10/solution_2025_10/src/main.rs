use lpsolve::prelude::*;
use rayon::prelude::*;
use std::{
    collections::{BinaryHeap, HashSet},
    fs,
};

const BRACKETS: [char; 6] = ['(', ')', '[', ']', '{', '}'];

#[derive(Eq, PartialEq)]
struct Entry {
    score: usize,
    light_int: usize,
    pressed: usize,
}

impl Ord for Entry {
    fn cmp(&self, other: &Self) -> std::cmp::Ordering {
        other
            .score
            .cmp(&self.score)
            .then_with(|| self.light_int.cmp(&other.light_int))
    }
}

impl PartialOrd for Entry {
    fn partial_cmp(&self, other: &Self) -> Option<std::cmp::Ordering> {
        Some(self.cmp(other))
    }
}

#[derive(Debug)]
struct Machine {
    // light: String,
    light_int: usize,
    buttons: Vec<Vec<usize>>,
    button_ints: Vec<usize>,
    joltages: Vec<usize>,
}

impl Machine {
    fn new(line: &str) -> Machine {
        let parts: Vec<&str> = line.split_whitespace().collect();
        let light = parts[0].trim_matches(BRACKETS);
        let light_int = light
            .chars()
            .enumerate()
            .fold(0, |a, (i, c)| if c == '#' { a | (1 << i) } else { a });
        let buttons = parts[1..parts.len() - 1]
            .iter()
            .map(|p| {
                p.trim_matches(BRACKETS)
                    .split(',')
                    .map(|n| n.parse().unwrap())
                    .collect()
            })
            .collect::<Vec<Vec<usize>>>();
        let button_ints = buttons
            .iter()
            .map(|b| b.iter().fold(0, |a, b| a | (1 << b)))
            .collect();
        let joltages = parts
            .last()
            .unwrap()
            .trim_matches(BRACKETS)
            .split(',')
            .map(|n| n.parse().unwrap())
            .collect();
        Machine {
            // light: light.to_string(),
            light_int,
            buttons,
            button_ints,
            joltages,
        }
    }

    fn fewest_presses_for_lights(&self) -> usize {
        let mut heap = BinaryHeap::new();
        heap.push(Entry {
            score: 0,
            light_int: 0,
            pressed: 0,
        });
        let mut checked: HashSet<usize> = HashSet::new();

        while let Some(Entry {
            score,
            light_int,
            pressed,
        }) = heap.pop()
        {
            if light_int == self.light_int {
                return score;
            }

            if checked.contains(&light_int) {
                continue;
            }
            checked.insert(light_int);

            for (i, button) in self.button_ints.iter().enumerate() {
                if (1 << i) & pressed == 0 {
                    heap.push(Entry {
                        score: score + 1,
                        light_int: light_int ^ button,
                        pressed: pressed | (1 << i),
                    })
                }
            }
        }
        usize::MAX
    }

    fn fewest_presses_for_joltage(&self) -> f64 {
        let n = self.buttons.len();
        let mut lp = Problem::builder()
            .cols(n as i32)
            .min(&vec![1.0; n])
            .non_negative_integers()
            .verbosity(Important);
        for (i, j) in self.joltages.iter().enumerate() {
            let mut con = vec![0.0; n];
            for (b, button) in self.buttons.iter().enumerate() {
                if button.contains(&i) {
                    con[b] = 1.0;
                }
            }
            lp = lp.eq(&con, *j as f64);
        }
        let soln = lp.solve().unwrap();
        soln.variables().unwrap().iter().sum()
    }
}

fn main() {
    let puzzle_input = fs::read_to_string("../input.txt").expect("Error reading input.txt");
    let machines = parse_input(&puzzle_input);

    println!("{}", part1(&machines));
    println!("{}", part2(&machines));
}

fn parse_input(puzzle_input: &str) -> Vec<Machine> {
    puzzle_input.lines().map(Machine::new).collect()
}

fn part1(machines: &[Machine]) -> usize {
    machines
        .par_iter()
        .map(|m| m.fewest_presses_for_lights())
        .sum()
}

fn part2(machines: &[Machine]) -> f64 {
    machines
        .par_iter()
        .map(|m| m.fewest_presses_for_joltage())
        .sum()
}
