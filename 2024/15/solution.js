const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n\n");

const puzzleMap = input[0].split("\n").map((line) => line.split(""));
const puzzleMap2 = input[0]
  .replace(/#/g, "##")
  .replace(/O/g, "[]")
  .replace(/\./g, "..")
  .replace(/@/g, "@.")
  .split("\n")
  .map((line) => line.split(""));
const movements = input[1].split("\n").map((line) => line.split(""));

const boxes = new Map();
const directions = new Map([
  [">", [0, 1]],
  ["<", [0, -1]],
  ["^", [-1, 0]],
  ["v", [1, 0]],
]);

class Robot {
  constructor() {
    this.r = 0;
    this.c = 0;
  }
  move(chr) {
    const [dr, dc] = directions.get(chr);
    let [r, c] = [this.r + dr, this.c + dc];
    while (puzzleMap[r][c] === "O") {
      r += dr;
      c += dc;
    }
    if (puzzleMap[r][c] === "#") return;
    r -= dr;
    c -= dc;
    while (puzzleMap[r][c] === "O") {
      puzzleMap[r][c] = ".";
      puzzleMap[r + dr][c + dc] = "O";
      r -= dr;
      c -= dc;
    }
    puzzleMap[r][c] = ".";
    puzzleMap[r + dr][c + dc] = "@";
    this.r += dr;
    this.c += dc;
  }

  move2(chr) {
    const [dr, dc] = directions.get(chr);
    if (dc != 0) {
      let [r, c] = [this.r, this.c + dc];
      while (puzzleMap2[r][c] === "[" || puzzleMap2[r][c] === "]") {
        c += dc;
      }
      if (puzzleMap2[r][c] === "#") return;
      c -= dc;
      while (puzzleMap2[r][c] === "[" || puzzleMap2[r][c] === "]") {
        puzzleMap2[r][c - dc] = ".";
        if (dc > 0) {
          puzzleMap2[r][c] = "[";
          puzzleMap2[r][c + dc] = "]";
        } else {
          puzzleMap2[r][c] = "]";
          puzzleMap2[r][c + dc] = "[";
        }
        c -= 2 * dc;
      }
      puzzleMap2[r][c] = ".";
      puzzleMap2[r][c + dc] = "@";
      this.c += dc;
    } else {
      let [r, c] = [this.r + dr, this.c];
      const affectedCols = [new Set([c])];
      while (true) {
        const newRange = new Set();
        for (const col of affectedCols[affectedCols.length - 1]) {
          if (puzzleMap2[r][col] === "#") return;
          if (puzzleMap2[r][col] === "[") {
            newRange.add(col);
            newRange.add(col + 1);
          } else if (puzzleMap2[r][col] === "]") {
            newRange.add(col);
            newRange.add(col - 1);
          }
        }
        if (newRange.size === 0) break;
        affectedCols.push(newRange);
        r += dr;
      }
      for (let i = affectedCols.length - 1; i >= 0; i--) {
        r -= dr;
        for (const col of affectedCols[i]) {
          puzzleMap2[r + dr][col] = puzzleMap2[r][col];
          puzzleMap2[r][col] = ".";
        }
      }
      this.r += dr;
    }
  }
}

const robot = new Robot();

for (let r = 0; r < puzzleMap.length; r++) {
  for (let c = 0; c < puzzleMap[r].length; c++) {
    if (puzzleMap[r][c] === "@") {
      robot.r = r;
      robot.c = c;
    }
  }
}

for (const line of movements) {
  for (const movement of line) {
    robot.move(movement);
  }
}

let part1 = 0;

for (let r = 0; r < puzzleMap.length; r++) {
  for (let c = 0; c < puzzleMap[r].length; c++) {
    if (puzzleMap[r][c] === "O") part1 += 100 * r + c;
  }
}

console.log(part1);

for (let r = 0; r < puzzleMap2.length; r++) {
  for (let c = 0; c < puzzleMap2[r].length; c++) {
    if (puzzleMap2[r][c] === "@") {
      robot.r = r;
      robot.c = c;
    }
  }
}

for (const line of movements) {
  for (const movement of line) {
    robot.move2(movement);
  }
}

let part2 = 0;

for (let r = 0; r < puzzleMap2.length; r++) {
  for (let c = 0; c < puzzleMap2[r].length; c++) {
    if (puzzleMap2[r][c] === "[") part2 += 100 * r + c;
  }
}
console.log(part2);
