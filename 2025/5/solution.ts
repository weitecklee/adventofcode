import * as fs from "fs";
import * as path from "path";

const puzzleInput = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n\n");

function parseInput(data: string[]): [number[][], number[]] {
  const ranges: number[][] = [];
  for (const line of data[0].split("\n")) {
    ranges.push(line.split("-").map(Number));
  }
  const ingredientIDs: number[] = data[1].split("\n").map(Number);
  return [ranges, ingredientIDs];
}

function condenseRanges(ranges: number[][]): number[][] {
  ranges.sort((a, b) => a[0] - b[0]);
  const res: number[][] = [];
  let tmp: number[] = ranges[0];
  for (const range of ranges) {
    if (range[0] <= tmp[1]) {
      tmp[1] = Math.max(range[1], tmp[1]);
    } else {
      res.push(tmp);
      tmp = range;
    }
  }
  res.push(tmp);
  return res;
}

function solve(ranges: number[][], ingredientIDs: number[]): [number, number] {
  ranges = condenseRanges(ranges);

  let part1 = 0;
  for (const n of ingredientIDs) {
    for (const range of ranges) {
      if (range[0] <= n && n <= range[1]) {
        part1++;
        break;
      }
    }
  }

  const part2 = ranges.reduce((a, b) => a + b[1] - b[0] + 1, 0);
  return [part1, part2];
}

const [ranges, ingredientIDs] = parseInput(puzzleInput);
console.log(solve(ranges, ingredientIDs));
