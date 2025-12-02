import * as fs from "fs";
import * as path from "path";

const puzzleInput = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split(",");

function parseInput(data: string[]): number[][] {
  return data.map((s) => s.split("-").map(Number));
}

function part1(ranges: number[][]): number {
  let res = 0;
  for (const range of ranges) {
    for (let i = range[0]; i <= range[1]; i++) {
      if (isInvalidId(i)) res += i;
    }
  }
  return res;
}

function part2(ranges: number[][]): number {
  let res = 0;
  for (const range of ranges) {
    for (let i = range[0]; i <= range[1]; i++) {
      if (isInvalidId2(i)) res += i;
    }
  }
  return res;
}

// So much easier when the regexp engine supports backreferences!
// Looking at you, Go
const reg1 = /^(.+)\1$/g;
const reg2 = /^(.+)\1+$/g;

function isInvalidId(n: number): boolean {
  return reg1.test(n.toString());
}

function isInvalidId2(n: number): boolean {
  return reg2.test(n.toString());
}

const ranges = parseInput(puzzleInput);
console.log(part1(ranges));
console.log(part2(ranges));
