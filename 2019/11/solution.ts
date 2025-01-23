import * as fs from "fs";
import * as path from "path";

const puzzleInput = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split(",")
  .map(Number);

class Memory<K, V> extends Map<K, V> {
  get(key: K): V {
    return (super.get(key) ?? 0) as V;
  }
}

type IntcodeGenerator = Generator<number, number, number>;

function* intcodeGenerator(prog: number[]): IntcodeGenerator {
  function getParams(
    program: Memory<number, number>,
    parameterModes: number[],
    nParams: number,
    i: number,
    relativeBase: number
  ): number[] {
    const parameters: number[] = [];
    for (let j = 0; j < nParams; j++) {
      if (parameterModes[j]) {
        if (parameterModes[j] === 1) {
          // immediate mode
          parameters.push(i + j + 1);
        } else if (parameterModes[j] === 2) {
          // relative mode
          parameters.push(program.get(i + j + 1) + relativeBase);
        }
      } else {
        // position mode
        parameters.push(program.get(i + j + 1));
      }
    }
    return parameters;
  }

  const program = new Memory(prog.map((a, i) => [i, a]));
  let i = 0; // program index
  let relativeBase = 0;
  while (i >= 0) {
    const opcode = program.get(i) % 100;
    const parameterModes = Math.trunc(program.get(i) / 100)
      .toString()
      .split("")
      .reverse()
      .map(Number);

    let params: number[] = [];
    switch (opcode) {
      case 5:
      case 6:
        params = getParams(program, parameterModes, 2, i, relativeBase).map(
          (a) => program.get(a)
        );
        break;
      case 1:
      case 2:
      case 7:
      case 8:
        params = getParams(program, parameterModes, 3, i, relativeBase);
        break;
      case 3:
      case 4:
      case 9:
        params = getParams(program, parameterModes, 1, i, relativeBase);
        break;
    }

    switch (opcode) {
      case 1:
        // add
        program.set(params[2], program.get(params[0]) + program.get(params[1]));
        i += 3;
        break;
      case 2:
        // multiply
        program.set(params[2], program.get(params[0]) * program.get(params[1]));
        i += 3;
        break;
      case 3:
        // save input
        program.set(params[0], yield 9999);
        i++;
        break;
      case 4:
        //output
        yield program.get(params[0]);
        i++;
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
        if (program.get(params[0]) < program.get(params[1])) {
          program.set(params[2], 1);
        } else {
          program.set(params[2], 0);
        }
        i += 3;
        break;
      case 8:
        // equal
        if (program.get(params[0]) === program.get(params[1])) {
          program.set(params[2], 1);
        } else {
          program.set(params[2], 0);
        }
        i += 3;
        break;
      case 9:
        // adjust relative base
        relativeBase += program.get(params[0]);
        i++;
        break;
      case 99:
        // halt
        i = -99;
        break;
      default:
        throw new Error("Unknown opcode: " + program.get(i));
    }
    i++;
  }
  return -1;
}

function runProgram(startingPanelColor: number): Map<string, number> {
  const robot = intcodeGenerator(puzzleInput);
  const robotPos = [0, 0]; // using RC coordinate system
  const robotDir = [-1, 0];
  const panels: Map<string, number> = new Map([
    [robotPos.join(","), startingPanelColor],
  ]);

  while (true) {
    robot.next();
    const ret = robot.next(panels.get(robotPos.join(",")) ?? 0);
    if (ret.done) break;
    const ret2 = robot.next();
    if (ret2.done) break;
    panels.set(robotPos.join(","), ret.value);
    if (ret2.value === 1) {
      // turn right
      [robotDir[0], robotDir[1]] = [robotDir[1], -robotDir[0]];
    } else {
      // turn left
      [robotDir[0], robotDir[1]] = [-robotDir[1], robotDir[0]];
    }
    robotPos[0] += robotDir[0];
    robotPos[1] += robotDir[1];
  }
  return panels;
}

const part1Panels = runProgram(0);
console.log(part1Panels.size);

const part2Panels = runProgram(1);
const coords = Array.from(part2Panels.keys()).map((a) =>
  a.split(",").map(Number)
);
let rMin = Number.MAX_SAFE_INTEGER;
let rMax = Number.MIN_SAFE_INTEGER;
let cMin = rMin;
let cMax = rMax;

for (const [r, c] of coords) {
  if (r < rMin) rMin = r;
  if (r > rMax) rMax = r;
  if (c < cMin) cMin = c;
  if (c > cMax) cMax = c;
}

const rRange = rMax - rMin + 1;
const cRange = cMax - cMin + 1;

const part2Paint = Array(rRange)
  .fill("")
  .map(() => Array(cRange).fill(" "));

for (const [coordString, paint] of part2Panels) {
  let [r, c] = coordString.split(",").map(Number);
  r -= rMin;
  c -= cMin;
  part2Paint[r][c] = paint === 1 ? "#" : " ";
}

part2Paint.forEach((r) => console.log(r.join("")));
