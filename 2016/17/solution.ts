import * as fs from "fs";
import * as path from "path";
import * as crypto from "crypto";
import MinHeap from "../../utils/MinHeap";

const puzzleInput = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8");

const RMAX = 4 - 1;
const CMAX = 4 - 1;
const directions: [string, [number, number]][] = [
  ["U", [-1, 0]],
  ["D", [1, 0]],
  ["L", [0, -1]],
  ["R", [0, 1]],
];

const openRegex = /[bcdef]/;
function openDoors(s: string): boolean[] {
  return crypto
    .createHash("md5")
    .update(s)
    .digest("hex")
    .slice(0, 4)
    .split("")
    .map((a) => openRegex.test(a));
}

function heuristic(r: number, c: number): number {
  return Math.abs(r - RMAX) + Math.abs(c - CMAX);
}

function solve(puzzleInput: string): [string, number] {
  type QueueEntry = [number, number, string, number, number];
  const queue: QueueEntry[] = [[0, 0, "", 0, 0]];
  let minRoute = "";
  let maxSteps = 0;

  while (queue.length) {
    let [_, steps, route, r, c] = MinHeap.pop(queue) as QueueEntry;

    if (r === RMAX && c === CMAX) {
      if (!minRoute) minRoute = route;
      if (steps > maxSteps) maxSteps = steps;
      continue;
    }

    steps++;
    const doors = openDoors(puzzleInput + route);
    for (let i = 0; i < doors.length; i++) {
      if (!doors[i]) continue;
      const [dir, [dr, dc]] = directions[i];
      const [r2, c2] = [r + dr, c + dc];
      if (r2 < 0 || c2 < 0 || r2 > RMAX || c2 > CMAX) continue;
      MinHeap.push(queue, [
        heuristic(r2, c2) + steps,
        steps,
        route + dir,
        r2,
        c2,
      ]);
    }
  }
  return [minRoute, maxSteps];
}

const [part1, part2] = solve(puzzleInput);
console.log(part1, part2);
