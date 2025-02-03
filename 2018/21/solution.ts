import * as fs from "fs";
import * as path from "path";

const puzzleInput: [string, ...number[]][] = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split(" "))
  .map((a) => {
    let [opcode, ...rest] = a;
    return [opcode, ...rest.map(Number)];
  });

// Only one line involves register 0, line 28 (eqrr 3 0 1)
// When this line evaluates to true (register 3 === register 0),
// the program ends soon after.
// General solution is to run the program repeatedly and check the value
// of register[3] whenever that line is run. The first value seen is the
// answer to part 1. The last unique value seen is answer to part 2.
// It takes just over a minute to complete...

let ip = puzzleInput.shift()![1];
const register3Set: Set<number> = new Set();

const register = Array(6).fill(0);

while (register[ip] >= 0 && register[ip] < puzzleInput.length) {
  if (register[ip] === 28) {
    if (register3Set.has(register[3])) break;
    register3Set.add(register[3]);
  }
  const [opcode, a, b, c] = puzzleInput[register[ip]];
  switch (opcode) {
    case "addr":
      register[c] = register[a] + register[b];
      break;
    case "addi":
      register[c] = register[a] + b;
      break;
    case "mulr":
      register[c] = register[a] * register[b];
      break;
    case "muli":
      register[c] = register[a] * b;
      break;
    case "banr":
      register[c] = register[a] & register[b];
      break;
    case "bani":
      register[c] = register[a] & b;
      break;
    case "borr":
      register[c] = register[a] | register[b];
      break;
    case "bori":
      register[c] = register[a] | b;
      break;
    case "setr":
      register[c] = register[a];
      break;
    case "seti":
      register[c] = a;
      break;
    case "gtir":
      register[c] = a > register[b] ? 1 : 0;
      break;
    case "gtri":
      register[c] = register[a] > b ? 1 : 0;
      break;
    case "gtrr":
      register[c] = register[a] > register[b] ? 1 : 0;
      break;
    case "eqir":
      register[c] = a === register[b] ? 1 : 0;
      break;
    case "eqri":
      register[c] = register[a] === b ? 1 : 0;
      break;
    case "eqrr":
      register[c] = register[a] === register[b] ? 1 : 0;
      break;
    default:
      throw new Error("Unknown opcode: " + opcode);
  }
  register[ip]++;
}

const values = Array.from(register3Set);
console.log(values[0]);
console.log(values[values.length - 1]);
