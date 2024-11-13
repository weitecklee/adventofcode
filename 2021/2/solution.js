const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split(" "))
  .map(([a, b]) => [a, Number(b)]);

let hor = 0;
let dep = 0;
for (const line of input) {
  switch (line[0]) {
    case "forward":
      hor += line[1];
      break;
    case "down":
      dep += line[1];
      break;
    case "up":
      dep -= line[1];
      break;
  }
}
const part1 = hor * dep;
console.log(part1);

hor = 0;
dep = 0;
let aim = 0;
for (const line of input) {
  switch (line[0]) {
    case "forward":
      hor += line[1];
      dep += line[1] * aim;
      break;
    case "down":
      aim += line[1];
      break;
    case "up":
      aim -= line[1];
      break;
  }
}
const part2 = hor * dep;
console.log(part2);
