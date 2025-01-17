import * as fs from "fs";
import * as path from "path";

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split("-"));

class Cave {
  name: string;
  isSmall: boolean;
  neighbors: Cave[];

  constructor(name: string) {
    this.name = name;
    this.isSmall = name[0] === name[0].toLowerCase();
    this.neighbors = [];
  }

  addNeighbor(cave: Cave) {
    this.neighbors.push(cave);
  }
}

const caveMap: Map<string, Cave> = new Map();

for (const [name1, name2] of input) {
  if (!caveMap.has(name1)) caveMap.set(name1, new Cave(name1));
  if (!caveMap.has(name2)) caveMap.set(name2, new Cave(name2));
  const cave1 = caveMap.get(name1)!;
  const cave2 = caveMap.get(name2)!;
  cave1.addNeighbor(cave2);
  cave2.addNeighbor(cave1);
}

let part1 = 0;

interface QueueEntry {
  current: Cave;
  visited: Set<Cave>;
}

const queue: QueueEntry[] = [
  { current: caveMap.get("start")!, visited: new Set([caveMap.get("start")!]) },
];

while (queue.length) {
  const { current, visited } = queue.pop()!;
  if (current.name === "end") {
    part1++;
    continue;
  }
  for (const neighbor of current.neighbors) {
    if (neighbor.isSmall && visited.has(neighbor)) continue;
    const visited2 = new Set(visited);
    visited2.add(neighbor);
    queue.push({ current: neighbor, visited: visited2 });
  }
}

console.log(part1);

let part2 = 0;

interface QueueEntry2 {
  current: Cave;
  visited: Set<Cave>;
  revisitedSmall: boolean;
}

const queue2: QueueEntry2[] = [
  {
    current: caveMap.get("start")!,
    visited: new Set([caveMap.get("start")!]),
    revisitedSmall: false,
  },
];

while (queue2.length) {
  const { current, visited, revisitedSmall } = queue2.pop()!;
  if (current.name === "end") {
    part2++;
    continue;
  }
  for (const neighbor of current.neighbors) {
    if (neighbor.name === "start") continue;
    let revisitedSmall2 = revisitedSmall;
    if (neighbor.isSmall && visited.has(neighbor)) {
      if (!revisitedSmall) revisitedSmall2 = true;
      else continue;
    }
    const visited2 = new Set(visited);
    visited2.add(neighbor);
    queue2.push({
      current: neighbor,
      visited: visited2,
      revisitedSmall: revisitedSmall2,
    });
  }
}

console.log(part2);
