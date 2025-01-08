const fs = require("fs");
const path = require("path");
const MinHeap = require("../../utils/MinHeap");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split("").map(Number));

function findMinRisk(input) {
  const turnDirections = [
    [0, 1],
    [1, 0],
    [-1, 0],
    [0, -1],
  ];

  const rMax = input.length - 1;
  const cMax = input[0].length - 1;

  const queue = [[0, 0, 0]];
  const visited = new Map();

  while (queue.length) {
    const [risk, r, c] = MinHeap.pop(queue);
    if (r === rMax && c === cMax) {
      return risk;
    }
    for (const [dr, dc] of turnDirections) {
      const r2 = r + dr;
      const c2 = c + dc;
      if (r2 < 0 || c2 < 0 || r2 > rMax || c2 > cMax) continue;
      const risk2 = risk + input[r2][c2];
      if (visited.has(`${r2},${c2}`) && visited.get(`${r2},${c2}`) <= risk2)
        continue;
      visited.set(`${r2},${c2}`, risk2);
      MinHeap.push(queue, [risk + input[r2][c2], r2, c2]);
    }
  }
}

function expandInput(input) {
  const rows = input.length;
  const cols = input[0].length;
  for (let i = 0; i < 4; i++) {
    for (let r = rows * i; r < rows * (i + 1); r++) {
      const newRow = [];
      for (let c = 0; c < cols; c++) {
        let val = input[r][c] + 1;
        if (val > 9) val = 1;
        newRow.push(val);
      }
      input.push(newRow);
    }
  }
  for (let i = 0; i < 4; i++) {
    for (let r = 0; r < input.length; r++) {
      for (let c = cols * i; c < cols * (i + 1); c++) {
        let val = input[r][c] + 1;
        if (val > 9) val = 1;
        input[r].push(val);
      }
    }
  }

  return input;
}

console.log(findMinRisk(input));
console.log(findMinRisk(expandInput(input)));
