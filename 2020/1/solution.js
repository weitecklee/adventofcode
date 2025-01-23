const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8", (err, data) => {
    if (err) {
      console.log(err);
    } else {
      return data;
    }
  })
  .split("\n")
  .map(Number);

input.sort((a, b) => a - b);

const target = 2020;

let a = 0;
let b = input.length - 1;

while (a < b && input[a] + input[b] != target) {
  if (input[a] + input[b] > target) {
    b--;
  } else {
    a++;
  }
}

if (a >= b) {
  console.log("Could not find solution for part 1.");
  return;
}
console.log(input[a] * input[b]);

for (let i = 0; i < input.length - 2; i++) {
  let j = i + 1;
  let k = input.length - 1;
  while (j < k && input[i] + input[j] + input[k] != target) {
    if (input[i] + input[j] + input[k] > target) {
      k--;
    } else {
      j++;
    }
  }
  if (input[i] + input[j] + input[k] === target) {
    console.log(input[i] * input[j] * input[k]);
    return;
  }
}

console.log("Could not find solution for part 2.");
