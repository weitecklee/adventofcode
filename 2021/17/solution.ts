import * as fs from "fs";
import * as path from "path";

const input = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8");
const targetArea = input.match(/-?\d+/g)!.map((match) => Number(match));

class Probe {
  xPos: number;
  yPos: number;
  xVel: number;
  yVel: number;
  yMax: number;
  constructor(xVel: number, yVel: number) {
    this.xPos = 0;
    this.yPos = 0;
    this.xVel = xVel;
    this.yVel = yVel;
    this.yMax = 0;
  }

  fly(targetArea: number[]): number {
    while (this.yPos >= targetArea[2]) {
      if (
        this.xPos >= targetArea[0] &&
        this.xPos <= targetArea[1] &&
        this.yPos <= targetArea[3]
      ) {
        return this.yMax;
      }
      this.xPos += this.xVel;
      this.yPos += this.yVel;
      this.yMax = Math.max(this.yMax, this.yPos);
      this.xVel -= Math.sign(this.xVel);
      this.yVel--;
    }
    return -1;
  }
}

let part1 = 0;
let part2 = 0;
for (let xVel = 0; xVel <= targetArea[1]; xVel++) {
  for (let yVel = targetArea[2]; yVel < 1000; yVel++) {
    const probe = new Probe(xVel, yVel);
    const yMax = probe.fly(targetArea);
    if (yMax >= 0) {
      part1 = Math.max(part1, yMax);
      part2++;
    }
  }
}

console.log(part1);
console.log(part2);
