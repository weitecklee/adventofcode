import * as fs from "fs";
import * as path from "path";
import intcodeGenerator, { IntcodeGenerator } from "../intcode";

const puzzleInput = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split(",")
  .map(Number);

const visited: Set<string> = new Set(["0,0"]);
const queue: [IntcodeGenerator, number[], number[]][] = [
  [intcodeGenerator(puzzleInput), [], [0, 0]],
];

const cloneRC = (commands: number[]) => {
  const newRC = intcodeGenerator(puzzleInput);
  for (const n of commands) {
    newRC.next();
    newRC.next(n);
  }
  return newRC;
};

const directions = [
  [0, 1],
  [0, -1],
  [1, 0],
  [-1, 0],
];

const reverse = (d: number) => {
  if (d < 3) return 3 - d;
  return 7 - d;
};

const areaMap: Map<string, number> = new Map([["0,0", 1]]);
let oxygenLocation: number[] = [0, 0];

let i = 0;
while (i < queue.length) {
  const [remoteControl, commands, pos] = queue[i];
  i++;

  const candidates: [number, number[]][] = [];
  for (let j = 1; j < 5; j++) {
    const [dx, dy] = directions[j - 1];
    const pos2 = [pos[0] + dx, pos[1] + dy];
    if (visited.has(pos2.join(","))) continue;
    visited.add(pos2.join(","));
    remoteControl.next();
    const ret = remoteControl.next(j);
    areaMap.set(pos2.join(","), ret.value);
    if (ret.value !== 0) {
      candidates.push([j, pos2]);
      if (ret.value === 2) {
        console.log(commands.length + 1);
        oxygenLocation = pos2;
      }
      remoteControl.next();
      remoteControl.next(reverse(j));
    }
  }

  if (candidates.length === 0) continue;
  for (let j = 0; j < candidates.length - 1; j++) {
    const [command, pos2] = candidates[j];
    const remoteControl2 = cloneRC(commands);
    remoteControl2.next();
    remoteControl2.next(command);
    queue.push([remoteControl2, commands.concat(command), pos2]);
  }
  const [command, pos2] = candidates[candidates.length - 1];
  remoteControl.next();
  remoteControl.next(command);
  queue.push([remoteControl, commands.concat(command), pos2]);
}

const visited2: Set<string> = new Set([oxygenLocation.join(",")]);
const queue2: [number[], number][] = [[oxygenLocation, 0]];

let part2 = 0;
while (queue2.length) {
  const [pos, steps] = queue2.pop()!;
  part2 = Math.max(part2, steps);
  for (const [dx, dy] of directions) {
    const pos2 = [pos[0] + dx, pos[1] + dy];
    const coord2 = pos2.join(",");
    if (areaMap.get(coord2)! === 0) continue;
    if (visited2.has(coord2)) continue;
    visited2.add(coord2);
    queue2.push([pos2, steps + 1]);
  }
}

console.log(part2);

// const coords = Array.from(areaMap.keys()).map((a) => a.split(",").map(Number));
// const xMin = Math.min(...coords.map((a) => a[0]));
// const xMax = Math.max(...coords.map((a) => a[0]));
// const yMin = Math.min(...coords.map((a) => a[1]));
// const yMax = Math.max(...coords.map((a) => a[1]));
// console.log(xMin, xMax, yMin, yMax);
// const xRange = xMax - xMin + 1;
// const yRange = yMax - yMin + 1;
// const area = Array(yRange)
//   .fill(0)
//   .map(() => Array(xRange).fill("#"));
// const tiles = new Map([
//   [0, "#"],
//   [1, " "],
//   [2, "@"],
//   [3, "R"],
// ]);
// for (const [k, v] of areaMap) {
//   const [x, y] = k.split(",").map(Number);
//   area[y - yMin][x - xMin] = tiles.get(v)!;
// }
// area.forEach((a) => console.log(a.join("")));
