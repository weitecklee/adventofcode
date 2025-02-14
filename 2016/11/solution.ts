import * as fs from "fs";
import * as path from "path";
import MinHeap from "../../utils/MinHeap";

const puzzleInput = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n");

const generatorRegex = /\w+(?= generator)/g;
const microchipRegex = /\w+(?=\-compatible)/g;

type Floor = [Set<string>, Set<string>];
type Floors = Floor[];

function parseLine(line: string): Floor {
  const generators: Set<string> = new Set(line.match(generatorRegex) || []);
  const microchips: Set<string> = new Set(line.match(microchipRegex) || []);
  return [generators, microchips];
}

function copyFloors(floors: Floors): Floors {
  return floors.map((a) => [new Set(a[0]), new Set(a[1])]);
}

function checkFloors(floors: Floors): boolean {
  for (const [generators, microchips] of floors) {
    for (const microchip of microchips) {
      if (!generators.has(microchip) && generators.size > 0) return false;
    }
  }
  return true;
}

function heuristic(floors: Floors): number {
  return floors.reduce(
    (a, b, i) => a + (b[0].size + b[1].size) * (floors.length - i),
    0
  );
}

function isDone(floors: Floors): boolean {
  for (let i = 0; i < floors.length - 1; i++) {
    if (floors[i][0].size > 0 || floors[i][1].size > 0) return false;
  }
  return true;
}

function stateString(floors: Floors, currLoc: number): string {
  return (
    currLoc +
    floors
      .map((a) => a.map((b) => Array.from(b).sort().join(",")).join(";"))
      .join("|")
  );
}

function pushNewEntry(
  floors: Floors,
  currLoc: number,
  destLoc: number,
  steps: number,
  takenGenerators: string[],
  takenMicrochips: string[],
  queue: [number, number, number, Floors][],
  visited: Map<string, number>
) {
  const copy = copyFloors(floors);
  for (const generator of takenGenerators) {
    copy[currLoc][0].delete(generator);
    copy[destLoc][0].add(generator);
  }
  for (const microchip of takenMicrochips) {
    copy[currLoc][1].delete(microchip);
    copy[destLoc][1].add(microchip);
  }
  if (!checkFloors(copy)) return;
  const state = stateString(copy, destLoc);
  if (visited.has(state) && visited.get(state)! <= steps) return;
  visited.set(state, steps);
  MinHeap.push(queue, [
    heuristic(copy) * 0.9 + steps * 0.1,
    steps,
    destLoc,
    copy,
  ]);
}

const floors = puzzleInput.map(parseLine);

function solve(floors: Floors): number {
  const queue: [number, number, number, Floors][] = [[0, 0, 0, floors]];
  const visited: Map<string, number> = new Map();

  while (queue.length) {
    let [_, steps, currLoc, currFloors] = MinHeap.pop(queue) as [
      number,
      number,
      number,
      Floors
    ];
    if (isDone(currFloors)) {
      return steps;
    }

    steps++;
    for (let i = -1; i <= 1; i += 2) {
      if (currLoc + i < 0 || currLoc + i > floors.length - 1) continue;
      const generators = currFloors[currLoc][0];
      const microchips = currFloors[currLoc][1];
      for (const generator of generators) {
        for (const microchip of microchips) {
          pushNewEntry(
            currFloors,
            currLoc,
            currLoc + i,
            steps,
            [generator],
            [microchip],
            queue,
            visited
          );
        }
        for (const generator2 of generators) {
          pushNewEntry(
            currFloors,
            currLoc,
            currLoc + i,
            steps,
            [generator, generator2],
            [],
            queue,
            visited
          );
        }
      }
      for (const microchip of microchips) {
        for (const microchip2 of microchips) {
          pushNewEntry(
            currFloors,
            currLoc,
            currLoc + i,
            steps,
            [],
            [microchip, microchip2],
            queue,
            visited
          );
        }
      }
    }
  }

  return -1;
}

console.log(solve(floors));

floors[0][0].add("elerium");
floors[0][0].add("dilithium");
floors[0][1].add("elerium");
floors[0][1].add("dilithium");

console.log(solve(floors));
