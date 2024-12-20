const fs = require("fs");
const path = require("path");

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

let part1 = 0;
for (const [coord, steps] of stepMap) {
  const [r, c] = coord.split(",").map(Number);
  for (const [dr, dc] of directions) {
    const [r2, c2] = [r + dr, c + dc];
    if (isEdge(r2, c2)) continue;
    if (input[r2][c2] === ".") continue;
    const [r3, c3] = [r + 2 * dr, c + 2 * dc];
    if (isEdge(r3, c3)) continue;
    if (input[r3][c3] === "#") continue;
    const steps2 = stepMap.get(`${r3},${c3}`);
    if (steps2 >= steps + 102) part1++;
    // 2 steps to walk through wall, then check if at least 100 saved
  }
}

console.log(part1);
