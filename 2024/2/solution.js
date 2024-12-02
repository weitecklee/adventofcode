const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split(" ").map(Number));

function isSafe(level) {
  const isIncreasing = level[1] - level[0] > 0;
  for (let i = 1; i < level.length; i++) {
    const diff = Math.abs(level[i] - level[i - 1]);
    if (diff < 1 || diff > 3) return false;
    if (level[i] - level[i - 1] > 0 !== isIncreasing) return false;
  }
  return true;
}

function isSafeWithTolerance(level) {
  if (isSafe(level)) return true;
  for (let i = 0; i < level.length; i++) {
    if (isSafe(level.toSpliced(i, 1))) return true;
  }
  return false;
}

const part1 = input.filter(isSafe).length;
console.log(part1);
const part2 = input.filter(isSafeWithTolerance).length;
console.log(part2);
