const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map(Number);

function part1(input) {
  const preamble = new Set(input.slice(0, 25));

  function checkForSum(n, preamble) {
    for (const num of preamble) {
      if (preamble.has(n - num)) {
        return true;
      }
    }
    return false;
  }

  if (!checkForSum(input[25], preamble)) {
    return input[25];
  }

  for (let i = 26; i < input.length; i++) {
    preamble.delete(input[i - 26]);
    preamble.add(input[i - 1]);
    if (!checkForSum(input[i], preamble)) {
      return input[i];
    }
  }
}

const target = part1(input);
console.log(target);

let a = 0;
let b = 0;
let sum = input[0];
while (sum !== target) {
  if (sum < target) {
    // b++;
    // sum += input[b];
    sum += input[++b];
  } else {
    // sum -= input[a];
    // a++;
    sum -= input[a++];
  }
}

const rng = input.slice(a, b + 1);
console.log(Math.min(...rng) + Math.max(...rng));
