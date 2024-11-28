const fs = require("fs");
const path = require("path");
const MinHeap = require("../../utils/MinHeap");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split("").map(Number));

const turnDirections = [
  [0, 1],
  [1, 0],
  [-1, 0],
  [0, -1],
];

function stateKey(x, y, dir, steps) {
  return `${x},${y},${dir},${steps}`;
}

function findMinPath(queue, visited, minSteps, maxSteps, input) {
  // queue: [][heatLoss, x, y, dirIndex, steps]
  // visited: Map[stateKey] = heatLoss
  // minSteps: minimum number of steps to stay in the same direction
  // maxSteps: maximum number of steps before turning
  // input: grid

  while (queue.length) {
    const [heatLoss, x, y, dirIndex, steps] = MinHeap.pop(queue);
    if (x === input[0].length - 1 && y === input.length - 1) {
      return heatLoss;
    }
    for (let i = 0; i < turnDirections.length; i++) {
      if (steps < minSteps && i != dirIndex) continue; // must stay in the same direction for at least minSteps
      if (i === 3 - dirIndex) continue; // no u-turns)
      if (i === dirIndex && steps === maxSteps) continue; // must turn after maxSteps
      const [dx, dy] = turnDirections[i];
      const [newX, newY] = [x + dx, y + dy];
      if (
        newX < 0 ||
        newX >= input[0].length ||
        newY < 0 ||
        newY >= input.length
      )
        continue;
      const newSteps = i === dirIndex ? steps + 1 : 1;
      const newHeatLoss = heatLoss + input[newY][newX];
      const nextOne = [newHeatLoss, newX, newY, i, newSteps];
      const state = stateKey(newX, newY, i, newSteps);
      if (visited.has(state) && visited.get(state) <= newHeatLoss) continue;
      visited.set(state, newHeatLoss);
      MinHeap.push(queue, nextOne);
    }
  }
  return -1;
}

const part1 = findMinPath([[0, 0, 0, -1, 0]], new Map(), 0, 3, input);
console.log(part1);
const part2 = findMinPath([[0, 0, 0, -1, 5]], new Map(), 4, 10, input);
// use starting steps of 5 to make it choose a new direction
console.log(part2);

/*
  Old solution:
  part1: 6.462s
  part2: 51.564s

  New solution with minHeap:
  part1: 230.407ms
  part2: 744.543ms
*/
