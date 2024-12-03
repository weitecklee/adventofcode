const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split(" @ ").map((b) => b.split(", ").map(Number)));

class Hailstone {
  constructor(pos, vel) {
    this.pos = pos;
    this.vel = vel;
    this.m = vel[1] / vel[0];
    this.b = pos[1] - this.m * pos[0];
  }
  findIntersection(hailstone) {
    const x = (hailstone.b - this.b) / (this.m - hailstone.m);
    const y = this.m * x + this.b;
    const t1 = (x - this.pos[0]) / this.vel[0];
    const t2 = (x - hailstone.pos[0]) / hailstone.vel[0];
    return [x, y, t1, t2];
  }
}

const hailstones = input.map(([pos, vel]) => new Hailstone(pos, vel));
let part1 = 0;
for (let i = 0; i < hailstones.length; i++) {
  for (let j = i + 1; j < hailstones.length; j++) {
    const intersection = hailstones[i].findIntersection(hailstones[j]);
    if (
      intersection[2] > 0 &&
      intersection[3] > 0 &&
      intersection[0] >= 200000000000000 &&
      intersection[0] <= 400000000000000 &&
      intersection[1] >= 200000000000000 &&
      intersection[1] <= 400000000000000
    )
      part1++;
  }
}

console.log(part1);
