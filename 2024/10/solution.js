const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split("").map(Number));

const trailheads = [];
for (let r = 0; r < input.length; r++) {
  for (let c = 0; c < input[0].length; c++) {
    if (input[r][c] === 0) {
      trailheads.push([r, c]);
    }
  }
}

const directions = [
  [0, 1],
  [1, 0],
  [-1, 0],
  [0, -1],
];

let part1 = 0;
for (const trailhead of trailheads) {
  let count = 0;
  const visited = new Set();
  const queue = [trailhead];
  while (queue.length) {
    const [r, c] = queue.shift();
    if (visited.has(`${r},${c}`)) continue;
    visited.add(`${r},${c}`);
    if (input[r][c] === 9) {
      count++;
      continue;
    }
    const curr = input[r][c];
    for (const [dr, dc] of directions) {
      const r2 = r + dr;
      const c2 = c + dc;
      if (r2 < 0 || r2 >= input.length || c2 < 0 || c2 >= input[0].length)
        continue;
      if (input[r2][c2] === curr + 1) {
        queue.push([r2, c2]);
      }
    }
  }

  part1 += count;
}

console.log(part1);

let part2 = 0;
for (const trailhead of trailheads) {
  let count = 0;
  const queue = [[...trailhead, new Set()]];
  while (queue.length) {
    const [r, c, visited] = queue.shift();
    if (visited.has(`${r},${c}`)) continue;
    visited.add(`${r},${c}`);
    if (input[r][c] === 9) {
      count++;
      continue;
    }
    const curr = input[r][c];
    for (const [dr, dc] of directions) {
      const r2 = r + dr;
      const c2 = c + dc;
      if (r2 < 0 || r2 >= input.length || c2 < 0 || c2 >= input[0].length)
        continue;
      if (input[r2][c2] === curr + 1) {
        queue.push([r2, c2, new Set(visited)]);
      }
    }
  }

  part2 += count;
}

console.log(part2);
