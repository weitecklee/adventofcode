const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split(""));

const width = input[0].length;
const height = input.length;

class Seat {
  constructor(r, c, occupied, seatMap) {
    this.r = r;
    this.c = c;
    this.occupied = occupied;
    this.seatMap = seatMap;
    this.adjacentOccupied = 0;
  }

  evalStatus() {
    let changed = false;
    if (this.occupied && this.adjacentOccupied >= 4) {
      this.occupied = false;
      changed = true;
    } else if (!this.occupied && this.adjacentOccupied === 0) {
      this.occupied = true;
      changed = true;
    }
    this.adjacentOccupied = 0;
    return changed;
  }

  evalSurroundingSeats() {
    if (this.occupied) {
      let count = 0;
      for (let r = this.r - 1; r <= this.r + 1; r++) {
        for (let c = this.c - 1; c <= this.c + 1; c++) {
          if (r === this.r && c === this.c) continue;
          if (this.seatMap.has(`${r},${c}`)) {
            this.seatMap.get(`${r},${c}`).adjacentOccupied++;
          }
        }
      }
    }
  }

  evalStatus2() {
    let changed = false;
    if (this.occupied && this.adjacentOccupied >= 5) {
      this.occupied = false;
      changed = true;
    } else if (!this.occupied && this.adjacentOccupied === 0) {
      this.occupied = true;
      changed = true;
    }
    this.adjacentOccupied = 0;
    return changed;
  }

  evalSurroundingSeats2() {
    if (this.occupied) {
      let count = 0;
      for (let dx = -1; dx < 2; dx++) {
        for (let dy = -1; dy < 2; dy++) {
          if (dx === 0 && dy === 0) continue;
          let i = 1;
          while (true) {
            const r = this.r + dx * i;
            const c = this.c + dy * i;
            if (r < 0 || r >= height || c < 0 || c >= width) break;
            if (this.seatMap.has(`${r},${c}`)) {
              this.seatMap.get(`${r},${c}`).adjacentOccupied++;
              break;
            }
            i++;
          }
        }
      }
    }
  }
}

function parseInput(input) {
  const seatMap = new Map();
  for (let r = 0; r < height; r++) {
    for (let c = 0; c < width; c++) {
      if (input[r][c] === ".") continue;
      seatMap.set(`${r},${c}`, new Seat(r, c, input[r][c] === "#", seatMap));
    }
  }
  return seatMap;
}

function part1() {
  const seatMap = parseInput(input);
  while (true) {
    let changes = 0;

    for (const seat of seatMap.values()) {
      seat.evalSurroundingSeats();
    }
    for (const seat of seatMap.values()) {
      changes += seat.evalStatus();
    }

    if (changes === 0) {
      return Array.from(seatMap.values()).reduce((a, b) => a + b.occupied, 0);
    }
  }
}

console.log(part1());

function part2() {
  const seatMap = parseInput(input);
  while (true) {
    let changes = 0;

    for (const seat of seatMap.values()) {
      seat.evalSurroundingSeats2();
    }
    for (const seat of seatMap.values()) {
      changes += seat.evalStatus2();
    }

    if (changes === 0) {
      return Array.from(seatMap.values()).reduce((a, b) => a + b.occupied, 0);
    }
  }
}

console.log(part2());
