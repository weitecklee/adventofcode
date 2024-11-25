const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n");

let part1 = 0;

for (let c = 0; c < input[0].length; c++) {
  let rowAfterTilt = 0;
  for (let r = 0; r < input.length; r++) {
    if (input[r][c] === "#") {
      rowAfterTilt = r + 1;
    } else if (input[r][c] === "O") {
      part1 += input.length - rowAfterTilt;
      rowAfterTilt++;
    }
  }
}

console.log(part1);
