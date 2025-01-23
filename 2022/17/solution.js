const fs = require("fs");
const path = require("path");

const input = fs.readFileSync(
  path.join(__dirname, "input.txt"),
  "utf-8",
  (err, data) => {
    if (err) {
      console.log(err);
    } else {
      return data;
    }
  }
);

const chamber = new Set();

class Block {
  constructor(start) {
    this.start = start;
    this.space = [];
  }

  get height() {
    let max = 0;
    for (const pos of this.space) {
      max = Math.max(max, pos[1]);
    }
    return max;
  }

  moveDown() {
    const space2 = [];
    for (const pos of this.space) {
      const pos2 = [pos[0], pos[1] - 1];
      if (pos2[1] === 0 || chamber.has(pos2.join(","))) {
        return false;
      }
      space2.push(pos2);
    }
    this.space = space2;
    return true;
  }

  moveLeft() {
    const space2 = [];
    for (const pos of this.space) {
      const pos2 = [pos[0] - 1, pos[1]];
      if (pos2[0] === 0 || chamber.has(pos2.join(","))) {
        return false;
      }
      space2.push(pos2);
    }
    this.space = space2;
    return true;
  }

  moveRight() {
    const space2 = [];
    for (const pos of this.space) {
      const pos2 = [pos[0] + 1, pos[1]];
      if (pos2[0] === 8 || chamber.has(pos2.join(","))) {
        return false;
      }
      space2.push(pos2);
    }
    this.space = space2;
    return true;
  }
}

class HorizontalBlock extends Block {
  constructor(start) {
    super(start);
    for (let i = 3; i < 7; i++) {
      this.space.push([i, this.start]);
    }
  }
}

class PlusBlock extends Block {
  constructor(start) {
    super(start);
    this.space.push([4, this.start + 2]);
    for (let i = 3; i < 6; i++) {
      this.space.push([i, this.start + 1]);
    }
    this.space.push([4, this.start]);
  }
}

class ReverseLBlock extends Block {
  constructor(start) {
    super(start);
    this.space.push([5, this.start + 2]);
    this.space.push([5, this.start + 1]);
    for (let i = 3; i < 6; i++) {
      this.space.push([i, this.start]);
    }
  }
}

class VerticalBlock extends Block {
  constructor(start) {
    super(start);
    for (let i = 0; i < 4; i++) {
      this.space.push([3, this.start + i]);
    }
  }
}

class SquareBlock extends Block {
  constructor(start) {
    super(start);
    this.space.push([3, this.start]);
    this.space.push([4, this.start]);
    this.space.push([3, this.start + 1]);
    this.space.push([4, this.start + 1]);
  }
}

const newBlock = (n, start) => {
  if (n % 5 === 0) {
    return new HorizontalBlock(start);
  }
  if (n % 5 === 1) {
    return new PlusBlock(start);
  }
  if (n % 5 === 2) {
    return new ReverseLBlock(start);
  }
  if (n % 5 === 3) {
    return new VerticalBlock(start);
  }
  if (n % 5 === 4) {
    return new SquareBlock(start);
  }
};

let height = 0;
let i = 0;
for (let n = 0; n < 2022; n++) {
  const currBlock = newBlock(n, height + 4);
  while (true) {
    if (input[i] === "<") {
      currBlock.moveLeft();
    } else if (input[i] === ">") {
      currBlock.moveRight();
    }
    i++;
    if (i >= input.length) {
      i = 0;
    }
    if (!currBlock.moveDown()) {
      for (const pos of currBlock.space) {
        chamber.add(pos.join(","));
      }
      height = Math.max(height, currBlock.height);
      break;
    }
  }
}

console.log(height);

// Find cycle by tracking height profile after each block has dropped
// Combine this with which block just dropped and place in input
// These three parameters form a cycleID
// In a map, use the cycleID as the key and [n, maxHeight] as the value
// where n = number of blocks dropped, maxHeight = ... you know
// Keep dropping blocks until a repeat cycleID has been found
// From there, you can determine how many blocks have been dropped (the period of the cycle)
// and the height that has been added (the amplitude?)
// Calculate how many cycles were run within 1000000000000 rock drops
// and how many leftover rocks need to be dropped to get to that big number
// Simulate those leftover rock drops and don't forget to add the height from all those cycles

const profile = (h) => {
  const prof = new Array(7).fill(0);
  for (let i = 0; i < 7; i++) {
    let y = h;
    while (y > 0 && !chamber.has(i + 1 + "," + y)) {
      y--;
    }
    prof[i] = h - y;
  }
  return prof.join(".");
};

chamber.clear();
i = 0;
height = 0;

const cycles = new Map();
let cycleFound = false;
let cycleProfile;
let n = 0;

while (!cycleFound) {
  const currBlock = newBlock(n, height + 4);
  while (true) {
    if (input[i] === "<") {
      currBlock.moveLeft();
    } else if (input[i] === ">") {
      currBlock.moveRight();
    }
    i++;
    if (i >= input.length) {
      i = 0;
    }
    if (!currBlock.moveDown()) {
      for (const pos of currBlock.space) {
        chamber.add(pos.join(","));
      }
      height = Math.max(height, currBlock.height);
      break;
    }
  }
  const cycleID = (n % 5) + "|" + i + "|" + profile(height);
  if (!cycleFound && !cycles.has(cycleID)) {
    cycles.set(cycleID, [n, height]);
  } else if (!cycleFound && cycles.has(cycleID)) {
    cycleProfile = cycleID;
    cycleFound = true;
  }
  n++;
}

const [firstN, firstH] = cycles.get(cycleProfile);
const cycleN = n - firstN - 1;
const cycleH = height - firstH;

const simulationN = 1000000000000;
const cyclesRan = Math.floor((simulationN - n) / cycleN);
const leftOverN = simulationN - n - cyclesRan * cycleN;
const heightToAdd = cyclesRan * cycleH;

for (let j = 0; j < leftOverN; j++) {
  const currBlock = newBlock(n + j, height + 4);
  while (true) {
    if (input[i] === "<") {
      currBlock.moveLeft();
    } else if (input[i] === ">") {
      currBlock.moveRight();
    }
    i++;
    if (i >= input.length) {
      i = 0;
    }
    if (!currBlock.moveDown()) {
      for (const pos of currBlock.space) {
        chamber.add(pos.join(","));
      }
      height = Math.max(height, currBlock.height);
      break;
    }
  }
}

console.log(height + heightToAdd);
