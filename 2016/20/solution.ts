import * as fs from "fs";
import * as path from "path";

const puzzleInput = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split("-").map(Number)) as [number, number][];

const MAXIMUMIP = 4294967295;

class RangeCollection {
  ranges: [number, number][];
  constructor(ranges?: [number, number][]) {
    if (ranges && ranges.length) {
      this.ranges = ranges.sort((a, b) => a[0] - b[0]);
      this.addRange(ranges[0]);
    } else {
      this.ranges = [];
    }
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
    if (this.ranges.length === 0) return 0;
    return this.ranges[0][1] + 1;
  }

  get nNonBlockedIPs(): number {
    if (this.ranges.length === 0) return MAXIMUMIP + 1;
    let res = 0;
    let curr = -1;
    for (const range of this.ranges) {
      res += range[0] - curr - 1;
      curr = range[1];
    }
    return res + MAXIMUMIP - this.ranges[this.ranges.length - 1][1];
  }
}

const ranges = new RangeCollection(puzzleInput);

console.log(ranges.lowestNonBlockedIP);
console.log(ranges.nNonBlockedIPs);
