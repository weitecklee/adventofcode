import * as fs from "fs";
import * as path from "path";

const puzzleInput = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split(",").map(Number));

function calcDist(p1: number[], p2: number[]): number {
  return p1.reduce((a, b, i) => a + Math.abs(b - p2[i]), 0);
}

const edges: [number, number][] = [];

for (let i = 0; i < puzzleInput.length; i++) {
  for (let j = i + 1; j < puzzleInput.length; j++) {
    if (calcDist(puzzleInput[i], puzzleInput[j]) <= 3) edges.push([i, j]);
  }
}

class UnionFind {
  parent: number[];
  size: number[];
  count: number;
  constructor(nElements: number, edges: [number, number][]) {
    this.parent = Array(nElements)
      .fill(undefined)
      .map((_, i) => i);
    this.size = Array(nElements).fill(1);
    this.count = nElements;
    this.initialize(edges);
  }

  initialize(edges: [number, number][]) {
    for (const [a, b] of edges) {
      this.union(a, b);
    }
  }

  findParent(a: number): number {
    while (a !== this.parent[a]) {
      this.parent[a] = this.parent[this.parent[a]];
      a = this.parent[a];
    }
    return a;
  }

  union(a: number, b: number): void {
    const parentA = this.findParent(a);
    const parentB = this.findParent(b);
    if (parentA !== parentB) {
      if (this.size[parentA] > this.size[parentB]) {
        this.parent[parentB] = parentA;
        this.size[parentA]++;
      } else {
        this.parent[parentA] = parentB;
        this.size[parentB]++;
      }
      this.count--;
    }
  }
}

const uf = new UnionFind(puzzleInput.length, edges);
console.log(uf.count);
