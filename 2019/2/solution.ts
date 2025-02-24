import * as fs from "fs";
import * as path from "path";
import intcodeGenerator from "../intcode/intcode";

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split(",")
  .map(Number);

function runProgram(input: number[], a: number, b: number): number {
  const inputCopy = input.slice();
  inputCopy[1] = a;
  inputCopy[2] = b;
  const intcode = intcodeGenerator(inputCopy);
  return intcode.next().value;
}

console.log(runProgram(input, 12, 2));

// Analysis of output for various pairs of nouns and verbs shows that the output
// follows the equation `output = x * noun + y + verb` (for my input).
// Calculate the values of x and y then find an integer solution.
// Also it's only Day 2, we're not gonna be asked to do anything too crazy.
// A search area of 100 x 100 is enough.

let foundAnswer = false;
for (let i = 0; i < 100; i++) {
  for (let j = 0; j < 100; j++) {
    if (runProgram(input, i, j) === 19690720) {
      console.log(100 * i + j);
      foundAnswer = true;
      break;
    }
  }
  if (foundAnswer) break;
}
