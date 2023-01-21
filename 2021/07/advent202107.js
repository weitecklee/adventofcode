const fs = require('fs');
const path = require('path');

let input = fs.readFileSync(path.resolve(__dirname, 'input202107.txt'), 'utf-8', (err, data) => {
  if (err) {
    console.log(err)
  } else {
    return data;
  }
});

input = input.split(',').map(Number);

const steps = (align) => {
  let sum = 0;
  for (const num of input) {
    sum += Math.abs(num - align);
  }
  return sum;
}

let left = Math.min(...input);
let right = Math.max(...input);
const candidates = [[left, steps(left)], [right, steps(right)]];

while (Math.abs(candidates[0][0] - candidates[1][0]) > 1) {
  const mid = Math.floor((candidates[0][0] + candidates[1][0]) / 2);
  candidates.push([mid, steps(mid)]);
  candidates.sort((a, b) => a[1] - b[1]);
}
console.log(candidates[0][1]);

const steps2 = (align) => {
  let sum = 0;
  for (const num of input) {
    const d = Math.abs(num - align);
    sum += d * (d + 1) / 2;
  }
  return sum;
}

const candidates2 = [[left, steps2(left)], [right, steps2(right)]];

while (Math.abs(candidates2[0][0] - candidates2[1][0]) > 1) {
  const mid = Math.floor((candidates2[0][0] + candidates2[1][0]) / 2);
  candidates2.push([mid, steps2(mid)]);
  candidates2.sort((a, b) => a[1] - b[1]);
}
console.log(candidates2[0][1]);
