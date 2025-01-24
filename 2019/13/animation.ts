import * as fs from "fs";
import * as path from "path";
import intcodeGenerator from "../intcode";

const puzzleInput = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split(",")
  .map(Number);

puzzleInput[0] = 2;

const freePlayArcade = intcodeGenerator(puzzleInput);
const arcadeScreen = Array(24)
  .fill("")
  .map(() => Array(44).fill(" "));
let move = 0;
let xBall = 0;
let xPaddle = 0;
let score = 0;
let xRet: IteratorResult<number, number>;
const tileBlocks = new Map([
  [0, " "],
  [1, "█"],
  [2, "▒"],
  [3, "▔"],
  [4, "O"],
]);
console.log("\x1b[?25l");

function draw() {
  while (true) {
    xRet = freePlayArcade.next(move);
    if (xRet.done || xRet.value === -9999) break;
    const yRet = freePlayArcade.next();
    const tileRet = freePlayArcade.next();
    if (xRet.value === -1 && yRet.value === 0) {
      score = tileRet.value;
    }
    if (tileRet.value === 4) {
      xBall = xRet.value;
    } else if (tileRet.value === 3) {
      xPaddle = xRet.value;
    }
    if (xRet.value >= 0 && yRet.value >= 0)
      arcadeScreen[yRet.value][xRet.value] = tileBlocks.get(tileRet.value);
  }

  if (xBall < xPaddle) move = -1;
  else if (xBall > xPaddle) move = 1;
  else move = 0;

  arcadeScreen.forEach((a) => console.log(a.join("")));
  console.log(score);
  console.log("\x1b[26A\x1b[44D");

  if (!xRet.done) setTimeout(draw, 1000 / 24);
  else console.log("\x1b[26B\x1b[?25h");
}

function cleanup() {
  console.log("\x1b[26B\x1b[?25h");
  process.exit(0);
}

draw();

process.on("SIGINT", cleanup);
process.on("SIGTERM", cleanup);
