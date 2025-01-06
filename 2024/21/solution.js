const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n");

/*
  Cannot just return full movement string, quickly becomes untenable for part 2.
  Instead, use a moveCounts map to keep track of how many times each movement (in
  the form of "A->", "<-^", etc.) is made after each layer of robot.
  For each movement, Keypad.maneuver(movement) will determine the optimal move string,
  which is then converted to a moveCounts map and returned.
  Keypad.typeIn(moveCounts) is then used to iterate on the current moveCounts map,
  calculate the necessary moves and counts to simulate that current moveCounts map,
  and return a new moveCounts map.
*/

function convertCodeToMoveCounts(code) {
  const moveCounts = new Map();
  code.split("").forEach((c, i) => {
    if (i === 0) {
      // always start with robot on "A" button
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

class Keypad {
  constructor() {
    this.pad;
    this.keyPositions = new Map();
    this.memo = new Map();
  }

  initialize() {
    // initialize keyPositions map
    for (let r = 0; r < this.pad.length; r++) {
      for (let c = 0; c < this.pad[0].length; c++) {
        this.keyPositions.set(this.pad[r][c], [r, c]);
      }
    }
  }

  maneuver(movement) {
    if (this.memo.has(movement)) {
      return this.memo.get(movement);
    }
    const [curr, target] = movement.split("-");
    const [r, c] = this.keyPositions.get(curr);
    const [r2, c2] = this.keyPositions.get(target);
    // Optimal move string is determined by 2 rules:
    // 1. never change direction more than once (e.g., "<<^" or "^<<" are better than "<^<")
    // 2. favor buttons farther away from "A"
    //    UNLESS this causes robot to go over gap ("_" button).
    // With the setup of the dirpad, "<" button is always favored most,
    // followed by "v", then "^" and ">". Then there is a check to see if robot goes over gap.
    // If there is, simply reverse the move string. This works because gap is always at a
    // corner. If "<<^" causes robot to go over gap, go the reverse way "^<<" instead.
    // Always end move string with "A".
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
    const moveCounts = convertCodeToMoveCounts(moves);
    this.memo.set(movement, moveCounts);
    return moveCounts;
  }

  typeIn(moveCounts) {
    const newMoveCounts = new Map();
    for (const [movement, n] of moveCounts) {
      const tmp = this.maneuver(movement);
      for (const [mvmt, a] of tmp) {
        newMoveCounts.set(mvmt, (newMoveCounts.get(mvmt) || 0) + a * n);
      }
    }
    return newMoveCounts;
  }
}

class Numpad extends Keypad {
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

class Dirpad extends Keypad {
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
const dirpad = new Dirpad();

function solve(code, nDirectionalRobots) {
  // make first moveCounts map
  let moveCounts = convertCodeToMoveCounts(code);
  // iterate with numpad first
  moveCounts = numpad.typeIn(moveCounts);
  // iterate for however many robots with directional keypads
  for (let i = 0; i < nDirectionalRobots; i++) {
    moveCounts = dirpad.typeIn(moveCounts);
  }
  // calculate complexity
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
