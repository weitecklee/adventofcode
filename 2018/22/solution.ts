import * as fs from "fs";
import * as path from "path";

const puzzleInput = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split(": ")[1]);

const depth = Number(puzzleInput[0]);
const target = puzzleInput[1].split(",").map(Number);

const erosionMap: Map<string, number> = new Map();
const geologicMap: Map<string, number> = new Map();
geologicMap.set("0,0", 0);
geologicMap.set(puzzleInput[1], 0);

function calcErosionLevel(x: number, y: number): number {
  if (!erosionMap.has(`${x},${y}`)) {
    erosionMap.set(`${x},${y}`, (calcGeologicIndex(x, y) + depth) % 20183);
  }
  return erosionMap.get(`${x},${y}`)!;
}

function calcGeologicIndex(x: number, y: number): number {
  if (!geologicMap.has(`${x},${y}`)) {
    let gi = 0;
    if (y === 0) gi = 16807 * x;
    else if (x === 0) gi = 48271 * y;
    else gi = calcErosionLevel(x - 1, y) * calcErosionLevel(x, y - 1);
    geologicMap.set(`${x},${y}`, gi);
  }
  return geologicMap.get(`${x},${y}`)!;
}

let part1 = 0;
for (let x = 0; x <= target[0]; x++) {
  for (let y = 0; y <= target[1]; y++) {
    part1 += calcErosionLevel(x, y) % 3;
  }
}

console.log(part1);
