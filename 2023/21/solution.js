const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split(""));

const start = [0, 0];
let startFound = false;

for (let r = 0; r < input.length; r++) {
  for (let c = 0; c < input[0].length; c++) {
    if (input[r][c] === "S") {
      [start[0], start[1]] = [c, r];
      startFound = true;
      break;
    }
  }
  if (startFound) break;
}

const reachedPlots = new Set();
const directions = [
  [0, 1],
  [1, 0],
  [-1, 0],
  [0, -1],
];

function walk(x, y, steps) {
  if (steps % 2 === 0) {
    reachedPlots.add(`${x},${y}`);
  }
  if (steps === 0) {
    return [];
  }
  const next = [];
  for (const [dx, dy] of directions) {
    const x2 = x + dx;
    const y2 = y + dy;
    if (x2 < 0 || x2 >= input[0].length || y2 < 0 || y2 >= input.length)
      continue;
    if (input[y2][x2] !== ".") continue;
    input[y2][x2] = "X";
    next.push([x2, y2, steps - 1]);
  }
  return next;
}

const queue = [[start[0], start[1], 64]];

while (queue.length) {
  const [x, y, steps] = queue.shift();
  const next = walk(x, y, steps);
  queue.push(...next);
}
console.log(reachedPlots.size);
