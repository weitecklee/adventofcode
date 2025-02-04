import * as fs from "fs";
import * as path from "path";

const puzzleInput = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8");

const directions = [
  [-1, 0],
  [0, -1],
  [0, 1],
  [1, 0],
];

type QueueEntry = [number, number[], number[], number[]];

class Unit {
  attackPower: number;
  hitPoints: number;
  pos: [number, number];
  species: string;
  moved: boolean;
  area: (string | Unit)[][];
  unitSetMap: Map<string, Set<Unit>>;

  constructor(
    pos: [number, number],
    species: string,
    attackPower: number,
    area: (string | Unit)[][],
    unitSetMap: Map<string, Set<Unit>>
  ) {
    this.attackPower = attackPower;
    this.hitPoints = 200;
    this.pos = pos;
    this.species = species;
    this.moved = false;
    this.area = area;
    this.unitSetMap = unitSetMap;
  }

  move() {
    if (!this.attackTarget()) {
      this.findTarget();
      this.attackTarget();
    }
    this.moved = true;
  }

  attackTarget(): boolean {
    const targets: Unit[] = [];
    for (const dir of directions) {
      const [r, c] = [this.pos[0] + dir[0], this.pos[1] + dir[1]];
      if (
        this.area[r][c] instanceof Unit &&
        this.area[r][c].species !== this.species
      ) {
        targets.push(this.area[r][c]);
      }
    }
    if (targets.length === 0) return false;
    targets.sort((a, b) => a.hitPoints - b.hitPoints);
    const target = targets[0];
    target.hitPoints -= this.attackPower;
    if (target.hitPoints <= 0) {
      this.area[target.pos[0]][target.pos[1]] = ".";
      this.unitSetMap.get(target.species)!.delete(target);
    }
    return true;
  }

  findTarget() {
    const queue: QueueEntry[] = directions.map((a) => [0, this.pos, a, a]);
    const targets: number[][] = [];
    const visited: Set<string> = new Set([this.pos.join(",")]);
    let dMin = -1;

    for (let i = 0; i < queue.length; i++) {
      const [d, [r, c], move0, [dr, dc]] = queue[i];
      const [r2, c2] = [r + dr, c + dc];
      if (dMin >= 0 && d > dMin) break;
      if (visited.has(`${r2},${c2}`)) continue;
      visited.add(`${r2},${c2}`);
      if (this.area[r2][c2] === "#") continue;
      if (this.area[r2][c2] instanceof Unit) {
        if ((this.area[r2][c2] as Unit).species !== this.species) {
          targets.push([r2, c2, ...move0]);
          dMin = d;
        }
        continue;
      }
      for (const [dr, dc] of directions) {
        queue.push([d + 1, [r2, c2], move0, [dr, dc]]);
      }
    }
    if (targets.length) {
      targets.sort((a, b) => {
        if (a[0] === b[0]) return a[1] - b[1];
        return a[0] - b[0];
      });
      const [r2, c2] = [
        this.pos[0] + targets[0][2],
        this.pos[1] + targets[0][3],
      ];
      this.area[r2][c2] = this;
      this.area[this.pos[0]][this.pos[1]] = ".";
      this.pos = [r2, c2];
    }
  }
}

function simulate(elfAttackPower: number): [number, number] {
  const area: (string | Unit)[][] = puzzleInput
    .split("\n")
    .map((a) => a.split(""));

  const unitSetMap: Map<string, Set<Unit>> = new Map([
    ["Goblin", new Set()],
    ["Elf", new Set()],
  ]);

  for (let r = 0; r < area.length; r++) {
    for (let c = 0; c < area[r].length; c++) {
      if (area[r][c] === "E") {
        area[r][c] = new Unit([r, c], "Elf", elfAttackPower, area, unitSetMap);
        unitSetMap.get("Elf")!.add(area[r][c] as Unit);
      } else if (area[r][c] === "G") {
        area[r][c] = new Unit([r, c], "Goblin", 3, area, unitSetMap);
        unitSetMap.get("Goblin")!.add(area[r][c] as Unit);
      }
    }
  }

  let round = 0;
  let gameOver = false;
  while (true) {
    for (let r = 0; r < area.length; r++) {
      for (let c = 0; c < area[r].length; c++) {
        if (area[r][c] instanceof Unit && !(area[r][c] as Unit).moved) {
          if (
            unitSetMap.get("Elf")!.size === 0 ||
            unitSetMap.get("Goblin")!.size === 0
          ) {
            gameOver = true;
            break;
          }
          (area[r][c] as Unit).move();
        }
      }
      if (gameOver) break;
    }
    if (gameOver) break;
    unitSetMap.get("Elf")!.forEach((u) => {
      u.moved = false;
    });
    unitSetMap.get("Goblin")!.forEach((u) => {
      u.moved = false;
    });
    round++;
  }

  return [
    unitSetMap.get("Elf")!.size,
    round *
      (Array.from(unitSetMap.get("Elf")!).reduce(
        (a, b) => a + Math.max(b.hitPoints, 0),
        0
      ) +
        Array.from(unitSetMap.get("Goblin")!).reduce(
          (a, b) => a + Math.max(b.hitPoints, 0),
          0
        )),
  ];
}

console.log(simulate(3)[1]);

const nElves = puzzleInput.match(/E/g)!.length;

let lo = 4;
let hi = 200;
let part2 = 0;
while (lo < hi) {
  const mid = Math.floor(lo + (hi - lo) / 2);
  const [nElfSurvivors, outcome] = simulate(mid);
  if (nElfSurvivors === nElves) {
    hi = mid;
    part2 = outcome;
  } else {
    lo = mid + 1;
  }
}
console.log(part2);
