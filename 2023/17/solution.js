const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split("").map(Number));

const queue = [[0, 0, -1, 0, 0]]; // x, y, directionIndex, steps in current direction, heat loss
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
  const [x, y, dirIndex, steps, heatLoss] = queue.pop();
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
    const nextOne = [newX, newY, i, newSteps, newHeatLoss];
    const state = stateKey(newX, newY, i, newSteps);
    if (visited.has(state) && visited.get(state) <= newHeatLoss) continue;
    visited.set(state, newHeatLoss);
    queue.push(nextOne);
  }
  queue.sort((a, b) => b[4] - a[4]);
}

const queue2 = [[0, 0, -1, 5, 0]]; // use starting steps of 5 to make it choose a new direction
const visited2 = new Map();

while (queue2.length) {
  const [x, y, dirIndex, steps, heatLoss] = queue2.pop();
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
    const nextOne = [newX, newY, dirIndex, newSteps, newHeatLoss];
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
      const nextOne = [newX, newY, i, newSteps, newHeatLoss];
      const state = stateKey(newX, newY, i, newSteps);
      if (visited2.has(state) && visited2.get(state) <= newHeatLoss) continue;
      visited2.set(state, newHeatLoss);
      queue2.push(nextOne);
    }
  }
  queue2.sort((a, b) => b[4] - a[4]);
}
