import * as fs from "fs";
import * as path from "path";
import intcodeGenerator from "../intcode/intcode";

const puzzleInput = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split(",")
  .map(Number);

function boostProgram(input: number): number {
  const program = intcodeGenerator(puzzleInput);
  program.next();
  return program.next(input).value;
}

console.log(boostProgram(1));
console.log(boostProgram(2));
