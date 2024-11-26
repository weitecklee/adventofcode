const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n");

class Beam {
  constructor(pos, dir, energyHistory) {
    this.pos = pos;
    this.dir = dir;
    this.energyHistory = energyHistory;
  }

  advance() {
    // returns [success, newBeam]
    // if !success, discard beam (hit wall or already visited)
    // if success, keep beam and if newBeam is not null, add it to beams
    this.pos[0] += this.dir[0];
    this.pos[1] += this.dir[1];
    if (
      this.pos[0] < 0 ||
      this.pos[0] >= input[0].length ||
      this.pos[1] < 0 ||
      this.pos[1] >= input.length
    ) {
      return [false, null];
    }
    const beamPos = this.pos.join(",");
    const beamDir = this.dir.join(",");
    if (this.energyHistory.has(beamPos)) {
      const history = this.energyHistory.get(beamPos);
      if (history.has(beamDir)) {
        return [false, null];
      } else {
        history.add(beamDir);
      }
    } else {
      this.energyHistory.set(beamPos, new Set([beamDir]));
    }
    const c = input[this.pos[1]][this.pos[0]];
    switch (c) {
      case ".":
        return [true, null];
      case "|":
        if (this.dir[0] === 0) {
          return [true, null];
        } else {
          this.dir = [0, 1];
          return [
            true,
            new Beam(this.pos.slice(), [0, -1], this.energyHistory),
          ];
        }
      case "-":
        if (this.dir[1] === 0) {
          return [true, null];
        } else {
          this.dir = [1, 0];
          return [
            true,
            new Beam(this.pos.slice(), [-1, 0], this.energyHistory),
          ];
        }
      case "/":
        this.dir = [-this.dir[1], -this.dir[0]];
        return [true, null];
      case "\\":
        this.dir = [this.dir[1], this.dir[0]];
        return [true, null];
      default:
        throw new Error("Invalid character: ", c);
    }
  }
}

function calculateEnergizedTiles(pos, dir) {
  const energyHistory = new Map();
  let beams = [new Beam(pos, dir, energyHistory)];
  while (beams.length) {
    const newBeams = [];
    for (const beam of beams) {
      const [success, newBeam] = beam.advance();
      if (success) {
        newBeams.push(beam);
        if (newBeam) {
          newBeams.push(newBeam);
        }
      }
    }
    beams = newBeams;
  }

  return energyHistory.size;
}

const part1 = calculateEnergizedTiles([-1, 0], [1, 0]);
console.log(part1);

let part2 = 0;
const startingAlignments = [];

for (let c = 0; c < input[0].length; c++) {
  startingAlignments.push([
    [c, -1],
    [0, 1],
  ]);
  startingAlignments.push([
    [c, input.length],
    [0, -1],
  ]);
}

for (let r = 0; r < input.length; r++) {
  startingAlignments.push([
    [-1, r],
    [1, 0],
  ]);
  startingAlignments.push([
    [input[0].length, r],
    [-1, 0],
  ]);
}

for (const [pos, dir] of startingAlignments) {
  part2 = Math.max(calculateEnergizedTiles(pos, dir), part2);
}

console.log(part2);
