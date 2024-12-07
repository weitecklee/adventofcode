const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n");

const regex = /mul\((\d{1,3}),(\d{1,3})\)|do(?:n\'t)?\(\)/g;

let part1 = 0;
let part2 = 0;

let doMul = true;
for (const line of input) {
  let match;
  while ((match = regex.exec(line))) {
    if (match[0] === "do()") {
      doMul = true;
    } else if (match[0] === "don't()") {
      doMul = false;
    } else {
      const prod = Number(match[1]) * Number(match[2]);
      part1 += prod;
      part2 += doMul ? prod : 0;
    }
  }
}

console.log(part1, part2);
