import * as fs from "fs";
import * as path from "path";
import MinHeap from "../../utils/MinHeap";

const puzzleInput = Number(
  fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8")
);

const directions = [
  [0, 1],
  [1, 0],
  [0, -1],
  [-1, 0],
];

function isOpenSpace(x: number, y: number, puzzleInput: number): boolean {
  return (
    (x * x + 3 * x + 2 * x * y + y + y * y + puzzleInput)
      .toString(2)
      .replaceAll("0", "").length %
      2 ===
    0
  );
}

interface State {
  steps: number;
  pos: [number, number];
}

function part1(puzzleInput: number): number {
  const xGoal = 31;
  const yGoal = 39;

  const queue: [number, State][] = [[0, { steps: 0, pos: [1, 1] }]];
  const visited: Map<string, number> = new Map([["1,1", 0]]);

  function dist(x: number, y: number): number {
    return Math.abs(x - xGoal) + Math.abs(y - yGoal);
  }

  while (queue.length) {
    let [_, { steps, pos }] = MinHeap.pop(queue) as [number, State];
    const [x, y] = pos;
    if (x === xGoal && y === yGoal) return steps;
    steps++;

    for (const [dx, dy] of directions) {
      const [x2, y2] = [x + dx, y + dy];
      if (x2 < 0 || y2 < 0) continue;
      if (!isOpenSpace(x2, y2, puzzleInput)) continue;
      const coord = `${x2},${y2}`;
      if (visited.has(coord) && visited.get(coord)! <= steps) continue;
      visited.set(coord, steps);
      MinHeap.push(queue, [dist(x2, y2) + steps, { steps, pos: [x2, y2] }]);
    }
  }

  return -1;
}

function part2(puzzleInput: number): number {
  const queue: State[] = [{ steps: 0, pos: [1, 1] }];
  const visited: Map<string, number> = new Map([["1,1", 0]]);

  while (queue.length) {
    let { steps, pos } = queue.pop()!;
    if (steps === 50) continue;
    const [x, y] = pos;
    steps++;
    for (const [dx, dy] of directions) {
      const [x2, y2] = [x + dx, y + dy];
      if (x2 < 0 || y2 < 0) continue;
      if (!isOpenSpace(x2, y2, puzzleInput)) continue;
      const coord = `${x2},${y2}`;
      if (visited.has(coord) && visited.get(coord)! <= steps) continue;
      visited.set(coord, steps);
      queue.push({ steps, pos: [x2, y2] });
    }
  }

  return visited.size;
}

console.log(part1(puzzleInput));
console.log(part2(puzzleInput));
