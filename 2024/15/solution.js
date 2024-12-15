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
const movements = input[1].replace(/\n/g, "");

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

  locateSelf(map) {
    for (let r = 0; r < map.length; r++) {
      for (let c = 0; c < map[r].length; c++) {
        if (map[r][c] === "@") {
          robot.r = r;
          robot.c = c;
          return;
        }
      }
    }
  }

  moves(movements, map, part2 = false) {
    for (const dir of movements) {
      if (part2) this.move2(dir, map);
      else this.move(dir, map);
    }
  }

  move(dir, map) {
    const [dr, dc] = directions.get(dir);
    let [r, c] = [this.r + dr, this.c + dc];
    while (map[r][c] === "O") {
      r += dr;
      c += dc;
    }
    if (map[r][c] === "#") return;
    r -= dr;
    c -= dc;
    while (map[r][c] === "O") {
      map[r][c] = ".";
      map[r + dr][c + dc] = "O";
      r -= dr;
      c -= dc;
    }
    map[r][c] = ".";
    map[r + dr][c + dc] = "@";
    this.r += dr;
    this.c += dc;
  }

  move2(dir, map) {
    const [dr, dc] = directions.get(dir);
    if (dc != 0) {
      let [r, c] = [this.r, this.c + dc];
      while (map[r][c] === "[" || map[r][c] === "]") {
        c += dc;
      }
      if (map[r][c] === "#") return;
      c -= dc;
      while (map[r][c] === "[" || map[r][c] === "]") {
        map[r][c - dc] = ".";
        if (dc > 0) {
          map[r][c] = "[";
          map[r][c + dc] = "]";
        } else {
          map[r][c] = "]";
          map[r][c + dc] = "[";
        }
        c -= 2 * dc;
      }
      map[r][c] = ".";
      map[r][c + dc] = "@";
      this.c += dc;
    } else {
      let [r, c] = [this.r + dr, this.c];
      const affectedCols = [new Set([c])];
      while (true) {
        const newAffectedCols = new Set();
        for (const col of affectedCols[affectedCols.length - 1]) {
          if (map[r][col] === "#") return;
          if (map[r][col] === "[") {
            newAffectedCols.add(col);
            newAffectedCols.add(col + 1);
          } else if (map[r][col] === "]") {
            newAffectedCols.add(col);
            newAffectedCols.add(col - 1);
          }
        }
        if (newAffectedCols.size === 0) break;
        affectedCols.push(newAffectedCols);
        r += dr;
      }
      for (let i = affectedCols.length - 1; i >= 0; i--) {
        r -= dr;
        for (const col of affectedCols[i]) {
          map[r + dr][col] = map[r][col];
          map[r][col] = ".";
        }
      }
      this.r += dr;
    }
  }
}

function sumGPS(map) {
  let sum = 0;
  for (let r = 0; r < map.length; r++) {
    for (let c = 0; c < map[r].length; c++) {
      if (map[r][c] === "O" || map[r][c] === "[") sum += 100 * r + c;
    }
  }
  return sum;
}

const robot = new Robot();

robot.locateSelf(puzzleMap);
robot.moves(movements, puzzleMap);
console.log(sumGPS(puzzleMap));

robot.locateSelf(puzzleMap2);
robot.moves(movements, puzzleMap2, true);
console.log(sumGPS(puzzleMap2));
