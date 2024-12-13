const fs = require("fs");
const path = require("path");
const mathjs = require("mathjs");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n\n")
  .map((a) => a.split("\n"));

class Machine {
  constructor(lines) {
    this.buttonA = lines[0].match(/\d+/g).map(Number);
    this.buttonB = lines[1].match(/\d+/g).map(Number);
    this.prize = lines[2].match(/\d+/g).map(Number);
  }

  get tokenCost() {
    const presses = mathjs
      .lusolve(
        mathjs.matrix([
          [this.buttonA[0], this.buttonB[0]],
          [this.buttonA[1], this.buttonB[1]],
        ]),
        mathjs.matrix(this.prize)
      )
      ._data.flat();
    if (!presses.every((a) => Math.abs(a - Math.round(a)) < 1e-6)) {
      return 0;
    }
    return presses[0] * 3 + presses[1];
  }

  get tokenCost2() {
    const presses = mathjs
      .lusolve(
        mathjs.matrix([
          [this.buttonA[0], this.buttonB[0]],
          [this.buttonA[1], this.buttonB[1]],
        ]),
        mathjs.matrix([
          this.prize[0] + 10000000000000,
          this.prize[1] + 10000000000000,
        ])
      )
      ._data.flat();
    if (!presses.every((a) => Math.abs(a - Math.round(a)) < 0.0001)) {
      return 0;
    }
    return presses[0] * 3 + presses[1];
  }
}

const machines = input.map((a) => new Machine(a));
console.log(machines.reduce((a, b) => a + b.tokenCost, 0));
console.log(machines.reduce((a, b) => a + b.tokenCost2, 0));
