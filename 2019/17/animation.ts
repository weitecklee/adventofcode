import * as fs from "fs";
import * as path from "path";
import intcodeGenerator from "../intcode/intcode";

const puzzleInput = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split(",")
  .map(Number);

const camera = intcodeGenerator(puzzleInput);
const scaffold: string[][] = [];
let row: string[] = [];

while (true) {
  const ret = camera.next();
  if (ret.done) break;
  if (ret.value === 10) {
    scaffold.push(row);
    row = [];
  } else {
    row.push(String.fromCharCode(ret.value));
  }
}

const tileBlocks: Map<string, string> = new Map([
  ["#", "█"],
  [".", " "],
]);

const directions = [
  [-1, 0],
  [1, 0],
  [0, -1],
  [0, 1],
];
const directionsASCII = ["^", "v", "<", ">"];

for (const c of directionsASCII) {
  tileBlocks.set(c, c);
}

let robotPos: number[] = [];
let robotDir: number[] = [];
const rMax = scaffold.length;
const cMax = scaffold[0].length;
for (let r = 1; r < rMax - 1; r++) {
  for (let c = 1; c < cMax - 1; c++) {
    if (scaffold[r][c] === ".") continue;
    if (directionsASCII.includes(scaffold[r][c])) {
      robotPos = [r, c];
      robotDir = directions[directionsASCII.indexOf(scaffold[r][c])];
    }
  }
}

const robotPath: (string | number)[] = [];
let steps = 0;
let endOfPath = false;
while (!endOfPath) {
  let [r2, c2] = [robotPos[0] + robotDir[0], robotPos[1] + robotDir[1]];
  if (
    r2 >= 0 &&
    c2 >= 0 &&
    r2 < rMax &&
    c2 < cMax &&
    scaffold[r2][c2] === "#"
  ) {
    steps++;
  } else {
    robotPath.push(steps);
    steps = 1;
    robotDir = [robotDir[1], -robotDir[0]];
    [r2, c2] = [robotPos[0] + robotDir[0], robotPos[1] + robotDir[1]];
    if (
      r2 >= 0 &&
      c2 >= 0 &&
      r2 < rMax &&
      c2 < cMax &&
      scaffold[r2][c2] === "#"
    ) {
      robotPath.push("R");
    } else {
      robotDir = [-robotDir[0], -robotDir[1]];
      [r2, c2] = [robotPos[0] + robotDir[0], robotPos[1] + robotDir[1]];
      if (
        r2 >= 0 &&
        c2 >= 0 &&
        r2 < rMax &&
        c2 < cMax &&
        scaffold[r2][c2] === "#"
      ) {
        robotPath.push("L");
      } else {
        endOfPath = true;
      }
    }
  }
  robotPos = [r2, c2];
}
if (robotPath[0] === 0) robotPath.shift();

const pathString = robotPath.join(",");

function generateMovementFunctions(pathString: string): string[] {
  const reg =
    /^(.{1,19}[^,])(?:,|\1)*(.{1,19}[^,])(?:,|\1|\2)*(.{1,19}[^,])(?:,|\1|\2|\3)*$/;

  const match = pathString.match(reg);
  if (match) {
    const funcA = match[1];
    const funcB = match[2];
    const funcC = match[3];
    const funcMain = pathString
      .replace(new RegExp(funcA, "g"), "A")
      .replace(new RegExp(funcB, "g"), "B")
      .replace(new RegExp(funcC, "g"), "C");
    return [funcMain, funcA, funcB, funcC];
  }

  throw new Error("Could not generate movement functions with regex");
}

const [funcMain, funcA, funcB, funcC] = generateMovementFunctions(pathString);

puzzleInput[0] = 2;

const robot = intcodeGenerator(puzzleInput);

function displayMessage() {
  while (true) {
    const ret = robot.next();
    if (ret.value === -9999) break;
  }
}

function inputFunction(funcString: string) {
  displayMessage();
  for (let i = 0; i < funcString.length; i++) {
    robot.next(funcString.charCodeAt(i));
  }
  robot.next(10);
}

inputFunction(funcMain);
inputFunction(funcA);
inputFunction(funcB);
inputFunction(funcC);
displayMessage();
robot.next("y".charCodeAt(0));
robot.next(10);

console.log("\x1b[?25l");

function draw() {
  const messages: string[] = [];
  let message: string[] = [];
  while (true) {
    const ret = robot.next();
    if (ret.done) {
      console.log(`\x1b[${rMax}B\x1b[?${cMax}h`);
      return;
    }
    if (ret.value === 10) {
      if (message.length === 0) break;
      messages.push(message.join(""));
      message = [];
    } else {
      message.push(tileBlocks.get(String.fromCharCode(ret.value))!);
    }
  }
  console.log(messages.join("\n"));
  console.log(`\x1b[${rMax}A\x1b[${cMax}D`);
  setTimeout(draw, 1000 / 24);
}

draw();

function cleanup() {
  console.log(`\x1b[${rMax}B\x1b[?25h`);
  process.exit(0);
}

process.on("SIGINT", cleanup);
process.on("SIGTERM", cleanup);
