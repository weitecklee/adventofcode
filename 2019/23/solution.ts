import * as fs from "fs";
import * as path from "path";
import intcodeGenerator, { IntcodeGenerator } from "../intcode";

const puzzleInput = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split(",")
  .map(Number);

const computers: Map<number, IntcodeGenerator> = new Map();

let outbox: number[][] = [];

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
let ret: IteratorResult<number, number>;
let natPacket = [0, 0];
let part1 = 0;
const ySet = new Set();
while (true) {
  while (i < outbox.length) {
    const [dst, x, y] = outbox[i];
    i++;
    if (dst === 255) {
      natPacket = [x, y];
      if (!part1) part1 = y;
      continue;
    }
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
  outbox = [];
  i = 0;
  for (let j = 0; j < 50; j++) {
    const computer = computers.get(j)!;
    ret = computer.next(-1);
    while (ret.value != -9999) {
      const dst = ret.value;
      const x = computer.next().value;
      const y = computer.next().value;
      outbox.push([dst, x, y]);
      ret = computer.next();
    }
  }
  if (outbox.length === 0) {
    if (ySet.has(natPacket[1])) {
      console.log(part1);
      console.log(natPacket[1]);
      break;
    }
    ySet.add(natPacket[1]);
    outbox = [[0, ...natPacket]];
  }
}
