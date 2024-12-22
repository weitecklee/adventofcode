const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map(Number);

function* secretGenerator(n) {
  while (true) {
    n = (n << 6) ^ n;
    n = n & 16777215;
    n = (n >> 5) ^ n;
    n = n & 16777215;
    n = (n << 11) ^ n;
    n = n & 16777215;
    yield n;

    // n = (n * 64) ^ n;
    // n %= 16777216;
    // n = Math.trunc(n / 32) ^ n;
    // n %= 16777216;
    // n = (n * 2048) ^ n;
    // n %= 16777216;
  }
}

function secret2000(n) {
  const gen = secretGenerator(n);
  for (let i = 1; i < 2000; i++) {
    gen.next();
  }
  return gen.next().value;
}

let part1 = 0;
for (const n of input) {
  part1 += secret2000(n);
}
console.log(part1);
