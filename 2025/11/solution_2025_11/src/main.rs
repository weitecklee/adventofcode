use std::{collections::HashMap, fs, ops::Index};

fn main() {
    let puzzle_input = fs::read_to_string("../input.txt").expect("Error reading input.txt");
    let (nodes, name_map) = parse_input(&puzzle_input);
    let mut graph = Graph::new(name_map, nodes);
    println!("{}", part1(&mut graph));
    println!("{}", part2(&mut graph));
}

struct NameMap {
    map: HashMap<String, usize>,
}

impl NameMap {
    fn new() -> Self {
        Self {
            map: HashMap::new(),
        }
    }

    fn insert(&mut self, name: &str) -> usize {
        let n = self.map.len();
        self.map.insert(name.to_string(), n);
        n
    }

    fn get_or_insert(&mut self, name: &str) -> usize {
        if let Some(&idx) = self.map.get(name) {
            idx
        } else {
            self.insert(name)
        }
    }
}

impl Index<&str> for NameMap {
    type Output = usize;

    fn index(&self, name: &str) -> &Self::Output {
        self.map.get(name).unwrap()
    }
}

struct Graph {
    map: NameMap,
    memo: HashMap<(usize, usize), usize>,
    nodes: Vec<Vec<usize>>,
}

impl Graph {
    fn new(map: NameMap, nodes: Vec<Vec<usize>>) -> Self {
        let mut memo = HashMap::new();
        nodes.iter().enumerate().for_each(|(i, outs)| {
            outs.iter().for_each(|&j| {
                memo.insert((i, j), 1);
            });
        });

        Self { map, memo, nodes }
    }

    fn num_paths(&mut self, from: &str, to: &str) -> usize {
        let from_idx = self.map[from];
        let to_idx = self.map[to];
        self.dfs(from_idx, to_idx)
    }

    fn dfs(&mut self, from_idx: usize, to_idx: usize) -> usize {
        let pair = (from_idx, to_idx);
        if let Some(&n) = self.memo.get(&pair) {
            return n;
        }

        let mut res = 0;
        for i in 0..self.nodes[from_idx].len() {
            let out = self.nodes[from_idx][i];
            res += self.dfs(out, to_idx);
        }

        self.memo.insert(pair, res);
        res
    }
}

fn parse_input(puzzle_input: &str) -> (Vec<Vec<usize>>, NameMap) {
    let lines: Vec<&str> = puzzle_input.lines().collect();
    let mut nodes: Vec<Vec<usize>> = vec![Vec::new(); lines.len() + 1]; // +1 due to `out` node
    let mut name_map = NameMap::new();

    for line in lines {
        let (name, out_str) = line.split_once(": ").unwrap();
        let name_idx = name_map.get_or_insert(name);
        let outs: Vec<usize> = out_str
            .split_whitespace()
            .map(|s| name_map.get_or_insert(s))
            .collect();
        nodes[name_idx] = outs;
    }

    (nodes, name_map)
}

fn part1(graph: &mut Graph) -> usize {
    graph.num_paths("you", "out")
}

fn part2(graph: &mut Graph) -> usize {
    graph.num_paths("svr", "dac") * graph.num_paths("dac", "fft") * graph.num_paths("fft", "out")
        + graph.num_paths("svr", "fft")
            * graph.num_paths("fft", "dac")
            * graph.num_paths("dac", "out")
}
