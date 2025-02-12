import * as fs from "fs";
import * as path from "path";

const puzzleInput = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split(""));

function countAdjacentElves(
  r: number,
  c: number,
  elves: Set<string>
): [boolean, number[]] {
  const adjacentElves: number[] = Array(4).fill(0);
  if (elves.has(`${r - 1},${c}`)) {
    adjacentElves[0]++;
  }
  if (elves.has(`${r - 1},${c + 1}`)) {
    adjacentElves[0]++;
    adjacentElves[3]++;
  }
  if (elves.has(`${r},${c + 1}`)) {
    adjacentElves[3]++;
  }
  if (elves.has(`${r + 1},${c + 1}`)) {
    adjacentElves[1]++;
    adjacentElves[3]++;
  }
  if (elves.has(`${r + 1},${c}`)) {
    adjacentElves[1]++;
  }
  if (elves.has(`${r + 1},${c - 1}`)) {
    adjacentElves[1]++;
    adjacentElves[2]++;
  }
  if (elves.has(`${r},${c - 1}`)) {
    adjacentElves[2]++;
  }
  if (elves.has(`${r - 1},${c - 1}`)) {
    adjacentElves[0]++;
    adjacentElves[2]++;
  }
  return [adjacentElves.some((a) => a > 0), adjacentElves];
}

const directions = [
  [-1, 0],
  [1, 0],
  [0, -1],
  [0, 1],
];

function moveTheElves(
  elves: Set<string>,
  round: number
): [Set<string>, boolean] {
  const newElves: Set<string> = new Set();
  const proposals: Map<string, string> = new Map();
  const proposalCounts: Map<string, number> = new Map();
  let movement = false;
  for (const elf of elves) {
    const [r, c] = elf.split(",").map(Number);
    const [anyAdjacent, adjacentElves] = countAdjacentElves(r, c, elves);
    let [r2, c2] = [r, c];
    if (anyAdjacent) {
      for (let i = 0; i < 4; i++) {
        const j = (round + i) % 4;
        if (adjacentElves[j] === 0) {
          r2 += directions[j][0];
          c2 += directions[j][1];
          break;
        }
      }
    }
    proposals.set(elf, `${r2},${c2}`);
    proposalCounts.set(
      `${r2},${c2}`,
      (proposalCounts.get(`${r2},${c2}`) || 0) + 1
    );
  }
  for (const [elf, proposal] of proposals) {
    if (proposalCounts.get(proposal)! > 1) {
      newElves.add(elf);
    } else {
      newElves.add(proposal);
      if (elf !== proposal) movement = true;
    }
  }
  return [newElves, movement];
}

function part1(puzzleInput: string[][]): number {
  let elves: Set<string> = new Set();

  for (let r = 0; r < puzzleInput.length; r++) {
    for (let c = 0; c < puzzleInput[0].length; c++) {
      if (puzzleInput[r][c] === "#") elves.add(`${r},${c}`);
    }
  }

  for (let i = 0; i < 10; i++) {
    elves = moveTheElves(elves, i)[0];
  }

  const elfCoords = Array.from(elves).map((a) => a.split(",").map(Number));

  let rMin = Number.MAX_SAFE_INTEGER;
  let rMax = Number.MIN_SAFE_INTEGER;
  let cMin = rMin;
  let cMax = rMax;

  for (const [r, c] of elfCoords) {
    rMin = Math.min(rMin, r);
    rMax = Math.max(rMax, r);
    cMin = Math.min(cMin, c);
    cMax = Math.max(cMax, c);
  }

  return (rMax - rMin + 1) * (cMax - cMin + 1) - elves.size;
}

function part2(puzzleInput: string[][]): number {
  let elves: Set<string> = new Set();

  for (let r = 0; r < puzzleInput.length; r++) {
    for (let c = 0; c < puzzleInput[0].length; c++) {
      if (puzzleInput[r][c] === "#") elves.add(`${r},${c}`);
    }
  }

  let round = -1;
  let movement = true;
  while (movement) {
    round++;
    [elves, movement] = moveTheElves(elves, round);
  }
  return round + 1;
}

console.log(part1(puzzleInput));
console.log(part2(puzzleInput));
