import * as fs from "fs";
import * as path from "path";

const puzzleInput = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split(",")
  .map(Number);

class Intcode {
  constructor() {}

  runProgram(prog: number[], input: number): number {
    const program = prog.slice();
    let i = 0;
    let output = 0;
    while (i >= 0 && i < program.length) {
      const opcode = program[i] % 100;
      const parameterModes = Math.trunc(program[i] / 100)
        .toString()
        .split("")
        .reverse()
        .map(Number);

      let params: number[] = [];
      switch (opcode) {
        case 1:
        case 2:
        case 5:
        case 6:
          params = this.#getParams(program, parameterModes, 2, i).map(
            (a) => program[a]
          );
          break;
        case 7:
        case 8:
          params = this.#getParams(program, parameterModes, 3, i);
          break;
      }

      switch (opcode) {
        case 1:
          // add
          program[program[i + 3]] = params.reduce((a, b) => a + b, 0);
          i += 3;
          break;
        case 2:
          // multiply
          program[program[i + 3]] = params.reduce((a, b) => a * b, 1);
          i += 3;
          break;
        case 3:
          // save input
          program[program[++i]] = input;
          break;
        case 4:
          //output
          output = program[program[++i]];
          break;
        case 5:
          // jump if true
          if (params[0] !== 0) {
            i = params[1];
            continue;
          } else {
            i += 2;
          }
          break;
        case 6:
          // jump if false
          if (params[0] === 0) {
            i = params[1];
            continue;
          } else {
            i += 2;
          }
          break;
        case 7:
          // less than
          if (program[params[0]] < program[params[1]]) {
            program[params[2]] = 1;
          } else {
            program[params[2]] = 0;
          }
          i += 3;
          break;
        case 8:
          // equal
          if (program[params[0]] === program[params[1]]) {
            program[params[2]] = 1;
          } else {
            program[params[2]] = 0;
          }
          i += 3;
          break;
        case 99:
          // halt
          i = -99;
          break;
        default:
          throw new Error("Unknown opcode: " + program[i]);
      }
      i++;
    }
    return output;
  }

  #getParams(
    program: number[],
    parameterModes: number[],
    nParams: number,
    i: number
  ): number[] {
    const parameters: number[] = [];
    for (let j = 0; j < nParams; j++) {
      if (parameterModes[j] && parameterModes[j] === 1) {
        parameters.push(i + j + 1);
      } else {
        parameters.push(program[i + j + 1]);
      }
    }
    return parameters;
  }
}

const intcode = new Intcode();
console.log(intcode.runProgram(puzzleInput, 1));
console.log(intcode.runProgram(puzzleInput, 5));
