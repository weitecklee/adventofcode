const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split(" ")
  .map(Number);

const memo = new Map();

function blink(n) {
  if (memo.has(n)) {
    return memo.get(n);
  }
  if (n === 0) {
    memo.set(n, [1]);
    return [1];
  }
  const nString = n.toString();
  if (nString.length % 2 === 0) {
    memo.set(n, [
      Number(nString.slice(0, nString.length / 2)),
      Number(nString.slice(nString.length / 2)),
    ]);
    return [
      Number(nString.slice(0, nString.length / 2)),
      Number(nString.slice(nString.length / 2)),
    ];
  }
  memo.set(n, [n * 2024]);
  return [n * 2024];
}

class CustomMap {
  constructor(arr = []) {
    this.map = new Map();
    this.add(arr, 1);
  }
  add(arr, k) {
    for (const n of arr) {
      this.map.set(n, this.map.has(n) ? this.map.get(n) + k : k);
    }
  }
  get count() {
    return this.map.values().reduce((a, b) => a + b, 0);
  }
}

let curr = new CustomMap(input);
for (let i = 0; i < 75; i++) {
  const temp = new CustomMap();
  for (const [n, k] of curr.map) {
    const res = blink(n);
    temp.add(res, k);
  }
  curr = temp;
}

console.log(curr.count);
