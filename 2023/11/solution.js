const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split(""));

const emptyRows = [];
const emptyCols = [];
const galaxies = [];

for (let c = 0; c < input[0].length; c++) {
  let galaxyFound = false;
  for (let r = 0; r < input.length; r++) {
    if (input[r][c] === "#") {
      galaxies.push([r, c]);
      galaxyFound = true;
    }
  }
  if (!galaxyFound) emptyCols.push(c);
}

for (let r = 0; r < input.length; r++) {
  if (!input[r].includes("#")) emptyRows.push(r);
}

let part1 = 0;
let part2 = 0;

for (let i = 0; i < galaxies.length; i++) {
  for (let j = i + 1; j < galaxies.length; j++) {
    let [r1, c1] = galaxies[i];
    let [r2, c2] = galaxies[j];
    [r1, r2] = [Math.min(r1, r2), Math.max(r1, r2)];
    [c1, c2] = [Math.min(c1, c2), Math.max(c1, c2)];
    const expandedRows = emptyRows.filter((r) => r > r1 && r < r2).length;
    const expandedCols = emptyCols.filter((c) => c > c1 && c < c2).length;
    part1 += r2 - r1 + c2 - c1 + expandedRows + expandedCols;
    part2 +=
      r2 -
      r1 +
      c2 -
      c1 +
      (1000000 - 1) * expandedRows +
      (1000000 - 1) * expandedCols;
  }
}

console.log(part1);
console.log(part2);
