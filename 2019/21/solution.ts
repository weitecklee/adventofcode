import * as fs from "fs";
import * as path from "path";
import intcodeGenerator, { IntcodeGenerator } from "../intcode/intcode";

const puzzleInput = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split(",")
  .map(Number);

const springdroid = intcodeGenerator(puzzleInput);

function displayMessage(
  droid: IntcodeGenerator,
  retVal: number = -9999,
  display: boolean = false
) {
  let message: string[] = [];
  while (true) {
    const ret = droid.next();
    if (ret.value === retVal) break;
    message.push(String.fromCharCode(ret.value));
  }
  if (display) console.log(message.join(""));
}

function inputCommand(droid: IntcodeGenerator, command: string): number {
  for (let i = 0; i < command.length; i++) {
    droid.next(command.charCodeAt(i));
  }
  return droid.next(10).value;
}

function inputCommands(droid: IntcodeGenerator, commands: string[]): number {
  let res = -1;
  for (const cmd of commands) {
    res = inputCommand(droid, cmd);
  }
  return res;
}

displayMessage(springdroid);

// If any of A/B/C is false, J is set to true. (There is a hole to jump over.)
// T is set to D. (If true, D is ground and safe to jump to.)
// if both T and J are true, J is true. (There is a hole to jump over and ground to land on.)

const commands = [
  "NOT A J",
  "NOT B T",
  "OR T J",
  "NOT C T",
  "OR T J",
  "NOT D T",
  "NOT T T",
  "AND T J",
  "WALK",
];

inputCommands(springdroid, commands);

displayMessage(springdroid, 10);
springdroid.next();
console.log(springdroid.next().value);

const springdroid2 = intcodeGenerator(puzzleInput);

const commands2 = ["RUN"];

let line: string[] = [];

displayMessage(springdroid2);
line.push(String.fromCharCode(inputCommands(springdroid2, commands2)));

displayMessage(springdroid2, 10);
while (true) {
  const ret = springdroid2.next();
  if (ret.done) break;
  if (ret.value === 10) {
    console.log(line.join(""));
    line = [];
  } else {
    line.push(String.fromCharCode(ret.value));
  }
}
