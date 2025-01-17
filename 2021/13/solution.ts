import * as fs from "fs";
import * as path from "path";

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n\n");

class Dot {
  x: number;
  y: number;

  constructor(x: number, y: number) {
    this.x = x;
    this.y = y;
  }

  fold([x, y]: number[]) {
    if (x > 0 && this.x > x) {
      this.x = 2 * x - this.x;
    } else if (y > 0 && this.y > y) {
      this.y = 2 * y - this.y;
    }
  }

  get coords() {
    return `${this.x},${this.y}`;
  }
}

let dots = input[0].split("\n").map((a) => {
  const [x, y] = a.split(",").map(Number);
  return new Dot(x, y);
});

const folds = input[1].split("\n").map((a) => {
  const [partA, partB] = a.split("=");
  if (partA[partA.length - 1] === "x") {
    return [Number(partB), 0];
  }
  return [0, Number(partB)];
});

function foldTheDots(dots: Dot[], fold: number[]) {
  const coordsSet: Set<string> = new Set();
  const dots2: Dot[] = [];
  for (const dot of dots) {
    dot.fold(fold);
    if (!coordsSet.has(dot.coords)) {
      coordsSet.add(dot.coords);
      dots2.push(dot);
    }
  }
  return dots2;
}

dots = foldTheDots(dots, folds[0]);
console.log(dots.length);

for (let i = 1; i < folds.length; i++) {
  dots = foldTheDots(dots, folds[i]);
}

const xMax = Math.max(...dots.map((dot) => dot.x));
const yMax = Math.max(...dots.map((dot) => dot.y));

const res: string[][] = [];
for (let i = 0; i <= yMax; i++) {
  res.push(new Array(xMax + 1).fill(" "));
}

dots.forEach((dot) => {
  res[dot.y][dot.x] = "#";
});

res.forEach((r) => console.log(r.join("")));
