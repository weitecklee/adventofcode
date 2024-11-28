const fs = require("fs");
const path = require("path");
const MinHeap = require("../../utils/MinHeap");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n");

class Instruction {
  constructor(line) {
    const [a, b, c] = line.split(" ");
    this.dir = a;
    this.dist = Number(b);
    this.color = c.slice(1, -1);
  }
}

const pos = [0, 0];
const instructions = input.map((line) => new Instruction(line));
let minX = 0;
let minY = 0;
let maxX = 0;
let maxY = 0;
for (const instruction of instructions) {
  switch (instruction.dir) {
    case "U":
      pos[1] -= instruction.dist;
      break;
    case "D":
      pos[1] += instruction.dist;
      break;
    case "R":
      pos[0] += instruction.dist;
      break;
    case "L":
      pos[0] -= instruction.dist;
      break;
  }
  minX = Math.min(minX, pos[0]);
  minY = Math.min(minY, pos[1]);
  maxX = Math.max(maxX, pos[0]);
  maxY = Math.max(maxY, pos[1]);
}

const width = maxX - minX + 1;
const height = maxY - minY + 1;
const grid = Array(height)
  .fill(0)
  .map(() => Array(width).fill("."));

[pos[0], pos[1]] = [0 - minX, 0 - minY];

for (const instruction of instructions) {
  switch (instruction.dir) {
    case "U":
      for (let i = 0; i < instruction.dist; i++) {
        pos[1]--;
        grid[pos[1]][pos[0]] = "#";
      }
      break;
    case "D":
      for (let i = 0; i < instruction.dist; i++) {
        pos[1]++;
        grid[pos[1]][pos[0]] = "#";
      }
      break;
    case "R":
      for (let i = 0; i < instruction.dist; i++) {
        pos[0]++;
        grid[pos[1]][pos[0]] = "#";
      }
      break;
    case "L":
      for (let i = 0; i < instruction.dist; i++) {
        pos[0]--;
        grid[pos[1]][pos[0]] = "#";
      }
      break;
  }
}

let groundCount = 0;

function fillIn(x, y) {
  let count = 0;
  const queue = [[x, y]];
  while (queue.length) {
    const [x, y] = queue.pop();
    if (x < 0 || x >= width || y < 0 || y >= height) {
      continue;
    }
    if (grid[y][x] === "#") {
      continue;
    }
    if (grid[y][x] === "-") {
      continue;
    }
    if (grid[y][x] === ".") {
      count++;
    }
    grid[y][x] = "-";
    queue.push([x - 1, y]);
    queue.push([x + 1, y]);
    queue.push([x, y - 1]);
    queue.push([x, y + 1]);
  }
  return count;
}

for (let i = 0; i < height; i++) {
  if (grid[i][0] === ".") {
    groundCount += fillIn(0, i);
  }
}

for (let i = 0; i < height; i++) {
  if (grid[i][width - 1] === ".") {
    groundCount += fillIn(width - 1, i);
  }
}

for (let i = 0; i < width; i++) {
  if (grid[0][i] === ".") {
    groundCount += fillIn(i, 0);
  }
}

for (let i = 0; i < width; i++) {
  if (grid[height - 1][i] === ".") {
    groundCount += fillIn(i, height - 1);
  }
}

console.log(height * width - groundCount);
