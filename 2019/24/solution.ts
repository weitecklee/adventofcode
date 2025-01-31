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
