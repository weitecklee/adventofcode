use std::fs;

struct UnionFind {
    parents: Vec<usize>,
    sizes: Vec<usize>,
}

impl UnionFind {
    fn new(n: usize) -> Self {
        Self {
            parents: (0..n).collect(),
            sizes: vec![1; n],
        }
    }

    fn find(&mut self, a: usize) -> usize {
        if self.parents[a] != a {
            self.parents[a] = self.find(self.parents[a]);
        }
        self.parents[a]
    }

    fn union(&mut self, a: usize, b: usize) {
        let mut root_a = self.find(a);
        let mut root_b = self.find(b);

        if root_a == root_b {
            return;
        }

        if self.sizes[root_a] < self.sizes[root_b] {
            std::mem::swap(&mut root_a, &mut root_b);
        }

        self.parents[root_b] = root_a;
        self.sizes[root_a] += self.sizes[root_b];
    }
}

struct Pair {
    a: usize,
    b: usize,
    dist: u64,
}

fn main() {
    let puzzle_input = fs::read_to_string("../input.txt").expect("Error reading input.txt");
    let points = parse_input(&puzzle_input);
    let (pairs, mut uf) = prepare_data(&points);
    println!("{}", part1(&pairs, &mut uf));
    println!("{}", part2(&pairs, &mut uf, &points));
}

fn parse_input(puzzle_input: &str) -> Vec<[u32; 3]> {
    puzzle_input
        .lines()
        .map(|l| {
            let mut parts = l.split(',').map(|s| s.parse::<u32>().unwrap());
            [
                parts.next().unwrap(),
                parts.next().unwrap(),
                parts.next().unwrap(),
            ]
        })
        .collect()
}

fn distance_between_points(a: &[u32; 3], b: &[u32; 3]) -> u64 {
    a.iter()
        .zip(b)
        .map(|(x, y)| {
            let d = x.abs_diff(*y) as u64;
            d * d
        })
        .sum()
}

fn prepare_data(points: &[[u32; 3]]) -> (Vec<Pair>, UnionFind) {
    let mut pairs: Vec<Pair> = points
        .iter()
        .enumerate()
        .flat_map(|(i, p1)| {
            points[i + 1..].iter().enumerate().map(move |(j, p2)| Pair {
                a: i,
                b: i + j + 1,
                dist: distance_between_points(p1, p2),
            })
        })
        .collect();

    pairs.sort_by_key(|p| p.dist);

    let uf = UnionFind::new(points.len());

    (pairs, uf)
}

fn part1(pairs: &[Pair], uf: &mut UnionFind) -> usize {
    for pair in pairs.iter().take(1000) {
        uf.union(pair.a, pair.b);
    }

    let mut top3 = [0_usize; 3];

    for &s in &uf.sizes {
        if s >= top3[0] {
            top3[2] = top3[1];
            top3[1] = top3[0];
            top3[0] = s;
        } else if s >= top3[1] {
            top3[2] = top3[1];
            top3[1] = s;
        } else if s > top3[2] {
            top3[2] = s;
        }
    }

    top3.iter().product()
}

fn part2(pairs: &[Pair], uf: &mut UnionFind, points: &[[u32; 3]]) -> u32 {
    for pair in pairs.iter().skip(1000) {
        uf.union(pair.a, pair.b);
        let parent0 = uf.find(0);
        if uf.sizes[parent0] == 1000 {
            return points[pair.a][0] * points[pair.b][0];
        }
    }

    0
}
