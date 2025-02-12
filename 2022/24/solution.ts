import * as fs from "fs";
import * as path from "path";
import MinHeap from "../../utils/MinHeap";

const puzzleInput = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split(""));

const directions = [
  [-1, 0],
  [1, 0],
  [0, -1],
  [0, 1],
  [0, 0],
];

const entrance: [number, number] = [0, 1];
const exit: [number, number] = [
  puzzleInput.length - 1,
  puzzleInput[0].length - 2,
];

function isBlizzard(r: number, c: number, t: number): boolean {
  if (puzzleInput[((r + t - 1) % (puzzleInput.length - 2)) + 1][c] === "^")
    return true;
  if (puzzleInput[r][((c + t - 1) % (puzzleInput[0].length - 2)) + 1] === "<")
    return true;

  if (c - t <= 0) {
    if (
      puzzleInput[r][
        ((c - t) % (puzzleInput[0].length - 2)) + puzzleInput[0].length - 2
      ] === ">"
    )
      return true;
  } else if (puzzleInput[r][c - t] === ">") return true;

  if (r - t <= 0) {
    if (
      puzzleInput[
        ((r - t) % (puzzleInput.length - 2)) + puzzleInput.length - 2
      ][c] === "v"
    )
      return true;
  } else if (puzzleInput[r - t][c] === "v") return true;

  return false;
}

function calcDist(r: number, c: number): number {
  return Math.abs(r - exit[0]) + Math.abs(c - exit[1]);
}

const queue: [number, number, number, number][] = [[0, 0, ...entrance]];
const visited: Set<string> = new Set();

while (queue.length) {
  const [_, t, r, c] = MinHeap.pop(queue) as [number, number, number, number];
  if (visited.has(`${t},${r},${c}`)) continue;
  visited.add(`${t},${r},${c}`);
  if (r === exit[0] && c === exit[1]) {
    console.log(t);
    break;
  }
  for (const [dr, dc] of directions) {
    const [r2, c2] = [r + dr, c + dc];
    if (
      r2 < 0 ||
      c2 < 0 ||
      r2 >= puzzleInput.length ||
      c2 >= puzzleInput[0].length
    )
      continue;
    if (puzzleInput[r2][c2] === "#") continue;
    if (isBlizzard(r2, c2, t + 1)) continue;
    MinHeap.push(queue, [calcDist(r2, c2) + t + 1, t + 1, r2, c2]);
  }
}

const distanceBetweenEnds =
  Math.abs(entrance[0] - exit[0]) + Math.abs(entrance[1] - exit[1]);

function calcDist2(r: number, c: number, endsVisited: number): number {
  if (endsVisited === 0)
    return (
      Math.abs(r - exit[0]) + Math.abs(c - exit[1]) + distanceBetweenEnds * 2
    );
  if (endsVisited === 1)
    return (
      Math.abs(r - entrance[0]) +
      Math.abs(c - entrance[1]) +
      distanceBetweenEnds
    );
  return Math.abs(r - exit[0]) + Math.abs(c - exit[1]);
}

const queue2: [number, number, number, number, number][] = [
  [0, 0, 0, ...entrance],
];
const visited2: Set<string> = new Set();

while (queue2.length) {
  let [_, t, endsVisited, r, c] = MinHeap.pop(queue2) as [
    number,
    number,
    number,
    number,
    number
  ];
  if (visited2.has(`${t},${r},${c}`)) continue;
  visited2.add(`${t},${r},${c}`);
  if (r === exit[0] && c === exit[1]) {
    if (endsVisited === 2) {
      console.log(t);
      break;
    } else if (endsVisited === 0) {
      endsVisited = 1;
    }
  } else if (r === entrance[0] && c === entrance[1]) {
    if (endsVisited === 1) {
      endsVisited = 2;
    }
  }
  for (const [dr, dc] of directions) {
    const [r2, c2] = [r + dr, c + dc];
    if (
      r2 < 0 ||
      c2 < 0 ||
      r2 >= puzzleInput.length ||
      c2 >= puzzleInput[0].length
    )
      continue;
    if (puzzleInput[r2][c2] === "#") continue;
    if (isBlizzard(r2, c2, t + 1)) continue;
    MinHeap.push(queue2, [
      calcDist2(r2, c2, endsVisited) + t + 1,
      t + 1,
      endsVisited,
      r2,
      c2,
    ]);
  }
}
