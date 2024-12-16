const fs = require("fs");
const path = require("path");
const MinHeap = require("../../utils/MinHeap");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n");

const start = [0, 0];

const directions = [
  [0, 1],
  [1, 0],
  [-1, 0],
  [0, -1],
];

for (let r = 0; r < input.length; r++) {
  for (let c = 0; c < input[r].length; c++) {
    if (input[r][c] === "S") {
      start[0] = r;
      start[1] = c;
      break;
    }
  }
}

const queue = [[0, ...start, 0, new Set([`${start[0]},${start[1]}`])]];
const visitedScores = new Map([[`${start[0]},${start[1]},0`, 0]]);
let part2 = new Set();
let minScore = Infinity;

while (queue.length) {
  const [score, r, c, dirIndex, visited] = MinHeap.pop(queue);
  if (input[r][c] === "E") {
    if (score <= minScore) {
      minScore = score;
      part2 = part2.union(visited);
    } else {
      break;
    }
    continue;
  }
  for (let i = 0; i < directions.length; i++) {
    if (i === 3 - dirIndex) continue;
    const [dr, dc] = directions[i];
    const [newR, newC] = [r + dr, c + dc];
    if (
      newR < 0 ||
      newR >= input.length ||
      newC < 0 ||
      newC >= input[0].length ||
      input[newR][newC] === "#"
    )
      continue;
    const newScore = score + 1 + (i !== dirIndex ? 1000 : 0);
    if (
      visitedScores.has(`${newR},${newC},${i}`) &&
      visitedScores.get(`${newR},${newC},${i}`) < newScore
    )
      continue;
    visitedScores.set(`${newR},${newC},${i}`, newScore);
    const newVisited = new Set(visited);
    newVisited.add(`${newR},${newC}`);
    const nextOne = [newScore, newR, newC, i, newVisited];
    MinHeap.push(queue, nextOne);
  }
}

console.log(minScore);
console.log(part2.size);
