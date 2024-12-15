const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n\n");

const puzzleMap = input[0].split("\n").map((line) => line.split(""));
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
