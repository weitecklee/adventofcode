const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n");

class Pad {
  constructor() {
    this.pad;
    this.keyPositions;
    this.currentKey;
    this.memo = new Map();
  }

  initialize() {
    this.keyPositions = new Map();
    for (let r = 0; r < this.pad.length; r++) {
      for (let c = 0; c < this.pad[0].length; c++) {
        this.keyPositions.set(this.pad[r][c], [r, c]);
      }
    }
    this.currentKey = "A";
  }

  maneuverTo(target) {
    const memoKey = `${this.currentKey}-${target}`;
    if (this.memo.has(memoKey)) {
      this.currentKey = target;
      return this.memo.get(memoKey);
    }
    const [r, c] = this.keyPositions.get(this.currentKey);
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
    this.currentKey = target;
    if ((r === rBlank && c2 === cBlank) || (r2 === rBlank && c === cBlank))
      moves = moves.reverse();
    moves.push("A");
    moves = moves.join("");
    this.memo.set(memoKey, moves);
    return moves;
  }

  type(code) {
    const res = [];
    for (const c of code) {
      res.push(this.maneuverTo(c));
    }
    return res.join("");
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

function minLength(arr) {
  let min = arr[0].length;
  arr.forEach((a) => {
    min = Math.min(min, a.length);
  });
  return min;
}

function solve(code, n) {
  let res = numpad.type(code);
  for (let i = 0; i < n; i++) {
    // console.log(n, res.length);
    res = keypad.type(res);
  }
  return res;
}

function complexity(code) {
  return type(code).length * Number(code.slice(0, 3));
}

let part1 = 0;
for (const code of input) {
  const res = solve(code, 2);
  part1 += res.length * Number(code.slice(0, 3));
}
console.log(part1);

/*
  879A 70
  508A 72
  463A 70
  593A 74
  189A 74
*/
