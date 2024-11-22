const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((line) => line.split(" ").map(Number));

function extrapolation(seq) {
  const arr = [seq];
  while (arr[arr.length - 1].some((a) => a != 0)) {
    const curr = arr[arr.length - 1];
    const next = [];
    for (let i = 1; i < curr.length; i++) {
      next.push(curr[i] - curr[i - 1]);
    }
    arr.push(next);
  }
  for (let i = arr.length - 2; i >= 0; i--) {
    arr[i].push(arr[i][arr[i].length - 1] + arr[i + 1][arr[i + 1].length - 1]);
    arr[i].unshift(arr[i][0] - arr[i + 1][0]);
  }
  return [arr[0][0], arr[0][arr[0].length - 1]];
}

const res = input.map(extrapolation);
const [part2, part1] = res.reduce((a, b) => [a[0] + b[0], a[1] + b[1]], [0, 0]);
console.log(part1, part2);
