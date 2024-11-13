const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map(Number);

const part1 = input.reduce((a, b, i) => a + (i > 0 && b > input[i - 1]), 0);

console.log(part1);

// let sumA = input[0] + input[1] + input[2];
// let sumB = input[1] + input[2] + input[3];

// let part2 = sumB > sumA ? 1 : 0;

// for (let i = 4; i < input.length; i++) {
//   sumA += input[i - 1] - input[i - 4];
//   sumB += input[i] - input[i - 3];
//   part2 += sumB > sumA;
// }

// only need to compare input[i] and input[i - 3]
let part2 = 0;
for (let i = 3; i < input.length; i++) {
  part2 += input[i] > input[i - 3];
}

console.log(part2);
