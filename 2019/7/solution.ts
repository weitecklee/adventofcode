import * as fs from "fs";
import * as path from "path";
import intcodeGenerator from "../intcode";

const puzzleInput = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split(",")
  .map(Number);

class AmplifierControllerSoftware {
  program: number[];
  constructor(program: number[]) {
    this.program = program;
  }

  run(phaseSeq: number[]): number {
    const ampA = intcodeGenerator(this.program);
    const ampB = intcodeGenerator(this.program);
    const ampC = intcodeGenerator(this.program);
    const ampD = intcodeGenerator(this.program);
    const ampE = intcodeGenerator(this.program);
    let outputA = 0;
    let outputB = 0;
    let outputC = 0;
    let outputD = 0;
    let outputE = 0;
    ampA.next();
    ampA.next(phaseSeq[0]);
    ampB.next();
    ampB.next(phaseSeq[1]);
    ampC.next();
    ampC.next(phaseSeq[2]);
    ampD.next();
    ampD.next(phaseSeq[3]);
    ampE.next();
    ampE.next(phaseSeq[4]);
    while (true) {
      outputA = ampA.next(outputE).value;
      outputB = ampB.next(outputA).value;
      outputC = ampC.next(outputB).value;
      outputD = ampD.next(outputC).value;
      const ret = ampE.next(outputD);
      if (ret.done) {
        break;
      }
      outputE = ret.value;
      ampA.next();
      ampB.next();
      ampC.next();
      ampD.next();
      ampE.next();
    }
    return outputE;
  }
}

function permutations(arr: number[]): number[][] {
  if (arr.length <= 1) return [arr];
  const res: number[][] = [];
  for (let i = 0; i < arr.length; i++) {
    const others = arr.slice(0, i).concat(arr.slice(i + 1));
    for (const n of permutations(others)) {
      res.push([arr[i], ...n]);
    }
  }
  return res;
}

const acs = new AmplifierControllerSoftware(puzzleInput);

console.log(Math.max(...permutations([0, 1, 2, 3, 4]).map((a) => acs.run(a))));
console.log(Math.max(...permutations([5, 6, 7, 8, 9]).map((a) => acs.run(a))));
