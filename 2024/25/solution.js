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
  for (let r = 0; r < block.length; r++) {
    const line = [];
    for (let c = 0; c < block[r].length; c++) {
      if (block[r][c] === "#") line.push("1");
      else line.push("0");
    }
    design.push(parseInt(line.join(""), 2));
  }
  if (block[0][0] === "#") locks.push(design);
  else keys.push(design);
}

let part1 = 0;
for (const lock of locks) {
  for (const key of keys) {
    if (lock.every((a, i) => (a & key[i]) === 0)) part1++;
  }
}
console.log(part1);
