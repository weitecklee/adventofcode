import * as fs from "fs";
import * as path from "path";

const puzzleInput = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.match(/-?\d+/g)!.map((b) => Number(b)));

class Nanobot {
  pos: number[];
  range: number;
  numInRange: number;
  constructor(nums: number[]) {
    this.pos = nums.slice(0, 3);
    this.range = nums[3];
    this.numInRange = 1;
  }

  distanceFrom(bot: Nanobot) {
    return this.pos.reduce((a, b, i) => a + Math.abs(b - bot.pos[i]), 0);
  }
}

const nanobots = puzzleInput.map((a) => new Nanobot(a));

for (let i = 0; i < nanobots.length; i++) {
  const bot1 = nanobots[i];
  for (let j = i + 1; j < nanobots.length; j++) {
    const bot2 = nanobots[j];
    const dist = bot1.distanceFrom(bot2);
    if (dist <= bot1.range) bot1.numInRange++;
    if (dist <= bot2.range) bot2.numInRange++;
  }
}

let maxRange = 0;
let part1 = 0;
for (const bot of nanobots) {
  if (bot.range > maxRange) {
    maxRange = bot.range;
    part1 = bot.numInRange;
  }
}

console.log(part1);
