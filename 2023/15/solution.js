const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split(",");

function executeHASH(s) {
  let curr = 0;
  for (const c of s) {
    curr += c.charCodeAt(0);
    curr *= 17;
    curr %= 256;
  }
  return curr;
}

const part1 = input.reduce((a, b) => a + executeHASH(b), 0);
console.log(part1);
