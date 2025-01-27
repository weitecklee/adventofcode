import * as fs from "fs";
import * as path from "path";
import intcodeGenerator from "../intcode";

const puzzleInput = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split(",")
  .map(Number);

const visited: Set<string> = new Set(["0,0"]);
const queue: [number[], number[]][] = [[[], [0, 0]]];

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

const areaMap: Map<string, number> = new Map([["0,0", 1]]);
let oxygenLocation: number[] = [0, 0];

let i = 0;
while (i < queue.length) {
  const [commands, pos] = queue[i];
  i++;
  for (let j = 1; j < 5; j++) {
    const [dx, dy] = directions[j - 1];
    const pos2 = [pos[0] + dx, pos[1] + dy];
    if (visited.has(pos2.join(","))) continue;
    visited.add(pos2.join(","));
    const remoteControl2 = cloneRC(commands);
    remoteControl2.next();
    const ret = remoteControl2.next(j);
    areaMap.set(pos2.join(","), ret.value);
    if (ret.value === 0) {
      continue;
    }
    if (ret.value === 2) {
      console.log(commands.length + 1);
      oxygenLocation = pos2;
    }
    queue.push([commands.concat(j), pos2]);
  }
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
