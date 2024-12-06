const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split(""));

class Guard {
  constructor() {
    this.pos = [-1, -1];
    this.dir = [0, -1];
  }

  setPos(pos) {
    this.pos = pos.slice();
  }

  setDir(dir) {
    this.dir = dir.slice();
  }

  turn() {
    this.dir = [-this.dir[1], this.dir[0]];
  }

  patrol(patrolArea = input) {
    const patrolPath = new Set();
    const patrolPathWithDirection = new Set();
    while (
      this.pos[1] >= 0 &&
      this.pos[1] < patrolArea.length &&
      this.pos[0] >= 0 &&
      this.pos[0] < patrolArea[0].length
    ) {
      patrolPath.add(`${this.pos[1]},${this.pos[0]}`);
      while (
        this.pos[1] + this.dir[1] >= 0 &&
        this.pos[1] + this.dir[1] < patrolArea.length &&
        this.pos[0] + this.dir[0] >= 0 &&
        this.pos[0] + this.dir[0] < patrolArea[0].length &&
        patrolArea[this.pos[1] + this.dir[1]][this.pos[0] + this.dir[0]] === "#"
      ) {
        this.turn();
      }
      if (
        patrolPathWithDirection.has(
          `${this.pos[1]},${this.pos[0]},${this.dir[1]},${this.dir[0]}`
        )
      )
        return null;
      patrolPathWithDirection.add(
        `${this.pos[1]},${this.pos[0]},${this.dir[1]},${this.dir[0]}`
      );
      this.pos[1] += this.dir[1];
      this.pos[0] += this.dir[0];
    }
    return patrolPath;
  }
}

const guard = new Guard();

for (let r = 0; r < input.length; r++) {
  for (let c = 0; c < input[r].length; c++) {
    if (input[r][c] === "^") {
      guard.setPos([c, r]);
      break;
    }
  }
}
const origPos = guard.pos.slice();

const patrolPath = guard.patrol();
console.log(patrolPath.size);

console.time("part2");
let part2 = 0;
for (const coord of patrolPath) {
  guard.setPos(origPos);
  guard.setDir([0, -1]);
  const [r, c] = coord.split(",").map(Number);
  input[r][c] = "#";
  const tmpPath = guard.patrol(input);
  if (!tmpPath) {
    part2++;
  }
  input[r][c] = ".";
}
console.log(part2);
console.timeEnd("part2");

// Brute force solution : ~3.7s
