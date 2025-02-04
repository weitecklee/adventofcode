import * as fs from "fs";
import * as path from "path";

const puzzleInput: (string | Unit)[][] = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split(""));

const directions = [
  [-1, 0],
  [0, -1],
  [0, 1],
  [1, 0],
];

const unitCounts: Map<string, number> = new Map([
  ["Goblin", 0],
  ["Elf", 0],
]);

const unitMap: Map<string, Unit[]> = new Map([
  ["Goblin", []],
  ["Elf", []],
]);

type QueueEntry = [number, number[], number[], number[]];

class Unit {
  attackPower: number;
  hitPoints: number;
  pos: [number, number];
  species: string;
  moved: boolean;

  constructor(pos: [number, number], species: string) {
    this.attackPower = 3;
    this.hitPoints = 200;
    this.pos = pos;
    this.species = species;
    this.moved = false;
    unitCounts.set(species, (unitCounts.get(species) || 0) + 1);
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
        puzzleInput[r][c] instanceof Unit &&
        puzzleInput[r][c].species !== this.species
      ) {
        targets.push(puzzleInput[r][c]);
      }
    }
    if (targets.length === 0) return false;
    targets.sort((a, b) => a.hitPoints - b.hitPoints);
    targets[0].hitPoints -= this.attackPower;
    if (targets[0].hitPoints <= 0) {
      puzzleInput[targets[0].pos[0]][targets[0].pos[1]] = ".";
      unitCounts.set(
        targets[0].species,
        unitCounts.get(targets[0].species)! - 1
      );
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
      if (puzzleInput[r2][c2] === "#") continue;
      if (puzzleInput[r2][c2] instanceof Unit) {
        if ((puzzleInput[r2][c2] as Unit).species !== this.species) {
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
      puzzleInput[r2][c2] = this;
      puzzleInput[this.pos[0]][this.pos[1]] = ".";
      this.pos = [r2, c2];
    }
  }
}

for (let r = 0; r < puzzleInput.length; r++) {
  for (let c = 0; c < puzzleInput[r].length; c++) {
    if (puzzleInput[r][c] === "E") {
      puzzleInput[r][c] = new Unit([r, c], "Elf");
      unitMap.get("Elf")!.push(puzzleInput[r][c] as Unit);
    } else if (puzzleInput[r][c] === "G") {
      puzzleInput[r][c] = new Unit([r, c], "Goblin");
      unitMap.get("Goblin")!.push(puzzleInput[r][c] as Unit);
    }
  }
}

let round = 0;
let gameOver = false;
while (true) {
  for (let r = 0; r < puzzleInput.length; r++) {
    for (let c = 0; c < puzzleInput[r].length; c++) {
      if (
        puzzleInput[r][c] instanceof Unit &&
        !(puzzleInput[r][c] as Unit).moved
      ) {
        if (unitCounts.get("Elf") === 0 || unitCounts.get("Goblin") === 0) {
          gameOver = true;
          break;
        }
        (puzzleInput[r][c] as Unit).move();
      }
    }
    if (gameOver) break;
  }
  if (gameOver) break;
  unitMap.get("Elf")!.forEach((u) => {
    u.moved = false;
  });
  unitMap.get("Goblin")!.forEach((u) => {
    u.moved = false;
  });
  round++;
}

console.log(
  round *
    (unitMap.get("Elf")!.reduce((a, b) => a + Math.max(b.hitPoints, 0), 0) +
      unitMap.get("Goblin")!.reduce((a, b) => a + Math.max(b.hitPoints, 0), 0))
);
