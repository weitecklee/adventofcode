const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map(Number);

input.push(0);
input.sort((a, b) => a - b);

const diffs = [0, 0, 1];

for (let i = 1; i < input.length; i++) {
  diffs[input[i] - input[i - 1] - 1]++;
}

const part1 = diffs[0] * diffs[2];
console.log(part1);

const arrangementMap = new Map();
arrangementMap.set(input[input.length - 1], 1);
for (let i = input.length - 2; i >= 0; i--) {
  let count = 0;
  for (let j = 1; j <= 3; j++) {
    count += arrangementMap.has(input[i] + j)
      ? arrangementMap.get(input[i] + j)
      : 0;
  }
  arrangementMap.set(input[i], count);
}
console.log(arrangementMap.get(0));
