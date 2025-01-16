import * as fs from "fs";
import * as path from "path";

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split("").map(Number));

const rMax = input.length - 1;
const cMax = input[0].length - 1;

let part1 = 0;

function increaseEnergy(input: number[][]): number[][] {
  const flashQueue: number[][] = [];
  for (let r = 0; r < input.length; r++) {
    for (let c = 0; c < input[r].length; c++) {
      input[r][c]++;
      if (input[r][c] === 10) {
        flashQueue.push([r, c]);
      }
    }
  }
  return flashQueue;
}

function flash(input: number[][]): number {
  const flashQueue = increaseEnergy(input);
  for (let j = 0; j < flashQueue.length; j++) {
    const [r, c] = flashQueue[j];
    input[r][c] = 0;
    for (let dr = -1; dr <= 1; dr++) {
      for (let dc = -1; dc <= 1; dc++) {
        if (dr === 0 && dc === 0) continue;
        const r2 = r + dr;
        const c2 = c + dc;
        if (r2 < 0 || r2 > rMax || c2 < 0 || c2 > cMax) continue;
        if (input[r2][c2] === 0) continue;
        input[r2][c2]++;
        if (input[r2][c2] === 10) {
          flashQueue.push([r2, c2]);
        }
      }
    }
  }
  return flashQueue.length;
}

for (let i = 0; i < 100; i++) {
  part1 += flash(input);
}

console.log(part1);

const n = input.length * input[0].length;

let part2 = 100;
while (true) {
  part2++;
  if (flash(input) === n) break;
}

console.log(part2);
