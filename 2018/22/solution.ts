import * as fs from "fs";
import * as path from "path";
import MinHeap from "../../utils/MinHeap";

const puzzleInput = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split(": ")[1]);

const depth = Number(puzzleInput[0]);
const target = puzzleInput[1].split(",").map(Number);

const erosionMap: Map<string, number> = new Map();
const geologicMap: Map<string, number> = new Map();
const riskMap: Map<string, number> = new Map();
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

function calcRisk(x: number, y: number): number {
  if (!riskMap.has(`${x},${y}`)) {
    riskMap.set(`${x},${y}`, calcErosionLevel(x, y) % 3);
  }
  return riskMap.get(`${x},${y}`)!;
}

let part1 = 0;
for (let x = 0; x <= target[0]; x++) {
  for (let y = 0; y <= target[1]; y++) {
    part1 += calcRisk(x, y);
  }
}

console.log(part1);

const toolOptions = [
  new Set(["torch", "climbing gear"]),
  new Set(["climbing gear", "neither"]),
  new Set(["torch", "neither"]),
];
const directions = [
  [0, 1],
  [1, 0],
  [-1, 0],
  [0, -1],
];

function calcDist(x: number, y: number): number {
  return Math.abs(x - target[0]) + Math.abs(y - target[1]);
}

const queue: [number, number, number, number, string][] = [
  [calcDist(0, 0), 0, 0, 0, "torch"],
];
const visited = new Map([["0,0", 0]]);
let part2 = Number.MAX_SAFE_INTEGER;

while (queue.length) {
  const [_, t, x, y, tool] = MinHeap.pop(queue) as [
    number,
    number,
    number,
    number,
    string
  ];
  if (t > part2) continue;
  if (
    visited.has(`${x},${y},${tool}`) &&
    visited.get(`${x},${y},${tool}`)! <= t
  )
    continue;
  visited.set(`${x},${y},${tool}`, t);
  if (x === target[0] && y === target[1]) {
    if (tool === "torch") {
      if (t < part2) {
        part2 = t;
      }
    } else {
      MinHeap.push(queue, [t + 7, t + 7, x, y, "torch"]);
    }
    continue;
  }
  const risk = calcRisk(x, y);
  const d = calcDist(x, y);
  for (const [dx, dy] of directions) {
    const [x2, y2] = [x + dx, y + dy];
    if (x2 < 0 || y2 < 0) continue;
    const d2 = calcDist(x2, y2);
    const risk2 = calcRisk(x2, y2);
    for (const tool2 of toolOptions[risk]) {
      if (!toolOptions[risk2].has(tool2)) continue;
      if (tool === tool2)
        MinHeap.push(queue, [d2 + t + 1, t + 1, x2, y2, tool2]);
      else MinHeap.push(queue, [d + t + 7, t + 7, x, y, tool2]);
    }
  }
}

console.log(part2);
