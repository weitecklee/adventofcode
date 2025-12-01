import * as fs from "fs";
import * as path from "path";

const puzzleInput = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n");

function parseInput(data: string[]): number[] {
  const res: number[] = [];
  for (const line of data) {
    let n = Number(line.slice(1));
    if (line[0] === "L") {
      n *= -1;
    }
    res.push(n);
  }
  return res;
}

function part1(turns: number[]): number {
  let dial = 50;
  let res = 0;
  for (const turn of turns) {
    dial += turn;
    dial %= 100;
    if (dial === 0) {
      res++;
    }
  }
  return res;
}

function part2(turns: number[]): number {
  let dial = 50;
  let res = 0;
  let prev = dial;
  for (let turn of turns) {
    res += Math.floor(Math.abs(turn) / 100);
    dial += turn % 100;
    if ((prev < 0 && dial >= 0) || (prev > 0 && dial <= 0)) {
      res++;
    }
    if (prev < 100 && dial >= 100) {
      res++;
      dial -= 100;
    }
    if (prev > -100 && dial <= -100) {
      res++;
      dial += 100;
    }
    prev = dial;
  }
  return res;
}

const turns = parseInput(puzzleInput);
console.log(part1(turns));
console.log(part2(turns));
