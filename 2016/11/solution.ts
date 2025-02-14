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
  return (
    floors[floors.length - 1].reduce((a, b) => a + b.size, 0) === totalItems
  );
}

function stateString(floors: Floors, currLoc: number): string {
  return (
    currLoc +
    floors
      .map((a) => a.map((b) => Array.from(b).sort().join(",")).join(";"))
      .join("|")
  );
}

const floors = puzzleInput.map(parseLine);
const totalItems = floors.reduce((a, b) => a + b[0].size + b[1].size, 0);

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
    console.log(steps);
    break;
  }
  steps++;
  for (let i = -1; i <= 1; i += 2) {
    if (currLoc + i < 0 || currLoc + i > floors.length - 1) continue;
    const generators = currFloors[currLoc][0];
    const microchips = currFloors[currLoc][1];
    for (const generator of generators) {
      for (const microchip of microchips) {
        const copy = copyFloors(currFloors);
        copy[currLoc][0].delete(generator);
        copy[currLoc][1].delete(microchip);
        copy[currLoc + i][0].add(generator);
        copy[currLoc + i][1].add(microchip);
        if (!checkFloors(copy)) continue;
        const state = stateString(copy, currLoc + i);
        if (visited.has(state) && visited.get(state)! <= steps) continue;
        visited.set(state, steps);
        MinHeap.push(queue, [
          heuristic(copy) + steps,
          steps,
          currLoc + i,
          copy,
        ]);
      }
      for (const generator2 of generators) {
        const copy = copyFloors(currFloors);
        copy[currLoc][0].delete(generator);
        copy[currLoc][0].delete(generator2);
        copy[currLoc + i][0].add(generator);
        copy[currLoc + i][0].add(generator2);
        if (!checkFloors(copy)) continue;
        const state = stateString(copy, currLoc + i);
        if (visited.has(state) && visited.get(state)! <= steps) continue;
        visited.set(state, steps);
        MinHeap.push(queue, [
          heuristic(copy) + steps,
          steps,
          currLoc + i,
          copy,
        ]);
      }
    }
    for (const microchip of microchips) {
      for (const microchip2 of microchips) {
        const copy = copyFloors(currFloors);
        copy[currLoc][1].delete(microchip);
        copy[currLoc][1].delete(microchip2);
        copy[currLoc + i][1].add(microchip);
        copy[currLoc + i][1].add(microchip2);
        if (!checkFloors(copy)) continue;
        const state = stateString(copy, currLoc + i);
        if (visited.has(state) && visited.get(state)! <= steps) continue;
        visited.set(state, steps);
        MinHeap.push(queue, [
          heuristic(copy) + steps,
          steps,
          currLoc + i,
          copy,
        ]);
      }
    }
  }
}
