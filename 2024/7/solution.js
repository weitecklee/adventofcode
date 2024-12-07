const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split(": "))
  .map(([a, b]) => [Number(a), b.split(" ").map(Number)]);

let part1 = 0;
let part2 = 0;
for (const [target, nums] of input) {
  let queue = [nums[0]];
  for (let i = 1; i < nums.length; i++) {
    let tmp = [];
    for (const n of queue) {
      const prod = n * nums[i];
      const sum = n + nums[i];
      if (prod <= target) tmp.push(prod);
      if (sum <= target) tmp.push(sum);
    }
    queue = tmp;
  }
  if (queue.some((a) => a === target)) {
    part1 += target;
    part2 += target;
    continue;
  }
  queue = [nums[0]];
  for (let i = 1; i < nums.length; i++) {
    let tmp = [];
    for (const n of queue) {
      const prod = n * nums[i];
      const sum = n + nums[i];
      const concat = Number(n.toString() + nums[i].toString());
      if (prod <= target) tmp.push(prod);
      if (sum <= target) tmp.push(sum);
      if (concat <= target) tmp.push(concat);
    }
    queue = tmp;
  }
  if (queue.some((a) => a === target)) {
    part2 += target;
  }
}

console.log(part1);
console.log(part2);
