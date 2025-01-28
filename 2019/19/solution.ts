import * as fs from "fs";
import * as path from "path";
import intcodeGenerator from "../intcode";

const puzzleInput = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split(",")
  .map(Number);

let part1 = 0;
for (let x = 0; x < 50; x++) {
  for (let y = 0; y < 50; y++) {
    const tractor = intcodeGenerator(puzzleInput);
    tractor.next();
    tractor.next(x);
    const ret = tractor.next(y);
    if (ret.value === 1) part1++;
  }
}
console.log(part1);
