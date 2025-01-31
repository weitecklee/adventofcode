import * as fs from "fs";
import * as path from "path";

const puzzleInput = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split(""));

const directions = [
  [-1, 0],
  [1, 0],
  [0, -1],
  [0, 1],
];

class Grid {
  grid: string[][];
  history: Set<string>;
  constructor(input: string[][]) {
    this.grid = input;
    this.history = new Set([this.gridKey]);
  }

  get gridKey(): string {
    return this.grid.map((a) => a.join("")).join("");
  }

  iterate(): boolean {
    const grid2 = Array(5)
      .fill("")
      .map(() => Array(5).fill("."));
    for (let r = 0; r < 5; r++) {
      for (let c = 0; c < 5; c++) {
        let bugs = 0;
        for (const [dr, dc] of directions) {
          const r2 = r + dr;
          const c2 = c + dc;
          if (r2 < 0 || r2 >= 5 || c2 < 0 || c2 >= 5) continue;
          if (this.grid[r2][c2] === "#") bugs++;
        }
        if (this.grid[r][c] === "#") {
          if (bugs === 1) grid2[r][c] = "#";
          else grid2[r][c] = ".";
        } else if (bugs === 1 || bugs === 2) grid2[r][c] = "#";
        else grid2[r][c] = ".";
      }
    }
    this.grid = grid2;
    const key = this.gridKey;
    if (this.history.has(key)) return true;
    else {
      this.history.add(key);
      return false;
    }
  }

  get biodiversityRating(): number {
    return parseInt(
      this.gridKey
        .split("")
        .reverse()
        .join("")
        .replaceAll(".", "0")
        .replaceAll("#", "1"),
      2
    );
  }
}

const grid = new Grid(puzzleInput);
while (!grid.iterate()) {}
console.log(grid.biodiversityRating);

class Hypergrid {
  bugs: number[][];
  constructor(input: string[][]) {
    this.bugs = [];
    for (let r = 0; r < input.length; r++) {
      for (let c = 0; c < input[r].length; c++) {
        if (input[r][c] === "#") this.bugs.push([r, c, 0]);
      }
    }
  }

  iterate() {
    const neighborMap: Map<string, number> = new Map();
    for (const [r, c, level] of this.bugs) {
      const neighbors: number[][] = [];
      for (const [dr, dc] of directions) {
        const r2 = r + dr;
        const c2 = c + dc;
        if (r2 < 0) {
          neighbors.push([1, 2, level - 1]);
        } else if (r2 > 4) {
          neighbors.push([3, 2, level - 1]);
        } else if (c2 < 0) {
          neighbors.push([2, 1, level - 1]);
        } else if (c2 > 4) {
          neighbors.push([2, 3, level - 1]);
        } else if (r === 1 && c === 2 && dr === 1) {
          for (let i = 0; i < 5; i++) {
            neighbors.push([0, i, level + 1]);
          }
        } else if (r === 3 && c == 2 && dr == -1) {
          for (let i = 0; i < 5; i++) {
            neighbors.push([4, i, level + 1]);
          }
        } else if (r === 2 && c === 1 && dc === 1) {
          for (let i = 0; i < 5; i++) {
            neighbors.push([i, 0, level + 1]);
          }
        } else if (r === 2 && c == 3 && dc == -1) {
          for (let i = 0; i < 5; i++) {
            neighbors.push([i, 4, level + 1]);
          }
        } else {
          neighbors.push([r2, c2, level]);
        }
      }
      for (const posArr of neighbors) {
        const key = posArr.join(",");
        neighborMap.set(key, (neighborMap.get(key) || 0) + 1);
      }
    }
    const bugs2: number[][] = [];
    for (const bug of this.bugs) {
      if ((neighborMap.get(bug.join(",")) || 0) === 1) {
        bugs2.push(bug);
      }
      neighborMap.delete(bug.join(","));
    }
    for (const [pos, n] of neighborMap) {
      if (n === 1 || n === 2) {
        bugs2.push(pos.split(",").map(Number));
      }
    }
    this.bugs = bugs2;
  }
}

const hypergrid = new Hypergrid(puzzleInput);
for (let i = 0; i < 200; i++) {
  hypergrid.iterate();
}
console.log(hypergrid.bugs.length);
