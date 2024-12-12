const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split(""));

const directions = [
  [0, 1],
  [1, 0],
  [-1, 0],
  [0, -1],
];

class Region {
  constructor(set) {
    this.set = set;
  }

  get area() {
    return this.set.size;
  }

  get perimeter() {
    let perimeter = 0;
    for (const coord of this.set) {
      const [r, c] = coord.split(",").map(Number);
      for (const [dr, dc] of directions) {
        if (!this.set.has([r + dr, c + dc].join(","))) {
          perimeter++;
        }
      }
    }
    return perimeter;
  }

  get sides() {
    if (this.set.size <= 2) return 4;
    const coords = Array.from(this.set).map((a) => a.split(",").map(Number));
    const [minR, maxR] = [
      Math.min(...coords.map((a) => a[0])),
      Math.max(...coords.map((a) => a[0])),
    ];
    const [minC, maxC] = [
      Math.min(...coords.map((a) => a[1])),
      Math.max(...coords.map((a) => a[1])),
    ];
    let count = 0;
    for (let r = minR; r <= maxR + 1; r++) {
      for (let c = minC; c <= maxC + 1; c++) {
        // e | f
        //-------
        // g | h
        const e = this.set.has(`${r - 1},${c - 1}`);
        const f = this.set.has(`${r - 1},${c}`);
        const g = this.set.has(`${r},${c - 1}`);
        const h = this.set.has(`${r},${c}`);
        if (h !== g && ((h === f && f === e) || h !== f)) {
          count++;
        }
      }
    }
    for (let c = minC; c <= maxC + 1; c++) {
      for (let r = minR; r <= maxR + 1; r++) {
        // e | f
        //-------
        // g | h
        const e = this.set.has(`${r - 1},${c - 1}`);
        const f = this.set.has(`${r - 1},${c}`);
        const g = this.set.has(`${r},${c - 1}`);
        const h = this.set.has(`${r},${c}`);
        if (h !== f && ((h === g && g === e) || h !== g)) {
          count++;
        }
      }
    }
    return count;
  }

  get price() {
    return this.area * this.perimeter;
  }

  get discountPrice() {
    return this.area * this.sides;
  }
}

function mapRegion(r, c) {
  const region = new Set();
  const regionCharacter = input[r][c];
  const queue = [[r, c]];
  while (queue.length) {
    const [r, c] = queue.shift();
    region.add(`${r},${c}`);
    for (const [dr, dc] of directions) {
      const r2 = r + dr;
      const c2 = c + dc;
      if (r2 < 0 || r2 >= input.length || c2 < 0 || c2 >= input[0].length)
        continue;
      if (input[r2][c2] !== regionCharacter) continue;
      input[r2][c2] = " ";
      queue.push([r2, c2]);
    }
  }
  return new Region(region);
}

const regions = [];
for (let i = 0; i < input.length; i++) {
  for (let j = 0; j < input[0].length; j++) {
    if (input[i][j] === " ") continue;
    regions.push(mapRegion(i, j));
  }
}

console.log(regions.reduce((a, b) => a + b.price, 0));
console.log(regions.reduce((a, b) => a + b.discountPrice, 0));
