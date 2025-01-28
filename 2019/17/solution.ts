import * as fs from "fs";
import * as path from "path";
import intcodeGenerator, { IntcodeGenerator } from "../intcode";

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

const directions = [
  [-1, 0],
  [1, 0],
  [0, -1],
  [0, 1],
];
const directionsASCII = ["^", "v", "<", ">"];

let part1 = 0;
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
    let isIntersection = true;
    for (const [dr, dc] of directions) {
      if (scaffold[r + dr][c + dc] === ".") {
        isIntersection = false;
        break;
      }
    }
    if (isIntersection) part1 += r * c;
  }
}
console.log(part1);

// const robotPath: (string | number)[] = [];
// let steps = 1;
// let endOfPath = false;
// while (!endOfPath) {
//   let [r2, c2] = [robotPos[0] + robotDir[0], robotPos[1] + robotDir[1]];
//   if (
//     r2 >= 0 &&
//     c2 >= 0 &&
//     r2 < rMax &&
//     c2 < cMax &&
//     scaffold[r2][c2] === "#"
//   ) {
//     steps++;
//   } else {
//     robotPath.push(steps);
//     steps = 1;
//     robotDir = [robotDir[1], -robotDir[0]];
//     [r2, c2] = [robotPos[0] + robotDir[0], robotPos[1] + robotDir[1]];
//     if (
//       r2 >= 0 &&
//       c2 >= 0 &&
//       r2 < rMax &&
//       c2 < cMax &&
//       scaffold[r2][c2] === "#"
//     ) {
//       robotPath.push("R");
//     } else {
//       robotDir = [-robotDir[0], -robotDir[1]];
//       [r2, c2] = [robotPos[0] + robotDir[0], robotPos[1] + robotDir[1]];
//       if (
//         r2 >= 0 &&
//         c2 >= 0 &&
//         r2 < rMax &&
//         c2 < cMax &&
//         scaffold[r2][c2] === "#"
//       ) {
//         robotPath.push("L");
//       } else {
//         endOfPath = true;
//       }
//     }
//   }
//   robotPos = [r2, c2];
// }
// robotPath.shift();
// console.log(robotPath.join(","));

// R,8,L,4,R,4,R,10,R,8,R,8,L,4,R,4,R,10,R,8,L,12,L,12,R,8,R,8,R,10,R,4,R,4,L,12,L,12,R,8,R,8,R,10,R,4,R,4,L,12,L,12,R,8,R,8,R,10,R,4,R,4,R,10,R,4,R,4,R,8,L,4,R,4,R,10,R,8
// main: A,A,B,C,B,C,B,C,C,A
// A: R,8,L,4,R,4,R,10,R,8
// B: L,12,L,12,R,8,R,8
// C: R,10,R,4,R,4

puzzleInput[0] = 2;

const robot = intcodeGenerator(puzzleInput);

const funcMain = "A,A,B,C,B,C,B,C,C,A";
const funcA = "R,8,L,4,R,4,R,10,R,8";
const funcB = "L,12,L,12,R,8,R,8";
const funcC = "R,10,R,4,R,4";

function displayMessage() {
  // let message: string[] = [];
  while (true) {
    const ret = robot.next();
    if (ret.value === -9999) break;
    // message.push(String.fromCharCode(ret.value));
  }
  // console.log(message.join(""));
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
robot.next("n".charCodeAt(0));
// robot.next("y".charCodeAt(0));
robot.next(10);
let part2 = 0;
// let message: string[] = [];
while (true) {
  const ret = robot.next();
  if (ret.done) break;
  // if (ret.value === 10) {
  //   console.log(message.join(""));
  //   message = [];
  // } else {
  //   message.push(String.fromCharCode(ret.value));
  // }
  part2 = ret.value;
}
console.log(part2);
