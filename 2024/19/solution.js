const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n\n");

const patterns = new Set(input[0].split(", "));
const memo = new Map([["", true]]);
const memo2 = new Map([["", 1]]);

function recur(design) {
  if (memo.has(design)) return memo.get(design);
  for (let i = 0; i < design.length; i++) {
    if (patterns.has(design.slice(0, i + 1))) {
      if (recur(design.slice(i + 1))) {
        memo.set(design, true);
        return true;
      }
    }
  }
  memo.set(design, false);
  return false;
}

function recur2(design) {
  if (memo2.has(design)) return memo2.get(design);
  let count = 0;
  for (let i = 0; i < design.length; i++) {
    if (patterns.has(design.slice(0, i + 1))) {
      count += recur2(design.slice(i + 1));
    }
  }
  memo2.set(design, count);
  return count;
}

let part1 = 0;
let part2 = 0;
for (const design of input[1].split("\n")) {
  if (recur(design)) {
    part1++;
    part2 += recur2(design);
  }
}
console.log(part1);
console.log(part2);
