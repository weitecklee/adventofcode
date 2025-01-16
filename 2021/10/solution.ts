import * as fs from "fs";
import * as path from "path";

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n");

const bracketPairs = new Map([
  ["(", ")"],
  ["[", "]"],
  ["{", "}"],
  ["<", ">"],
]);

const bracketScores = new Map([
  [")", 3],
  ["]", 57],
  ["}", 1197],
  [">", 25137],
]);

const bracketScores2 = new Map([
  [")", 1],
  ["]", 2],
  ["}", 3],
  [">", 4],
]);

let part1 = 0;
const part2: number[] = [];

for (const line of input) {
  const openers: string[] = [];
  let isCorrupted = false;
  for (const c of line) {
    if (bracketPairs.has(c)) openers.push(c);
    else {
      if (bracketPairs.get(openers[openers.length - 1]) === c) {
        openers.pop();
      } else {
        part1 += bracketScores.get(c)!;
        isCorrupted = true;
        break;
      }
    }
  }
  if (!isCorrupted) {
    let score = 0;
    for (let i = openers.length - 1; i >= 0; i--) {
      score *= 5;
      score += bracketScores2.get(bracketPairs.get(openers[i])!)!;
    }
    part2.push(score);
  }
}

console.log(part1);

part2.sort((a, b) => a - b);
console.log(part2[Math.floor(part2.length / 2)]);
