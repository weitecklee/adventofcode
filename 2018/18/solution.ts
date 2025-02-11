import * as fs from "fs";
import * as path from "path";

const puzzleInput = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n");

const directions = [
  [-1, 0],
  [1, 0],
  [0, -1],
  [0, 1],
  [-1, -1],
  [-1, 1],
  [1, -1],
  [1, 1],
];

class Acre {
  state: string;
  neighbors: Map<string, number>;
  r: number;
  c: number;
  area: Area;
  constructor(state: string, r: number, c: number, area: Area) {
    this.state = state;
    this.r = r;
    this.c = c;
    this.neighbors = new Map([
      [".", 0],
      ["|", 0],
      ["#", 0],
    ]);
    this.area = area;
  }

  propagate() {
    for (const [dr, dc] of directions) {
      const [r2, c2] = [this.r + dr, this.c + dc];
      if (r2 < 0 || r2 > this.area.rowMax || c2 < 0 || c2 > this.area.colMax)
        continue;
      const neighborMap = this.area.get(r2, c2).neighbors;
      neighborMap.set(this.state, neighborMap.get(this.state)! + 1);
    }
  }

  morph() {
    switch (this.state) {
      case ".":
        if (this.neighbors.get("|")! >= 3) this.state = "|";
        break;
      case "|":
        if (this.neighbors.get("#")! >= 3) this.state = "#";
        break;
      case "#":
        if (this.neighbors.get("#")! < 1 || this.neighbors.get("|")! < 1)
          this.state = ".";
        break;
    }
    this.reset();
  }

  reset() {
    this.neighbors.set(".", 0);
    this.neighbors.set("|", 0);
    this.neighbors.set("#", 0);
  }
}

class Area {
  area: Acre[][];
  constructor(puzzleInput: string[]) {
    this.area = puzzleInput.map((a, r) =>
      a.split("").map((b, c) => new Acre(b, r, c, this))
    );
  }

  get(r: number, c: number): Acre {
    return this.area[r][c];
  }

  get rowMax(): number {
    return this.area.length - 1;
  }

  get colMax(): number {
    return this.area[0].length - 1;
  }

  propagate() {
    for (const row of this.area) {
      for (const acre of row) {
        acre.propagate();
      }
    }
  }

  morph() {
    for (const row of this.area) {
      for (const acre of row) {
        acre.morph();
      }
    }
  }

  iterate() {
    this.propagate();
    this.morph();
  }

  get resourceValue(): number {
    let nWood = 0;
    let nLumber = 0;
    for (const row of this.area) {
      for (const acre of row) {
        if (acre.state === "|") nWood++;
        else if (acre.state === "#") nLumber++;
      }
    }
    return nWood * nLumber;
  }

  get stringForm(): string {
    return this.area.map((a) => a.map((b) => b.state).join("")).join("");
  }
}

function part1(): number {
  const area = new Area(puzzleInput);

  for (let i = 0; i < 10; i++) {
    area.iterate();
  }

  return area.resourceValue;
}

function part2(): number {
  const area = new Area(puzzleInput);
  let i = 0;
  const memo: Map<string, number> = new Map([[area.stringForm, 0]]);
  const memo2: Map<number, string> = new Map([[0, area.stringForm]]);

  while (true) {
    i++;
    area.iterate();
    const stringForm = area.stringForm;
    if (memo.has(stringForm)) {
      break;
    }
    memo.set(stringForm, i);
    memo2.set(i, stringForm);
  }

  const areaString = area.stringForm;
  const period = i - memo.get(areaString)!;
  const timeBeforePeriod = memo.get(areaString)!;
  const timeLeftOver = (1000000000 - timeBeforePeriod) % period;

  return memo2
    .get(timeBeforePeriod + timeLeftOver)!
    .split("")
    .reduce(
      (a, b) => {
        if (b === "|") {
          a[0]++;
        } else if (b === "#") {
          a[1]++;
        }
        return a;
      },
      [0, 0]
    )
    .reduce((a, b) => a * b, 1);
}

console.log(part1());
console.log(part2());
