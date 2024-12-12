const fs = require("fs");
const path = require("path");
const mathjs = require("mathjs");

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

const directions = [
  [0, 1],
  [1, 0],
  [-1, 0],
  [0, -1],
];

function walk(x, y, steps, totalSteps, reachedPlots, visited) {
  if (steps % 2 === totalSteps % 2) {
    reachedPlots.add(`${x},${y}`);
  }
  if (steps === totalSteps) {
    return [];
  }
  const next = [];
  for (const [dx, dy] of directions) {
    const x2 = x + dx;
    const y2 = y + dy;
    if (x2 < 0 || x2 >= input[0].length || y2 < 0 || y2 >= input.length)
      continue;
    if (input[y2][x2] === "#") continue;
    if (visited.has(`${x2},${y2}`)) continue;
    visited.add(`${x2},${y2}`);
    next.push([x2, y2, steps + 1]);
  }
  return next;
}

function part1(totalSteps) {
  const reachedPlots = new Set();
  const visited = new Set();
  const queue = [[start[0], start[1], 0]];

  let i = 0;
  while (i < queue.length) {
    const [x, y, steps] = queue[i];
    const next = walk(x, y, steps, totalSteps, reachedPlots, visited);
    queue.push(...next);
    i++;
  }
  return reachedPlots.size;
}

console.log(part1(64));

/*
  Analysis of input shows that S is (always?) at center of 131x131 grid,
  with a border of empty spaces 65 steps away from S in each direction
  (Manhattan distance).
  Curiously, the puzzle number 26501365 is 65 mod 131.
  If we examine the number of plots reached for steps that are 65 mod 131
  (e.g., 65, 196, 327, 458, etc.), there is a quadratic relationship between
  k and the number of plots reached, where steps = 131 * k + 65.
  Get 3 data points to calculate the quadratic equation and plug in our number.
*/

const wd = input[0].length;
const ht = input.length;

function walkInfinite(x, y, steps, totalSteps, reachedPlots, visited) {
  if (steps % 2 === totalSteps % 2) {
    reachedPlots.add(`${x},${y}`);
  }
  if (steps === totalSteps) {
    return [];
  }
  const next = [];
  for (const [dx, dy] of directions) {
    const x2 = x + dx;
    const y2 = y + dy;
    const x2mod = x2 % wd;
    const y2mod = y2 % ht;
    if (
      input[y2mod >= 0 ? y2mod : ht + y2mod][
        x2mod >= 0 ? x2mod : wd + x2mod
      ] === "#"
    )
      continue;
    if (visited.has(`${x2},${y2}`)) continue;
    visited.add(`${x2},${y2}`);
    next.push([x2, y2, steps + 1]);
  }
  return next;
}

function part2(totalSteps) {
  const reachedPlots = new Set();
  const visited = new Set();
  const queue = [[start[0], start[1], 0]];

  let i = 0;
  while (i < queue.length) {
    const [x, y, stepsRemaining] = queue[i];
    const next = walkInfinite(
      x,
      y,
      stepsRemaining,
      totalSteps,
      reachedPlots,
      visited
    );
    queue.push(...next);
    i++;
  }
  return reachedPlots.size;
}

const A = [0, 1, 2].map((x) => [x * x, x, 1]);
const b = [0, 1, 2].map((k) => part2(131 * k + 65));
const quadraticTerms = mathjs
  .lusolve(mathjs.matrix(A), mathjs.matrix(b))
  ._data.flat();

const n = 26501365;
const k = (n - 65) / 131;
function quadratic(x) {
  return quadraticTerms[0] * x * x + quadraticTerms[1] * x + quadraticTerms[2];
}

console.log(quadratic(k));
