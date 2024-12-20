const fs = require("fs");
const path = require("path");
const MinHeap = require("../../utils/MinHeap");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split(""));

let startPos, endPos;
const directions = [
  [1, 0],
  [-1, 0],
  [0, 1],
  [0, -1],
];

for (let r = 1; r < input.length - 1; r++) {
  for (let c = 1; c < input.length - 1; c++) {
    if (input[r][c] === "S") {
      startPos = [r, c];
      if (endPos) break;
    } else if (input[r][c] === "E") {
      endPos = [r, c];
      if (startPos) break;
    }
  }
  if (startPos && endPos) break;
}

const maxR = input.length - 1;
const maxC = input[0].length - 1;

function isEdge(r, c) {
  return r <= 0 || r >= maxR || c <= 0 || c >= maxC;
}

const stepMap = new Map([[startPos.join(","), 0]]);
let [r, c] = startPos;
let steps = 0;
while (r != endPos[0] || c != endPos[1]) {
  for (const [dr, dc] of directions) {
    const r2 = r + dr;
    const c2 = c + dc;
    if (isEdge(r2, c2)) continue;
    if (input[r2][c2] === "#") continue;
    const coord = `${r2},${c2}`;
    if (stepMap.has(coord)) continue;
    stepMap.set(coord, ++steps);
    r = r2;
    c = c2;
    break;
  }
}

function cheat(cheatLimit, stepsSavedTolerance) {
  let count = 0;
  for (const [coord, steps] of stepMap) {
    const [r, c] = coord.split(",").map(Number);
    const queue = [[0, r, c]];
    let i = 0;
    const visited = new Set([`${r},${c}`]);
    while (i < queue.length) {
      const [cheatSteps, r2, c2] = queue[i];
      i++;
      if (input[r2][c2] !== "#") {
        const stepsSaved = stepMap.get(`${r2},${c2}`) - steps - cheatSteps;
        if (stepsSaved >= stepsSavedTolerance) {
          count++;
        }
      }
      if (cheatSteps >= cheatLimit) continue;
      for (const [dr, dc] of directions) {
        const [r3, c3] = [r2 + dr, c2 + dc];
        if (isEdge(r3, c3)) continue;
        if (visited.has(`${r3},${c3}`)) continue;
        visited.add(`${r3},${c3}`);
        queue.push([cheatSteps + 1, r3, c3]);
      }
    }
  }
  return count;
}

console.log(cheat(2, 100));
console.log(cheat(20, 100));

/*
  Problem statement was kinda confusing, had to look up what others were saying.
  Basically, treat each spot on the track as a possible starting spot, then BFS
  to max 20 steps away. Whenever you land on the track again, see if you saved
  any time (compare how long it takes to get there without cheating and with
  cheating).
*/
