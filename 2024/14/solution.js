const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n");

class Robot {
  constructor(line) {
    const numbers = line.match(/-?\d+/g).map(Number);
    this.pos = [numbers[0], numbers[1]];
    this.vel = [numbers[2], numbers[3]];
  }
  simulate(time, spaceWidth, spaceHeight) {
    for (let i = 0; i < time; i++) {
      this.step(spaceWidth, spaceHeight);
    }
    return this.pos;
  }
  step(spaceWidth, spaceHeight) {
    this.pos[0] += this.vel[0];
    this.pos[1] += this.vel[1];
    if (this.pos[0] < 0) this.pos[0] += spaceWidth;
    else this.pos[0] %= spaceWidth;
    if (this.pos[1] < 0) this.pos[1] += spaceHeight;
    else this.pos[1] %= spaceHeight;
    return this.pos;
  }
}

const spaceWidth = 101;
const spaceHeight = 103;
const halfWidth = (spaceWidth - 1) / 2;
const halfHeight = (spaceHeight - 1) / 2;
const robots = input.map((line) => new Robot(line));
const quads = [0, 0, 0, 0];
robots.forEach((robot) => {
  const [x, y] = robot.simulate(100, spaceWidth, spaceHeight);
  if (x < halfWidth) {
    if (y < halfHeight) quads[0]++;
    else if (y > halfHeight) quads[1]++;
  } else if (x > halfWidth) {
    if (y < halfHeight) quads[2]++;
    else if (y > halfHeight) quads[3]++;
  }
});
console.log(quads.reduce((a, b) => a * b));

let i = 100;
while (true) {
  i++;
  const coords = new Set();
  robots.forEach((robot) => {
    const [x, y] = robot.step(spaceWidth, spaceHeight);
    coords.add(`${x},${y}`);
  });
  if (coords.size === 500) {
    console.log(i);
    console.log("\n");
    const canvas = Array(spaceHeight)
      .fill()
      .map(() => Array(spaceWidth).fill(" "));
    coords.forEach((coord) => {
      const [x, y] = coord.split(",").map(Number);
      canvas[y][x] = "#";
    });
    console.log(canvas.map((row) => row.join("")).join("\n"));
    break;
  }
}
