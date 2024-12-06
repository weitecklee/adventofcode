const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split(""));

class Guard {
  constructor() {
    this.dir = [0, -1];
  }

  setPos(pos) {
    this.pos = pos;
  }

  turn() {
    this.dir = [-this.dir[1], this.dir[0]];
  }

  patrol(patrolArea = input) {
    this.dir = [0, -1];
    let [r, c] = this.pos;
    const patrolPath = new Set();
    const patrolPathWithDirection = new Set();
    while (
      r >= 0 &&
      r < patrolArea.length &&
      c >= 0 &&
      c < patrolArea[r].length
    ) {
      patrolPath.add(`${r},${c}`);
      while (
        r + this.dir[1] >= 0 &&
        r + this.dir[1] < patrolArea.length &&
        c + this.dir[0] >= 0 &&
        c + this.dir[0] < patrolArea[r].length &&
        patrolArea[r + this.dir[1]][c + this.dir[0]] === "#"
      ) {
        this.turn();
      }
      if (patrolPathWithDirection.has(`${r},${c},${this.dir.join(",")}`))
        return null;
      patrolPathWithDirection.add(`${r},${c},${this.dir.join(",")}`);
      r += this.dir[1];
      c += this.dir[0];
    }
    return patrolPath;
  }
}

const guard = new Guard();
const obstructionRows = new Map();
const obstructionCols = new Map();

for (let r = 0; r < input.length; r++) {
  for (let c = 0; c < input[r].length; c++) {
    if (input[r][c] === "^") {
      guard.setPos([r, c]);
      break;
    }
  }
}

const patrolPath = guard.patrol();
console.log(patrolPath.size);

let part2 = 0;
for (const coord of patrolPath) {
  const [r, c] = coord.split(",").map(Number);
  input[r][c] = "#";
  const tmpPath = guard.patrol(input);
  if (!tmpPath) {
    part2++;
  }
  input[r][c] = ".";
}
console.log(part2);
