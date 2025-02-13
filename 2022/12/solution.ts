import * as fs from "fs";
import * as path from "path";
import MinHeap from "../../utils/MinHeap";

const puzzleInput = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split("").map((b) => b.charCodeAt(0)));

const directions = [
  [-1, 0],
  [1, 0],
  [0, -1],
  [0, 1],
];

const rMax = puzzleInput.length - 1;
const cMax = puzzleInput[0].length - 1;

let startPos: number[] = [];
let endPos: number[] = [];
let startFound = false;
let endFound = false;
for (let r = 0; !(startFound && endFound) && r <= rMax; r++) {
  for (let c = 0; !(startFound && endFound) && c <= cMax; c++) {
    if (puzzleInput[r][c] === "S".charCodeAt(0)) {
      startPos = [r, c];
      puzzleInput[r][c] = "a".charCodeAt(0);
      startFound = true;
    } else if (puzzleInput[r][c] === "E".charCodeAt(0)) {
      endPos = [r, c];
      puzzleInput[r][c] = "z".charCodeAt(0);
      endFound = true;
    }
  }
}

interface QueueEntry {
  steps: number;
  r: number;
  c: number;
}

function distanceToEnd(r: number, c: number): number {
  return Math.abs(r - endPos[0]) + Math.abs(c - endPos[1]);
}

function findMinSteps(startPositions: number[][]): number {
  const queue: [number, QueueEntry][] = startPositions.map((a) => [
    0,
    { steps: 0, r: a[0], c: a[1] },
  ]);
  const visited: Map<string, number> = new Map(
    startPositions.map((a) => [a.join(","), 0])
  );

  while (queue.length) {
    let [_, { steps, r, c }] = MinHeap.pop(queue) as [number, QueueEntry];
    if (r === endPos[0] && c === endPos[1]) {
      return steps;
    }
    steps++;
    const elev = puzzleInput[r][c];
    for (const [dr, dc] of directions) {
      const [r2, c2] = [r + dr, c + dc];
      if (r2 < 0 || c2 < 0 || r2 > rMax || c2 > cMax) continue;
      if (puzzleInput[r2][c2] - elev > 1) continue;
      if (visited.has(`${r2},${c2}`) && visited.get(`${r2},${c2}`)! <= steps)
        continue;
      visited.set(`${r2},${c2}`, steps);
      MinHeap.push(queue, [
        distanceToEnd(r2, c2) + steps,
        {
          steps,
          elev: puzzleInput[r2][c2],
          r: r2,
          c: c2,
        },
      ]);
    }
  }
  return -1;
}

function part1(): number {
  return findMinSteps([startPos]);
}

function part2(): number {
  const startPositions: number[][] = [];
  for (let r = 0; r <= rMax; r++) {
    for (let c = 0; c <= cMax; c++) {
      if (puzzleInput[r][c] === "a".charCodeAt(0)) {
        startPositions.push([r, c]);
      }
    }
  }
  return findMinSteps(startPositions);
}

console.log(part1());
console.log(part2());
