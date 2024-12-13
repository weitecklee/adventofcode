const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n\n")
  .map((a) => a.split("\n"));

function solve(A, b) {
  const det = A[0][0] * A[1][1] - A[0][1] * A[1][0];
  if (det === 0) {
    return [0, 0];
  }
  return [
    (b[0] * A[1][1] - b[1] * A[0][1]) / det,
    (A[0][0] * b[1] - A[1][0] * b[0]) / det,
  ];
}

class Machine {
  constructor(lines) {
    this.buttonA = lines[0].match(/\d+/g).map(Number);
    this.buttonB = lines[1].match(/\d+/g).map(Number);
    this.prize = lines[2].match(/\d+/g).map(Number);
    this.prize2 = this.prize.map((a) => a + 10000000000000);
  }

  tokenCost(part2 = false) {
    const [pressesA, pressesB] = solve(
      [
        [this.buttonA[0], this.buttonB[0]],
        [this.buttonA[1], this.buttonB[1]],
      ],
      part2 ? this.prize2 : this.prize
    );
    if (Number.isInteger(pressesA) && Number.isInteger(pressesB)) {
      return pressesA * 3 + pressesB;
    }
    return 0;
  }
}

const machines = input.map((a) => new Machine(a));
console.log(machines.reduce((a, b) => a + b.tokenCost(), 0));
console.log(machines.reduce((a, b) => a + b.tokenCost(true), 0));
