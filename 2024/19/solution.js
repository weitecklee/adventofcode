const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n\n");

const patterns = new Set(input[0].split(", "));
const memo = new Map();

function recur(design) {
  if (memo.has(design)) return memo.get(design);
  for (let i = 0; i < design.length; i++) {
    if (patterns.has(design.slice(0, i + 1))) {
      if (i === design.length - 1) return true;
      if (recur(design.slice(i + 1))) {
        memo.set(design, true);
        return true;
      }
    }
  }
  memo.set(design, false);
  return false;
}

let part1 = 0;
for (const design of input[1].split("\n")) {
  const possible = recur(design);
  if (possible) part1++;
}
console.log(part1);
