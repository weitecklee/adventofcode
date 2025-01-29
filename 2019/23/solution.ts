import * as fs from "fs";
import * as path from "path";
import intcodeGenerator, { IntcodeGenerator } from "../intcode";

const puzzleInput = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split(",")
  .map(Number);

const computers: Map<number, IntcodeGenerator> = new Map();

const outbox: number[][] = [];

for (let i = 0; i < 50; i++) {
  const computer = intcodeGenerator(puzzleInput);
  computer.next();
  computer.next(i);
  let ret = computer.next(-1);
  while (ret.value != -9999) {
    const dst = ret.value;
    const x = computer.next().value;
    const y = computer.next().value;
    outbox.push([dst, x, y]);
    ret = computer.next();
  }
  computers.set(i, computer);
}

let i = 0;
let ret;
while (i < outbox.length) {
  const [dst, x, y] = outbox[i];
  if (dst === 255) {
    console.log(y);
    break;
  }
  i++;
  const computer = computers.get(dst)!;
  computer.next(x);
  ret = computer.next(y);
  while (ret.value != -9999) {
    const dst = ret.value;
    const x = computer.next().value;
    const y = computer.next().value;
    outbox.push([dst, x, y]);
    ret = computer.next();
  }
}
