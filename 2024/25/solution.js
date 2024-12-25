const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n\n")
  .map((a) => a.split("\n"));

const locks = [];
const keys = [];

for (const block of input) {
  const design = [];
  for (let c = 0; c < block[0].length; c++) {
    let count = 0;
    for (let r = 0; r < block.length; r++) {
      if (block[r][c] === "#") count++;
    }
    design.push(count);
  }
  if (block[0][0] === "#") locks.push(design);
  else keys.push(design);
}

let part1 = 0;
for (const lock of locks) {
  for (const key of keys) {
    let fit = true;
    for (let i = 0; i < lock.length; i++) {
      if (lock[i] + key[i] > 7) {
        fit = false;
        break;
      }
    }
    if (fit) part1++;
  }
}

console.log(part1);
