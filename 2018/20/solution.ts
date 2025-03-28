import * as fs from "fs";
import * as path from "path";

const puzzleInput = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8");

// use stack to keep track of branching points.
// when `(` is encountered, a branch is started, add current pos to stack.
// when `|` is encountered, go back to start of branch, top of stack becomes current pos again.
// when `)` is encounted, branch is done, pop value off stack to become new current pos.

// facilityMap is map of <pos, isRoom>
// isRoom = !isDoor

function parseInput(puzzleInput: string): Map<string, boolean> {
  const facilityMap: Map<string, boolean> = new Map();
  facilityMap.set("0,0", true);
  const posStack: number[][] = [];
  let pos = [0, 0];
  for (const c of puzzleInput) {
    switch (c) {
      case "N":
        pos[1]++;
        facilityMap.set(pos.join(","), false);
        pos[1]++;
        facilityMap.set(pos.join(","), true);
        break;
      case "E":
        pos[0]++;
        facilityMap.set(pos.join(","), false);
        pos[0]++;
        facilityMap.set(pos.join(","), true);
        break;
      case "W":
        pos[0]--;
        facilityMap.set(pos.join(","), false);
        pos[0]--;
        facilityMap.set(pos.join(","), true);
        break;
      case "S":
        pos[1]--;
        facilityMap.set(pos.join(","), false);
        pos[1]--;
        facilityMap.set(pos.join(","), true);
        break;
      case "(":
        posStack.push(pos.slice());
        break;
      case "|":
        pos = posStack[posStack.length - 1].slice();
        break;
      case ")":
        pos = posStack.pop()!;
        break;
    }
  }
  return facilityMap;
}

const directions = [
  [-1, 0],
  [1, 0],
  [0, -1],
  [0, 1],
];

function solve(puzzleInput: string): [number, number] {
  const facilityMap = parseInput(puzzleInput);
  let part1 = 0;
  let part2 = 0;
  const visited: Set<string> = new Set();
  visited.add("0,0");
  const queue: [number[], number][] = [[[0, 0], 0]];
  let i = 0;
  while (i < queue.length) {
    let [pos, doors] = queue[i];
    i++;
    if (facilityMap.get(pos.join(","))!) {
      part1 = Math.max(part1, doors);
      if (doors >= 1000) {
        part2++;
      }
    } else {
      doors++;
    }
    for (const [dx, dy] of directions) {
      const pos2 = [pos[0] + dx, pos[1] + dy];
      const pos2String = pos2.join(",");
      if (!facilityMap.has(pos2String)) {
        // any pos not in facilityMap is not room or door, i.e., it's a wall
        continue;
      }
      if (visited.has(pos2String)) {
        continue;
      }
      visited.add(pos2String);
      queue.push([pos2, doors]);
    }
  }

  return [part1, part2];
}

console.log(solve(puzzleInput));
