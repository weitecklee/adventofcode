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

function* priceGenerator(n) {
  yield n % 10;
  while (true) {
    n = (n << 6) ^ n;
    n = n & 16777215;
    n = (n >> 5) ^ n;
    n = n & 16777215;
    n = (n << 11) ^ n;
    n = n & 16777215;
    yield n % 10;
  }
}

function price2000(n) {
  const gen = priceGenerator(n);
  const prices = [];
  for (let i = 0; i < 2000; i++) {
    prices.push(gen.next().value);
  }
  return prices;
}

const pricesColl = [];
for (const n of input) {
  pricesColl.push(price2000(n));
}

const sequencesMap = new Map();
for (const prices of pricesColl) {
  const sequences = new Set();
  const diffs = [];
  for (let i = 1; i < prices.length; i++) {
    diffs.push(prices[i] - prices[i - 1]);
  }
  for (let i = 3; i < diffs.length; i++) {
    const mapKey = diffs.slice(i - 3, i + 1).join(",");
    if (!sequences.has(mapKey)) {
      sequences.add(mapKey);
      sequencesMap.set(mapKey, (sequencesMap.get(mapKey) || 0) + prices[i + 1]);
    }
  }
}

console.log(Math.max(...Array.from(sequencesMap.values())));
