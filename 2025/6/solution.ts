import * as fs from "fs";
import * as path from "path";

const puzzleInput = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n");

const numRegex = /\d+/g;
const opRegex = /[\+\*]/g;

function parseInput(data: string[]): [number[][], string[]] {
  const opsString = data[data.length - 1];
  const ops = opsString.match(opRegex)!;

  const numbers: number[][] = [];
  for (const line of data.slice(0, -1)) {
    numbers.push(line.match(numRegex)!.map(Number));
  }

  return [numbers, ops];
}

function part1(numbers: number[][], ops: string[]): number {
  let res = 0;
  for (let i = 0; i < ops.length; i++) {
    let curr = 0;
    if (ops[i] == "*") {
      curr = 1;
      for (const nums of numbers) {
        curr *= nums[i];
      }
    } else {
      for (const nums of numbers) {
        curr += nums[i];
      }
    }
    res += curr;
  }
  return res;
}

function part2(puzzleInput: string[]): number {
  let opsString = puzzleInput[puzzleInput.length - 1];
  const numStrings = puzzleInput.slice(0, -1);

  let maxLen = 0;
  for (const line of numStrings) {
    maxLen = Math.max(maxLen, line.length);
  }

  for (let i = 0; i < numStrings.length; i++) {
    numStrings[i] = numStrings[i].padEnd(maxLen, " ");
  }
  opsString = opsString.padEnd(maxLen, " ");

  let res = 0;
  let currNums = [];
  for (let i = maxLen - 1; i >= 0; i--) {
    let curr = "";
    for (let j = 0; j < numStrings.length; j++) {
      curr += numStrings[j][i];
    }
    currNums.push(Number(curr.trim()));
    if (opsString[i] !== " ") {
      if (opsString[i] === "*") {
        res += currNums.reduce((a, b) => a * b, 1);
      } else if (opsString[i] === "+") {
        res += currNums.reduce((a, b) => a + b);
      }
      currNums = [];
      i--;
    }
  }

  return res;
}

const [numbers, ops] = parseInput(puzzleInput);
console.log(part1(numbers, ops));
console.log(part2(puzzleInput));
