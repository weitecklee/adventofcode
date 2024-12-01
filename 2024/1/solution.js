const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split("   ").map(Number));

const list1 = input.map((a) => a[0]);
const list2 = input.map((a) => a[1]);
list1.sort((a, b) => a - b);
list2.sort((a, b) => a - b);

const part1 = list1.reduce((acc, cur, i) => acc + Math.abs(cur - list2[i]), 0);
console.log(part1);

const count2 = new Map();
for (const n of list2) count2.set(n, (count2.get(n) || 0) + 1);

const part2 = list1.reduce((acc, cur) => acc + cur * (count2.get(cur) || 0), 0);
console.log(part2);
