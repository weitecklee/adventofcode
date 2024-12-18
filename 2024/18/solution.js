const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n");

const xMax = 70;
const yMax = 70;
const directions = [
  [0, 1],
  [1, 0],
  [0, -1],
  [-1, 0],
];

const memorySpace = new Set();

for (let i = 0; i < 1024; i++) {
  memorySpace.add(input[i]);
}

function findPath() {
  const queue = [[0, [0, 0]]];
  const visited = new Set(memorySpace);
  let i = 0;
  while (i < queue.length) {
    const [steps, [x, y]] = queue[i];
    i++;
    if (x === xMax && y === yMax) {
      return steps;
      break;
    }
    for (const [dx, dy] of directions) {
      const nx = x + dx;
      const ny = y + dy;
      if (nx < 0 || nx > xMax || ny < 0 || ny > yMax) continue;
      const coord = `${nx},${ny}`;
      if (!visited.has(coord)) {
        visited.add(coord);
        queue.push([steps + 1, [nx, ny]]);
      }
    }
  }
  return -1;
}

console.log(findPath());

for (let j = 1024; j < input.length; j++) {
  memorySpace.add(input[j]);
  if (findPath() === -1) {
    console.log(input[j]);
    break;
  }
}
