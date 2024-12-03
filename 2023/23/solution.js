const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n");

const start = [input[0].indexOf("."), 0];
const end = [input[input.length - 1].indexOf("."), input.length - 1];

const queue = [[...start, new Set()]];

const directions = [
  [0, 1],
  [1, 0],
  [-1, 0],
  [0, -1],
];
const slopeTiles = ["v", ">", "<", "^"];

let part1 = 0;

while (queue.length) {
  const [x, y, visited] = queue.pop();
  if (x === end[0] && y === end[1]) {
    part1 = Math.max(part1, visited.size);
    continue;
  }
  visited.add(`${x},${y}`);
  const slopeIndex = slopeTiles.indexOf(input[y][x]);
  if (slopeIndex >= 0) {
    const [dx, dy] = directions[slopeIndex];
    const x2 = x + dx;
    const y2 = y + dy;
    queue.push([x2, y2, new Set(visited)]);
    continue;
  }
  for (let i = 0; i < directions.length; i++) {
    const x2 = x + directions[i][0];
    const y2 = y + directions[i][1];
    if (x2 < 0 || x2 >= input[0].length || y2 < 0 || y2 >= input.length)
      continue;
    if (input[y2][x2] === "#") continue;
    if (slopeTiles.indexOf(input[y2][x2]) === 3 - i) continue;
    if (visited.has(`${x2},${y2}`)) continue;
    queue.push([x2, y2, new Set(visited)]);
  }
}

console.log(part1);
