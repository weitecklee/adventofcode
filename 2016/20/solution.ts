import * as fs from "fs";
import * as path from "path";

const puzzleInput = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split("-").map(Number)) as [number, number][];

class RangeCollection {
  ranges: [number, number][];
  constructor(ranges: [number, number][]) {
    this.ranges = ranges;
    this.initialize();
  }

  initialize() {
    this.ranges.sort((a, b) => a[0] - b[0]);
    const ranges: [number, number][] = [];
    let curr = this.ranges[0];
    for (let i = 1; i < this.ranges.length; i++) {
      if (this.ranges[i][1] < curr[0] - 1) {
        ranges.push(this.ranges[i]);
      } else if (curr[1] + 1 < this.ranges[i][0]) {
        ranges.push(curr);
        curr = this.ranges[i];
      } else {
        curr[0] = Math.min(curr[0], this.ranges[i][0]);
        curr[1] = Math.max(curr[1], this.ranges[i][1]);
      }
    }
    ranges.push(curr);
    this.ranges = ranges;
  }

  addRange(newRange: [number, number]) {
    const ranges: [number, number][] = [];
    let curr = newRange;
    for (const range of this.ranges) {
      if (range[1] < curr[0] - 1) {
        ranges.push(range);
      } else if (curr[1] + 1 < range[0]) {
        ranges.push(curr);
        curr = range;
      } else {
        curr[0] = Math.min(curr[0], range[0]);
        curr[1] = Math.max(curr[1], range[1]);
      }
    }
    ranges.push(curr);
    this.ranges = ranges;
  }

  get lowestNonBlockedIP(): number {
    return this.ranges[0][1] + 1;
  }

  get nNonBlockedIPs(): number {
    let res = 0;
    let curr = -1;
    for (const range of this.ranges) {
      res += range[0] - curr - 1;
      curr = range[1];
    }
    return res + 4294967295 - this.ranges[this.ranges.length - 1][1];
  }
}

const ranges = new RangeCollection(puzzleInput);

console.log(ranges.lowestNonBlockedIP);
console.log(ranges.nNonBlockedIPs);
