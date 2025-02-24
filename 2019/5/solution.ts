import * as fs from "fs";
import * as path from "path";
import intcodeGenerator from "../intcode/intcode";

const puzzleInput = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split(",")
  .map(Number);

const intcode = intcodeGenerator(puzzleInput);
const intcode2 = intcodeGenerator(puzzleInput);
intcode.next();
intcode.next(1);
let part1;
while (true) {
  const ret = intcode.next();
  if (ret.done) break;
  part1 = ret.value;
}
console.log(part1);
intcode2.next();
console.log(intcode2.next(5).value);
