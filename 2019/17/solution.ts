import * as fs from "fs";
import * as path from "path";
import intcodeGenerator, { IntcodeGenerator } from "../intcode";

const puzzleInput = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split(",")
  .map(Number);

const camera = intcodeGenerator(puzzleInput);
const scaffold: string[][] = [];
let row: string[] = [];

while (true) {
  const ret = camera.next();
  if (ret.done) break;
  if (ret.value === 10) {
    scaffold.push(row);
    row = [];
  } else {
    row.push(String.fromCharCode(ret.value));
  }
}

const directions = [
  [-1, 0],
  [1, 0],
  [0, -1],
  [0, 1],
];

let part1 = 0;
for (let r = 1; r < scaffold.length - 1; r++) {
  for (let c = 1; c < scaffold[r].length - 1; c++) {
    if (scaffold[r][c] === ".") continue;
    let isIntersection = true;
    for (const [dr, dc] of directions) {
      if (scaffold[r + dr][c + dc] === ".") {
        isIntersection = false;
        break;
      }
    }
    if (isIntersection) part1 += r * c;
  }
}
console.log(part1);
