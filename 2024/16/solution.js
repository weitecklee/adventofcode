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

const queue = [[0, ...start, 0]];
const visitedScores = new Map([[`${start[0]},${start[1]}`, 0]]);

while (queue.length) {
  const [score, r, c, dirIndex] = MinHeap.pop(queue);
  if (input[r][c] === "E") {
    console.log(score);
    break;
  }
  for (let i = 0; i < directions.length; i++) {
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
    const newScore =
      score + 1 + (i === 3 - dirIndex ? 2000 : i !== dirIndex ? 1000 : 0);
    if (
      visitedScores.has(`${newR},${newC}`) &&
      visitedScores.get(`${newR},${newC}`) <= newScore
    )
      continue;
    visitedScores.set(`${newR},${newC}`, newScore);
    const nextOne = [newScore, newR, newC, i];
    MinHeap.push(queue, nextOne);
  }
}
