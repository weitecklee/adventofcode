import * as fs from "fs";
import * as path from "path";

const puzzleInput: [string, ...number[]][] = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => {
    const nums = a.match(/\d+/g)!.map(Number);
    return [a[0], ...nums];
  });

// let xMin = Number.MAX_SAFE_INTEGER;
// let xMax = -1;

let yMin = Number.MAX_SAFE_INTEGER;
let yMax = -1;
const clayCoords: Set<string> = new Set();

for (const [ch, num0, num1, num2] of puzzleInput) {
  if (ch === "x") {
    yMin = Math.min(yMin, num1);
    yMax = Math.max(yMax, num2);
    // xMin = Math.min(num0, xMin);
    // xMax = Math.max(num0, xMax);
    for (let i = num1; i <= num2; i++) {
      clayCoords.add(`${num0},${i}`);
    }
  } else {
    yMin = Math.min(yMin, num0);
    yMax = Math.max(yMax, num0);
    // xMin = Math.min(num1, xMin);
    // xMax = Math.max(num2, xMax);
    for (let i = num1; i <= num2; i++) {
      clayCoords.add(`${i},${num0}`);
    }
  }
}

const settledCoords: Set<string> = new Set();
const flowingCoords: Set<string> = new Set();
const queue = [[500, yMin]];

while (queue.length) {
  let [x, y] = queue.pop()!;
  while (
    !flowingCoords.has(`${x},${y}`) &&
    !clayCoords.has(`${x},${y}`) &&
    y <= yMax
  ) {
    flowingCoords.add(`${x},${y}`);
    y++;
  }
  if (flowingCoords.has(`${x},${y}`)) continue;
  if (y > yMax) continue;
  y--;
  while (true) {
    const currSet: Set<string> = new Set();
    let leftWall = false;
    let rightWall = false;
    let xl = x;
    let xr = x;
    while (
      !clayCoords.has(`${xl},${y}`) &&
      (clayCoords.has(`${xl},${y + 1}`) || settledCoords.has(`${xl},${y + 1}`))
    ) {
      currSet.add(`${xl},${y}`);
      xl--;
    }
    if (clayCoords.has(`${xl},${y}`)) leftWall = true;
    else queue.push([xl, y]);
    while (
      !clayCoords.has(`${xr},${y}`) &&
      (clayCoords.has(`${xr},${y + 1}`) || settledCoords.has(`${xr},${y + 1}`))
    ) {
      currSet.add(`${xr},${y}`);
      xr++;
    }
    if (clayCoords.has(`${xr},${y}`)) rightWall = true;
    else queue.push([xr, y]);
    if (leftWall && rightWall) {
      for (const c of currSet) settledCoords.add(c);
      y--;
    } else {
      for (const c of currSet) flowingCoords.add(c);
      break;
    }
  }
}

const waterCoords = flowingCoords.union(settledCoords);
console.log(waterCoords.size);
console.log(settledCoords.size);

// for (const c of waterCoords) {
//   const [x, y] = c.split(",").map(Number);
//   xMin = Math.min(x, xMin);
//   xMax = Math.max(x, xMax);
// }

// fs.writeFileSync("out.txt", "");

// for (let y = yMin; y <= yMax; y++) {
//   const row: string[] = [];
//   for (let x = xMin; x <= xMax; x++) {
//     let c = ".";
//     if (clayCoords.has(`${x},${y}`)) c = "#";
//     else if (settledCoords.has(`${x},${y}`)) c = "~";
//     else if (flowingCoords.has(`${x},${y}`)) c = "|";
//     row.push(c);
//   }
//   row.push(`\n`);
//   fs.appendFileSync("out.txt", row.join(""));
// }
