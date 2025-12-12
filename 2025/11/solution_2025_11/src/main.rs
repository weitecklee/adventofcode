use std::{collections::HashMap, fs};

fn main() {
    let puzzle_input = fs::read_to_string("../input.txt").expect("Error reading input.txt");
    let (nodes, mut name_map) = parse_input(&puzzle_input);
    let mut memo = prepare_data(&nodes);
    println!("{}", part1(&nodes, &mut name_map, &mut memo));
    println!("{}", part2(&nodes, &mut name_map, &mut memo));
}

struct NameMap<'a> {
    map: HashMap<&'a str, usize>,
}

impl<'a> NameMap<'a> {
    fn new() -> Self {
        Self {
            map: HashMap::new(),
        }
    }

    fn insert(&mut self, name: &'a str) -> usize {
        let n = self.map.len();
        self.map.insert(name, n);
        n
    }

    fn get(&mut self, name: &'a str) -> usize {
        if let Some(&idx) = self.map.get(name) {
            idx
        } else {
            self.insert(name)
        }
    }
}

fn parse_input<'a>(puzzle_input: &'a str) -> (Vec<Vec<usize>>, NameMap<'a>) {
    let lines: Vec<&str> = puzzle_input.lines().collect();
    let mut nodes: Vec<Vec<usize>> = vec![vec![0_usize; 0]; lines.len() + 1];
    let mut name_map = NameMap::new();
    for line in lines {
        let (name, out_str) = line.split_once(": ").unwrap();
        let name_idx = name_map.get(name);
        let outs: Vec<usize> = out_str
            .split_whitespace()
            .map(|s| name_map.get(s))
            .collect();
        nodes[name_idx] = outs;
    }

    (nodes, name_map)
}

fn prepare_data(nodes: &[Vec<usize>]) -> HashMap<[usize; 2], usize> {
    let mut memo = HashMap::new();
    nodes.iter().enumerate().for_each(|(i, outs)| {
        outs.iter().for_each(|&j| {
            memo.insert([i, j], 1);
        });
    });

    memo
}

fn dfs(
    from_node_idx: usize,
    to_node_idx: usize,
    nodes: &Vec<Vec<usize>>,
    memo: &mut HashMap<[usize; 2], usize>,
) -> usize {
    let pair = [from_node_idx, to_node_idx];
    if let Some(&n) = memo.get(&pair) {
        return n;
    }
    let mut res = 0;
    for &out in &nodes[from_node_idx] {
        res += dfs(out, to_node_idx, nodes, memo);
    }
    memo.insert(pair, res);
    res
}

fn part1(
    nodes: &Vec<Vec<usize>>,
    name_map: &mut NameMap,
    memo: &mut HashMap<[usize; 2], usize>,
) -> usize {
    let you_node_idx = name_map.get("you");
    let out_node_idx = name_map.get("out");

    dfs(you_node_idx, out_node_idx, nodes, memo)
}

fn part2(
    nodes: &Vec<Vec<usize>>,
    name_map: &mut NameMap,
    memo: &mut HashMap<[usize; 2], usize>,
) -> usize {
    let svr_node_idx = name_map.get("svr");
    let dac_node_idx = name_map.get("dac");
    let fft_node_idx = name_map.get("fft");
    let out_node_idx = name_map.get("out");

    let svr2dac = dfs(svr_node_idx, dac_node_idx, nodes, memo);
    let dac2fft = dfs(dac_node_idx, fft_node_idx, nodes, memo);
    let fft2out = dfs(fft_node_idx, out_node_idx, nodes, memo);
    let svr2fft = dfs(svr_node_idx, fft_node_idx, nodes, memo);
    let fft2dac = dfs(fft_node_idx, dac_node_idx, nodes, memo);
    let dac2out = dfs(dac_node_idx, out_node_idx, nodes, memo);

    svr2dac * dac2fft * fft2out + svr2fft * fft2dac * dac2out
}
