import * as fs from "fs";
import * as path from "path";
import intcodeGenerator from "../intcode";

const puzzleInput = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split(",")
  .map(Number);

const springdroid = intcodeGenerator(puzzleInput);

function displayMessage(retVal: number = -9999, display: boolean = false) {
  let message: string[] = [];
  while (true) {
    const ret = springdroid.next();
    if (ret.value === retVal) break;
    message.push(String.fromCharCode(ret.value));
  }
  if (display) console.log(message.join(""));
}

function inputCommand(command: string): number {
  for (let i = 0; i < command.length; i++) {
    springdroid.next(command.charCodeAt(i));
  }
  return springdroid.next(10).value;
}

function inputCommands(commands: string[]): number {
  let res = -1;
  for (const cmd of commands) {
    res = inputCommand(cmd);
  }
  return res;
}

displayMessage();

// If any of A/B/C is false, J is set to true. (There is a hole to jump over.)
// T is set to D. (D is ground and safe to jump to.)
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

let line: string[] = [];

line.push(String.fromCharCode(inputCommands(commands)));

displayMessage(10);
springdroid.next();
console.log(springdroid.next().value);
