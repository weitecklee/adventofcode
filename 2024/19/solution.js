const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n\n");

const patterns = new Set(input[0].split(", "));
const designs = input[1].split("\n");
const memo = new Map([["", 1]]);

function countWays(design) {
  if (memo.has(design)) return memo.get(design);
  let count = 0;
  for (let i = 0; i < design.length; i++) {
    if (patterns.has(design.slice(0, i + 1))) {
      count += countWays(design.slice(i + 1));
    }
  }
  memo.set(design, count);
  return count;
}

let part1 = 0;
let part2 = 0;

for (const design of designs) {
  const ways = countWays(design);
  part1 += ways > 0 ? 1 : 0;
  part2 += ways;
}

console.log(part1);
console.log(part2);
