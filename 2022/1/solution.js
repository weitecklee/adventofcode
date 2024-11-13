const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => (a.length ? Number(a) : ""));

input.push("");
// let part1 = 0;
// let curr = 0;
// for (const line of input) {
//   if (line === "") {
//     part1 = Math.max(part1, curr);
//     curr = 0;
//   } else {
//     curr += line;
//   }
// }
// console.log(part1);

const elves = [];
let curr = 0;
for (const line of input) {
  if (line === "") {
    elves.push(curr);
    curr = 0;
  } else {
    curr += line;
  }
}
elves.sort((a, b) => b - a);
console.log(elves[0]);
console.log(elves[0] + elves[1] + elves[2]);
