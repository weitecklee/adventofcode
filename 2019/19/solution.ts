import * as fs from "fs";
import * as path from "path";
import intcodeGenerator from "../intcode/intcode";

const puzzleInput = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split(",")
  .map(Number);

function isAffectedPoint(x: number, y: number): boolean {
  const tractor = intcodeGenerator(puzzleInput);
  tractor.next();
  tractor.next(x);
  return tractor.next(y).value === 1;
}

let part1 = 0;
for (let x = 0; x < 50; x++) {
  for (let y = 0; y < 50; y++) {
    if (isAffectedPoint(x, y)) part1++;
  }
}
console.log(part1);

const beamEnds: Map<number, number[]> = new Map();
// Map with key y and values [xMin, xMax] to keep track of range of beam at each y

let y = 100;
let xMin = 0;
while (!isAffectedPoint(xMin, y)) {
  xMin++;
}
let xMax = xMin + 1;
while (isAffectedPoint(xMax, y)) {
  xMax++;
}
xMax--;
beamEnds.set(y, [xMin, xMax]);

while (true) {
  y++;
  while (isAffectedPoint(xMin, y)) {
    xMin--;
  }
  while (!isAffectedPoint(xMin, y)) {
    xMin++;
  }
  while (isAffectedPoint(xMax, y)) {
    xMax++;
  }
  while (!isAffectedPoint(xMax, y)) {
    xMax--;
  }
  beamEnds.set(y, [xMin, xMax]);
  if (xMax - xMin + 1 >= 100 && beamEnds.has(y - 99)) {
    const [xMin0, xMax0] = beamEnds.get(y - 99)!;
    if (
      xMin >= xMin0 &&
      xMin <= xMax0 &&
      xMin + 99 >= xMin0 &&
      xMin + 99 <= xMax0
    ) {
      console.log(xMin * 10000 + y - 99);
      break;
    }
  }
}
