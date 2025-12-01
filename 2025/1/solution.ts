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
  for (let turn of turns) {
    let inc = Math.sign(turn);
    turn = Math.abs(turn);
    for (let i = 0; i < turn; i++) {
      dial += inc;
      dial %= 100;
      if (dial === 0) {
        res++;
      }
    }
  }
  return res;
}

const turns = parseInput(puzzleInput);
console.log(part1(turns));
console.log(part2(turns));
