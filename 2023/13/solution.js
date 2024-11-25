const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n\n")
  .map((a) => a.split("\n"));

class Pattern {
  constructor(input) {
    this.pattern = input;
    [this.reflectionType, this.reflectionRow, this.reflectionCol] =
      this.findReflection();
  }

  findReflection() {
    for (let col = 1; col < this.pattern[0].length; col++) {
      let reflection = true;
      for (let row = 0; row < this.pattern.length; row++) {
        for (let i = 0; i < col && i + col < this.pattern[0].length; i++) {
          if (this.pattern[row][col - i - 1] !== this.pattern[row][i + col]) {
            reflection = false;
            break;
          }
        }
        if (!reflection) break;
      }
      if (reflection) return ["col", 0, col];
    }
    for (let row = 1; row < this.pattern.length; row++) {
      let reflection = true;
      for (let col = 0; col < this.pattern[0].length; col++) {
        for (let i = 0; i < row && i + row < this.pattern.length; i++) {
          if (this.pattern[row - i - 1][col] !== this.pattern[i + row][col]) {
            reflection = false;
            break;
          }
        }
        if (!reflection) break;
      }
      if (reflection) return ["row", row, 0];
    }
  }
}
const patterns = input.map((a) => new Pattern(a));
const part1 = patterns.reduce(
  (a, b) => a + 100 * b.reflectionRow + b.reflectionCol,
  0
);
console.log(part1);
