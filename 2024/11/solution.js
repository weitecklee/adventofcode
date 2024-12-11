const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split(" ")
  .map(Number);

function blink(n) {
  if (n === 0) {
    return [1];
  }
  const nString = n.toString();
  if (nString.length % 2 === 0) {
    return [
      Number(nString.slice(0, nString.length / 2)),
      Number(nString.slice(nString.length / 2)),
    ];
  }
  return [n * 2024];
}

class CustomMap extends Map {
  constructor(arr = []) {
    super();
    this.add(arr, 1);
  }

  add(arr, k) {
    for (const n of arr) {
      this.set(n, this.has(n) ? this.get(n) + k : k);
    }
  }

  get count() {
    return Array.from(this.values()).reduce((a, b) => a + b, 0);
  }
}

let curr = new CustomMap(input);

for (let i = 0; i < 25; i++) {
  const temp = new CustomMap();
  for (const [n, k] of curr) {
    temp.add(blink(n), k);
  }
  curr = temp;
}

console.log(curr.count);

for (let i = 25; i < 75; i++) {
  const temp = new CustomMap();
  for (const [n, k] of curr) {
    temp.add(blink(n), k);
  }
  curr = temp;
}

console.log(curr.count);
