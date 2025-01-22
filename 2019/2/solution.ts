import * as fs from "fs";
import * as path from "path";

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split(",")
  .map(Number);

class Intcode {
  constructor() {}

  runProgram(input: number[], noun: number, verb: number): number {
    const program = input.slice();
    program[1] = noun;
    program[2] = verb;
    let i = 0;
    while (i < program.length) {
      if (program[i] === 1) {
        // add
        const a = program[program[++i]];
        const b = program[program[++i]];
        program[program[++i]] = a + b;
      } else if (program[i] === 2) {
        // multiply
        const a = program[program[++i]];
        const b = program[program[++i]];
        program[program[++i]] = a * b;
      } else if (program[i] === 99) {
        break;
      } else {
        throw new Error("Unknown opcode: " + program[i]);
      }
      i++;
    }
    return program[0];
  }
}

const intcode = new Intcode();
console.log(intcode.runProgram(input, 12, 2));

// Analysis of output for various pairs of nouns and verbs shows that the output
// follows the equation `output = x * noun + y + verb` (for my input).
// Calculate the values of x and y then find an integer solution.
// Also it's only Day 2, we're not gonna be asked to do anything too crazy.
// A search area of 100 x 100 is enough.

let foundAnswer = false;
for (let i = 0; i < 100; i++) {
  for (let j = 0; j < 100; j++) {
    if (intcode.runProgram(input, i, j) === 19690720) {
      console.log(100 * i + j);
      foundAnswer = true;
      break;
    }
  }
  if (foundAnswer) break;
}
