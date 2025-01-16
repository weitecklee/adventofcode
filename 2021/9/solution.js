const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split("").map(Number));

const directions = [
  [-1, 0],
  [1, 0],
  [0, -1],
  [0, 1],
];
const rMax = input.length - 1;
const cMax = input[0].length - 1;

let part1 = 0;
const candidates = [];

for (let r = 0; r < input.length; r++) {
  for (let c = 0; c < input[0].length; c++) {
    let isLowPoint = true;
    for (const [dr, dc] of directions) {
      const r2 = r + dr;
      const c2 = c + dc;
      if (r2 < 0 || r2 > rMax || c2 < 0 || c2 > cMax) continue;
      if (input[r2][c2] <= input[r][c]) {
        isLowPoint = false;
        break;
      }
    }
    if (isLowPoint) {
      part1 += input[r][c] + 1;
      candidates.push([r, c]);
    }
  }
}

console.log(part1);

function mapBasin(row, col) {
  const queue = [[row, col]];
  const visited = new Set([`${row},${col}`]);
  let size = 0;
  for (let i = 0; i < queue.length; i++) {
    size++;
    const [r, c] = queue[i];
    for (const [dr, dc] of directions) {
      const r2 = r + dr;
      const c2 = c + dc;
      if (r2 < 0 || r2 > rMax || c2 < 0 || c2 > cMax) continue;
      if (input[r2][c2] === 9) continue;
      if (visited.has(`${r2},${c2}`)) continue;
      visited.add(`${r2},${c2}`);
      queue.push([r2, c2]);
    }
  }
  return size;
}

const basinSizes = [];
for (const [r, c] of candidates) {
  basinSizes.push(mapBasin(r, c));
}

basinSizes.sort((a, b) => b - a);

console.log(basinSizes.slice(0, 3).reduce((a, b) => a * b, 1));
