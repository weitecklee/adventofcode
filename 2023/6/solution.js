const fs = require('fs');
const path = require('path');

const input = fs.readFileSync(path.join(__dirname, 'input.txt'), 'utf-8', (err, data) => {
  if (err) {
    console.log(err)
  } else {
    return data;
  }
}).split('\n');

const numberPattern = /\d+/g;

const times = input[0].match(numberPattern).map(Number);
const distances = input[1].match(numberPattern).map(Number);

const numberOfWays = (time, distance) => {
  let j = 1;
  while (j < time && j * (time - j) <= distance) {
    j++;
  }
  let k = time - 1;
  while (k > 0 && k * (time - k) <= distance) {
    k--;
  }
  return k >= j ? k - j + 1 : 0;
}

let part1 = 1;

for (let i = 0; i < times.length; i++) {
  part1 *= numberOfWays(times[i], distances[i]);
}

console.log(part1);

const time2 = Number(input[0].match(numberPattern).reduce((a, b) => a + b));
const distance2 = Number(input[1].match(numberPattern).reduce((a, b) => a + b));

const part2 = numberOfWays(time2, distance2);
console.log(part2);
