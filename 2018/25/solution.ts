import * as fs from "fs";
import * as path from "path";

const puzzleInput = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split(",").map(Number));

type Point = [number, number, number, number];

function calcDist(p1: Point, p2: Point): number {
  return p1.reduce((a, b, i) => a + Math.abs(b - p2[i]), 0);
}

const points = puzzleInput as Point[];
const edges: [Point, Point][] = [];

for (let i = 0; i < points.length; i++) {
  for (let j = i + 1; j < points.length; j++) {
    if (calcDist(points[i], points[j]) <= 3) edges.push([points[i], points[j]]);
  }
}

class UnionFind<T> {
  parent: Map<T, T>;
  size: Map<T, number>;
  count: number;
  constructor(elements: T[], edges: [T, T][]) {
    this.count = elements.length;
    this.parent = new Map(elements.map((a) => [a, a]));
    this.size = new Map(elements.map((a) => [a, 1]));
    this.initialize(edges);
  }

  initialize(edges: [T, T][]) {
    for (const [a, b] of edges) {
      this.union(a, b);
    }
  }

  findParent(a: T): T {
    if (a !== this.parent.get(a)) {
      this.parent.set(a, this.findParent(this.parent.get(a)!));
    }
    return this.parent.get(a)!;
  }

  union(a: T, b: T): void {
    const rootA = this.findParent(a);
    const rootB = this.findParent(b);

    if (rootA !== rootB) {
      if (this.size.get(rootA)! > this.size.get(rootB)!) {
        this.parent.set(rootB, rootA);
        this.size.set(rootA, this.size.get(rootA)! + this.size.get(rootB)!);
      } else {
        this.parent.set(rootA, rootB);
        this.size.set(rootB, this.size.get(rootB)! + this.size.get(rootA)!);
      }
      this.count--;
    }
  }
}

const uf = new UnionFind(puzzleInput, edges);
console.log(uf.count);
