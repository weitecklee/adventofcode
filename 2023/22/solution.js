const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split("~").map((b) => b.split(",").map(Number)));

input.forEach((brick) => brick.sort((a, b) => a[2] - b[2]));
input.sort((a, b) => a[0][2] - b[0][2]);

function Brick(i) {
  this.index = i;
  this.atop = new Set();
  this.supporting = new Set();
  this.z = 0;
}

const coordMap = new Map();
const unsafeBricks = new Set();
const bricks = [];

let part1 = 0;
for (let i = 0; i < input.length; i++) {
  const currBrick = new Brick(i);
  const [endA, endB] = input[i];
  const [x0, y0, z0] = endA;
  const [x1, y1, z1] = endB;
  let z = 1;
  let count = 0;
  for (let x = x0; x <= x1; x++) {
    for (let y = y0; y <= y1; y++) {
      const coord = `${x},${y}`;
      if (coordMap.has(coord)) {
        const [h] = coordMap.get(coord);
        z = Math.max(h, z);
      }
    }
  }
  for (let x = x0; x <= x1; x++) {
    for (let y = y0; y <= y1; y++) {
      const coord = `${x},${y}`;
      if (coordMap.has(coord)) {
        const [h, brick] = coordMap.get(coord);
        if (h === z) {
          currBrick.atop.add(brick);
          bricks[brick].supporting.add(i);
        }
      }
      coordMap.set(coord, [z + z1 - z0 + 1, i]);
    }
  }
  if (currBrick.atop.size === 1) {
    unsafeBricks.add(currBrick.atop.values().next().value);
  }
  currBrick.z = z + z1 - z0 + 1;
  bricks.push(currBrick);
}

console.log(input.length - unsafeBricks.size);

let part2 = 0;

for (const brick of unsafeBricks) {
  const fallenBricks = new Set([brick]);
  let oldSize = 0;
  while (fallenBricks.size != oldSize) {
    oldSize = fallenBricks.size;
    for (let i = 0; i < bricks.length; i++) {
      if (bricks[i].atop.size && fallenBricks.isSupersetOf(bricks[i].atop)) {
        fallenBricks.add(i);
      }
    }
  }
  part2 += fallenBricks.size - 1;
}
console.log(part2);
