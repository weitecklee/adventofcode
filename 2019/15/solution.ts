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

let i = 0;
let found = false;
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
    if (ret.value === 0) {
      continue;
    }
    if (ret.value === 1) {
      queue.push([commands.concat(j), pos2]);
    } else if (ret.value === 2) {
      console.log(commands.length + 1);
      found = true;
      break;
    }
  }
  if (found) break;
}
