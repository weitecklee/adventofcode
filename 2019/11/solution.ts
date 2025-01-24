import * as fs from "fs";
import * as path from "path";
import intcodeGenerator from "../intcode";

const puzzleInput = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split(",")
  .map(Number);

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
