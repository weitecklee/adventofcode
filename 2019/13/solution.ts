import * as fs from "fs";
import * as path from "path";
import intcodeGenerator from "../intcode/intcode";

const puzzleInput = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split(",")
  .map(Number);

const arcadeCabinet = intcodeGenerator(puzzleInput);

let part1 = 0;

while (true) {
  if (arcadeCabinet.next().done) break;
  arcadeCabinet.next();
  const tileRet = arcadeCabinet.next();
  if (tileRet.value === 2) part1++;
}

console.log(part1);

const freePlayInput = puzzleInput.slice();
freePlayInput[0] = 2;

const freePlayArcade = intcodeGenerator(freePlayInput);
let move = 0;
let xBall = 0;
let xPaddle = 0;
let part2 = 0;
while (true) {
  let xRet: IteratorResult<number, number>;
  while (true) {
    xRet = freePlayArcade.next(move);
    if (xRet.done || xRet.value === -9999) break;
    const yRet = freePlayArcade.next();
    const tileRet = freePlayArcade.next();
    if (xRet.value === -1 && yRet.value === 0) {
      part2 = tileRet.value;
    }
    if (tileRet.value === 4) {
      xBall = xRet.value;
    } else if (tileRet.value === 3) {
      xPaddle = xRet.value;
    }
  }
  if (xRet.done) break;
  if (xBall < xPaddle) move = -1;
  else if (xBall > xPaddle) move = 1;
  else move = 0;
}
console.log(part2);
