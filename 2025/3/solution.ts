import * as fs from "fs";
import * as path from "path";

const puzzleInput = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n");

function parseInput(data: string[]): number[][] {
  return data.map((s) => s.split("").map(Number));
}

function part1(ranges: number[][]): number {
  return ranges.reduce((a, b) => a + findLargestJoltage(b, 2), 0);
}

function part2(ranges: number[][]): number {
  return ranges.reduce((a, b) => a + findLargestJoltage(b, 12), 0);
}

function findLargestJoltage(bank: number[], windowLength: number): number {
  let res = 0;
  let idx = 0;
  for (let i = 0; i < windowLength; i++) {
    let [n, idx2] = findLargestInWindow(
      bank.slice(idx, bank.length - windowLength + i + 1)
    );
    res = res * 10 + n;
    idx += idx2 + 1;
  }
  return res;
}

function findLargestInWindow(window: number[]): [number, number] {
  let max = 0;
  let idx = 0;
  for (let i = 0; i < window.length; i++) {
    if (window[i] > max) {
      max = window[i];
      idx = i;
    }
  }
  return [max, idx];
}

const banks = parseInput(puzzleInput);
console.log(part1(banks));
console.log(part2(banks));
