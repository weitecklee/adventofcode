import * as fs from "fs";
import * as path from "path";
import intcodeGenerator, { IntcodeGenerator } from "../intcode/intcode";

const puzzleInput = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split(",")
  .map(Number);

const computers: Map<number, IntcodeGenerator> = new Map();

let outbox: number[][] = [];

// initialize computers with network addresses
for (let i = 0; i < 50; i++) {
  const computer = intcodeGenerator(puzzleInput);
  computer.next();
  computer.next(i);
  computers.set(i, computer);
}

function fillOutbox(
  outbox: number[][],
  computer: IntcodeGenerator,
  lastRet: IteratorResult<number, number>
) {
  // fills outbox with output from computer until output is -9999
  // (computer is done sending and is now requesting input)
  while (lastRet.value != -9999) {
    const dst = lastRet.value;
    const x = computer.next().value;
    const y = computer.next().value;
    outbox.push([dst, x, y]);
    lastRet = computer.next();
  }
}

let i = 0;
let ret: IteratorResult<number, number>;
let natPacket = [0, 0];
let part1 = 0;
const ySet = new Set();
while (true) {
  // iterate through computers and fill outbox
  for (let j = 0; j < 50; j++) {
    const computer = computers.get(j)!;
    ret = computer.next(-1);
    fillOutbox(outbox, computer, ret);
  }

  // if nothing in outbox, network is idle.
  // NAT sends its last received packet to computer 0.
  if (outbox.length === 0) {
    if (ySet.has(natPacket[1])) {
      // y value has been seen before, puzzle solved!
      console.log(part1);
      console.log(natPacket[1]);
      break;
    }
    ySet.add(natPacket[1]);
    outbox = [[0, ...natPacket]];
  }

  // iterate through outbox (outbox may grow during this process)
  while (i < outbox.length) {
    const [dst, x, y] = outbox[i];
    i++;
    if (dst === 255) {
      // only keep track of last packet sent to 255
      natPacket = [x, y];
      // only keep track of first y value sent to 255
      if (!part1) part1 = y;
      continue;
    }
    const computer = computers.get(dst)!;
    computer.next(x);
    ret = computer.next(y);
    fillOutbox(outbox, computer, ret);
  }

  // reset outbox
  outbox = [];
  i = 0;
}
