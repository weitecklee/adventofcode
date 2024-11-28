const fs = require("fs");
const path = require("path");
const MinHeap = require("../../utils/MinHeap");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split("").map(Number));

const queue = [[0, 0, 0, -1, 0]]; // heat loss, x, y, directionIndex, steps in current direction
const turnDirections = [
  [0, 1],
  [1, 0],
  [-1, 0],
  [0, -1],
];
const visited = new Map();
// Map[statekey] = heatLoss
function stateKey(x, y, dir, steps) {
  return `${x},${y},${dir},${steps}`;
}

while (queue.length) {
  const [heatLoss, x, y, dirIndex, steps] = MinHeap.pop(queue);
  if (x === input[0].length - 1 && y === input.length - 1) {
    console.log(heatLoss);
    break;
  }
  for (let i = 0; i < turnDirections.length; i++) {
    if (i === 3 - dirIndex) continue; // no u-turns)
    if (i === dirIndex && steps === 3) continue; // must turn after 3 steps
    const [dx, dy] = turnDirections[i];
    const [newX, newY] = [x + dx, y + dy];
    if (newX < 0 || newX >= input[0].length || newY < 0 || newY >= input.length)
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

const queue2 = [[0, 0, 0, -1, 5]]; // use starting steps of 5 to make it choose a new direction
const visited2 = new Map();

while (queue2.length) {
  const [heatLoss, x, y, dirIndex, steps] = MinHeap.pop(queue2);
  if (x === input[0].length - 1 && y === input.length - 1) {
    console.log(heatLoss);
    break;
  }
  if (steps < 4) {
    // must stay in the same direction for at least 4 steps
    const [dx, dy] = turnDirections[dirIndex];
    const [newX, newY] = [x + dx, y + dy];
    if (newX < 0 || newX >= input[0].length || newY < 0 || newY >= input.length)
      continue;
    const newSteps = steps + 1;
    const newHeatLoss = heatLoss + input[newY][newX];
    const nextOne = [newHeatLoss, newX, newY, dirIndex, newSteps];
    const state = stateKey(newX, newY, dirIndex, newSteps);
    if (visited2.has(state) && visited2.get(state) <= newHeatLoss) continue;
    visited2.set(state, newHeatLoss);
    queue2.push(nextOne);
  } else {
    for (let i = 0; i < turnDirections.length; i++) {
      if (i === 3 - dirIndex) continue; // no u-turns)
      if (i === dirIndex && steps === 10) continue; // must turn after 10 steps
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
      if (visited2.has(state) && visited2.get(state) <= newHeatLoss) continue;
      visited2.set(state, newHeatLoss);
      MinHeap.push(queue2, nextOne);
    }
  }
}

/*
  Old solution:
  part1: 6.462s
  part2: 51.564s

  New solution with minHeap:
  part1: 298.899ms
  part2: 752.789ms
*/
