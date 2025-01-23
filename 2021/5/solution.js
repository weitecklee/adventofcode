const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8", (err, data) => {
    if (err) {
      console.log(err);
    } else {
      return data;
    }
  })
  .split("\n");

const grid = new Array(1000);
for (let i = 0; i < grid.length; i++) {
  grid[i] = new Array(1000).fill(0);
}

let count = 0;
for (const line of input) {
  const coords = line.match(/\d+/g).map(Number);
  if (coords[0] === coords[2] || coords[1] === coords[3]) {
    const a = Math.min(coords[0], coords[2]);
    const b = Math.min(coords[1], coords[3]);
    const x = Math.max(coords[0], coords[2]);
    const y = Math.max(coords[1], coords[3]);
    for (let i = a; i <= x; i++) {
      for (let j = b; j <= y; j++) {
        grid[i][j]++;
        if (grid[i][j] === 2) {
          count++;
        }
      }
    }
  }
}

console.log(count);

for (const line of input) {
  const coords = line.match(/\d+/g).map(Number);
  if (coords[0] !== coords[2] && coords[1] !== coords[3]) {
    const xstep = Math.sign(coords[2] - coords[0]);
    const ystep = Math.sign(coords[3] - coords[1]);
    for (let i = 0; i <= Math.abs(coords[2] - coords[0]); i++) {
      grid[coords[0] + i * xstep][coords[1] + i * ystep]++;
      if (grid[coords[0] + i * xstep][coords[1] + i * ystep] === 2) {
        count++;
      }
    }
  }
}

console.log(count);
