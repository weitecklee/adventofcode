const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n");

const rMax = input.length - 1;
const cMax = input[0].length - 1;

let eastCucumbers = [];
let southCucumbers = [];
let eastMap = new Set();
let southMap = new Set();

for (let r = 0; r < input.length; r++) {
  for (let c = 0; c < input[0].length; c++) {
    if (input[r][c] === ">") {
      eastCucumbers.push([r, c]);
      eastMap.add(`${r},${c}`);
    } else if (input[r][c] === "v") {
      southCucumbers.push([r, c]);
      southMap.add(`${r},${c}`);
    }
  }
}

let isMoving = true;
let part1 = 0;
while (isMoving) {
  part1++;
  isMoving = false;
  const eastCucumbers2 = [];
  const eastMap2 = new Set();
  for (const [r, c] of eastCucumbers) {
    let c2 = c + 1;
    if (c2 > cMax) c2 = 0;
    const coord = `${r},${c2}`;
    if (eastMap.has(coord) || southMap.has(coord)) {
      eastCucumbers2.push([r, c]);
      eastMap2.add(`${r},${c}`);
    } else {
      isMoving = true;
      eastCucumbers2.push([r, c2]);
      eastMap2.add(coord);
    }
  }
  eastCucumbers = eastCucumbers2;
  eastMap = eastMap2;
  const southCucumbers2 = [];
  const southMap2 = new Set();
  for (const [r, c] of southCucumbers) {
    let r2 = r + 1;
    if (r2 > rMax) r2 = 0;
    const coord = `${r2},${c}`;
    if (eastMap.has(coord) || southMap.has(coord)) {
      southCucumbers2.push([r, c]);
      southMap2.add(`${r},${c}`);
    } else {
      isMoving = true;
      southCucumbers2.push([r2, c]);
      southMap2.add(coord);
    }
  }
  southCucumbers = southCucumbers2;
  southMap = southMap2;
}

console.log(part1);
