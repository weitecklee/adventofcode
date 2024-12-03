const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n");

const mulRegex = /mul\((\d+),(\d+)\)/g;

let part1 = 0;

for (const line of input) {
  let match;
  while ((match = mulRegex.exec(line))) {
    part1 += Number(match[1]) * Number(match[2]);
  }
}

console.log(part1);

const doOrMulRegex = /mul\((\d+),(\d+)\)|do(n\'t)?\(\)/g;

let part2 = 0;

let doMul = true;
for (const line of input) {
  let match;
  while ((match = doOrMulRegex.exec(line))) {
    if (match[0] === "do()") {
      doMul = true;
    } else if (match[0] === "don't()") {
      doMul = false;
    } else {
      part2 += doMul ? Number(match[1]) * Number(match[2]) : 0;
    }
  }
}

console.log(part2);
