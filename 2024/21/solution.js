const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n");

function convertCodeToMoveCounts(code) {
  let moveCounts = new Map();
  code.split("").forEach((c, i) => {
    if (i === 0) {
      moveCounts.set(`A-${c}`, 1);
    } else {
      moveCounts.set(
        `${code[i - 1]}-${c}`,
        (moveCounts.get(`${code[i - 1]}-${c}`) || 0) + 1
      );
    }
  });
  return moveCounts;
}

class Pad {
  constructor() {
    this.pad;
    this.keyPositions;
    this.memo = new Map();
  }

  initialize() {
    this.keyPositions = new Map();
    for (let r = 0; r < this.pad.length; r++) {
      for (let c = 0; c < this.pad[0].length; c++) {
        this.keyPositions.set(this.pad[r][c], [r, c]);
      }
    }
  }

  maneuver(movement) {
    // returns string of button presses for movement (e.g., "A-<", "0-1", etc.)
    if (this.memo.has(movement)) {
      return this.memo.get(movement);
    }
    const [curr, target] = movement.split("-");
    const [r, c] = this.keyPositions.get(curr);
    const [r2, c2] = this.keyPositions.get(target);
    let moves = [];
    for (let i = c; i > c2; i--) {
      moves.push("<");
    }
    for (let i = r; i < r2; i++) {
      moves.push("v");
    }
    for (let i = r; i > r2; i--) {
      moves.push("^");
    }
    for (let i = c; i < c2; i++) {
      moves.push(">");
    }
    const [rBlank, cBlank] = this.keyPositions.get("_");
    if ((r === rBlank && c2 === cBlank) || (r2 === rBlank && c === cBlank))
      moves = moves.reverse();
    moves.push("A");
    moves = moves.join("");
    this.memo.set(movement, moves);
    return moves;
  }

  typeIn(moveCounts) {
    const newMoveCounts = new Map();
    for (const [movement, n] of moveCounts) {
      const code = this.maneuver(movement);
      const tmp = convertCodeToMoveCounts(code);
      for (const [mvmt, a] of tmp) {
        newMoveCounts.set(mvmt, (newMoveCounts.get(mvmt) || 0) + a * n);
      }
    }
    return newMoveCounts;
  }
}

class Numpad extends Pad {
  constructor() {
    super();
    this.pad = [
      ["7", "8", "9"],
      ["4", "5", "6"],
      ["1", "2", "3"],
      ["_", "0", "A"],
    ];
    this.initialize();
  }
}

class Keypad extends Pad {
  constructor() {
    super();
    this.pad = [
      ["_", "^", "A"],
      ["<", "v", ">"],
    ];
    this.initialize();
  }
}

const numpad = new Numpad();
const keypad = new Keypad();

function solve(code, nRobots) {
  let moveCounts = convertCodeToMoveCounts(code);
  moveCounts = numpad.typeIn(moveCounts);
  for (let i = 0; i < nRobots; i++) {
    moveCounts = keypad.typeIn(moveCounts);
  }
  return (
    Array.from(moveCounts.values()).reduce((a, b) => a + b) *
    Number(code.slice(0, 3))
  );
}

let part1 = 0;
for (const code of input) {
  part1 += solve(code, 2);
}
console.log(part1);

let part2 = 0;
for (const code of input) {
  part2 += solve(code, 25);
}
console.log(part2);
