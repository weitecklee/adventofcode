const fs = require('fs');
const path = require('path');

let input = fs.readFileSync(path.resolve(__dirname, 'input202103.txt'), 'utf-8', (err, data) => {
  if (err) {
    console.log(err)
  } else {
    return data;
  }
});

input = input.split('\n');

const bits = new Array(input[0].length).fill(0);

for (const line of input) {
  for (let i = 0; i < line.length; i++) {
    if (line[i] === '1') {
      bits[i]++;
    }
  }
}

let gamma = '';
let epsilon = '';
for (const bit of bits) {
  if (bit > input.length / 2) {
    gamma += '1';
    epsilon += '0';
  } else {
    gamma += '0';
    epsilon += '1';
  }
}
gamma = Number.parseInt(gamma, 2);
epsilon = Number.parseInt(epsilon, 2);
console.log(gamma * epsilon);

let place = 0;
let candidates = input;
while (candidates.length > 1) {
  let bit = 0;
  const ones = [];
  const zeros = [];
  for (const line of candidates) {
    if (line[place] === '1') {
      bit++;
      ones.push(line);
    } else {
      zeros.push(line);
    }
  }
  if (bit >= candidates.length / 2) {
    candidates = ones;
  } else {
    candidates = zeros;
  }
  place++;
}

const oxygenRating = Number.parseInt(candidates[0], 2);

place = 0;
candidates = input;
while (candidates.length > 1) {
  let bit = 0;
  const ones = [];
  const zeros = [];
  for (const line of candidates) {
    if (line[place] === '1') {
      bit++;
      ones.push(line);
    } else {
      zeros.push(line);
    }
  }
  if (bit >= candidates.length / 2) {
    candidates = zeros;
  } else {
    candidates = ones;
  }
  place++;
}

const CO2Rating = Number.parseInt(candidates[0], 2);

console.log(oxygenRating * CO2Rating);
