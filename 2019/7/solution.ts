import * as fs from "fs";
import * as path from "path";

const puzzleInput = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split(",")
  .map(Number);

type IntcodeGenerator = Generator<number, number, number>;

function* intcodeGenerator(prog: number[], input: number[]): IntcodeGenerator {
  function getParams(
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

  const program = prog.slice();
  let i = 0;
  let j = 0;
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
        params = getParams(program, parameterModes, 2, i).map(
          (a) => program[a]
        );
        break;
      case 7:
      case 8:
        params = getParams(program, parameterModes, 3, i);
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
        program[program[++i]] = input[j++];
        break;
      case 4:
        //output
        yield program[program[++i]];
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
  return -1;
}

class AmplifierControllerSoftware {
  program: number[];
  constructor(program: number[]) {
    this.program = program;
  }

  run(phaseSeq: number[]): number {
    const inputA = [phaseSeq[0], 0];
    const ampA = intcodeGenerator(this.program, inputA);
    const inputB = [phaseSeq[1]];
    const ampB = intcodeGenerator(this.program, inputB);
    const inputC = [phaseSeq[2]];
    const ampC = intcodeGenerator(this.program, inputC);
    const inputD = [phaseSeq[3]];
    const ampD = intcodeGenerator(this.program, inputD);
    const inputE = [phaseSeq[4]];
    const ampE = intcodeGenerator(this.program, inputE);
    let output = 0;
    while (true) {
      inputB.push(ampA.next().value);
      inputC.push(ampB.next().value);
      inputD.push(ampC.next().value);
      inputE.push(ampD.next().value);
      const ret = ampE.next();
      if (ret.done) {
        break;
      }
      output = ret.value;
      inputA.push(output);
    }
    return output;
  }
}

const acs = new AmplifierControllerSoftware(puzzleInput);

let part1 = 0;
const used: Set<number> = new Set();
for (let a = 0; a < 5; a++) {
  used.add(a);
  for (let b = 0; b < 5; b++) {
    if (used.has(b)) continue;
    used.add(b);
    for (let c = 0; c < 5; c++) {
      if (used.has(c)) continue;
      used.add(c);
      for (let d = 0; d < 5; d++) {
        if (used.has(d)) continue;
        used.add(d);
        for (let e = 0; e < 5; e++) {
          if (used.has(e)) continue;
          const output = acs.run([a, b, c, d, e]);
          if (output > part1) part1 = output;
        }
        used.delete(d);
      }
      used.delete(c);
    }
    used.delete(b);
  }
  used.delete(a);
}
console.log(part1);

let part2 = 0;
for (let a = 5; a < 10; a++) {
  used.add(a);
  for (let b = 5; b < 10; b++) {
    if (used.has(b)) continue;
    used.add(b);
    for (let c = 5; c < 10; c++) {
      if (used.has(c)) continue;
      used.add(c);
      for (let d = 5; d < 10; d++) {
        if (used.has(d)) continue;
        used.add(d);
        for (let e = 5; e < 10; e++) {
          if (used.has(e)) continue;
          const output = acs.run([a, b, c, d, e]);
          if (output > part2) part2 = output;
        }
        used.delete(d);
      }
      used.delete(c);
    }
    used.delete(b);
  }
  used.delete(a);
}
console.log(part2);
