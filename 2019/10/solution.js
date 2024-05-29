const fs = require('fs');
const path = require('path');

const input = fs.readFileSync(path.join(__dirname, 'input.txt'), 'utf-8', (err, data) => {
  if (err) {
    console.log(err)
  } else {
    return data;
  }
}).split('\n');

const height = input.length;
const width = input[0].length;
const maxDim = Math.max(height, width);

const gcd = (a, b) => {
  while (b != 0) {
    const tmp = b;
    b = a % b;
    a = tmp;
  }
  return a;
}

const isCoprime = (a, b) => gcd(a, b) === 1;

const calculateAngle = (a, b) => {
  const theta = Math.atan2(-b, a);
  const angle = Math.PI / 2 - theta;
  return angle >= 0 ? angle : angle + 2 * Math.PI;
}

const angles = [];
angles.push([1, 0, calculateAngle(1, 0)]);
angles.push([-1, 0, calculateAngle(-1, 0)]);
angles.push([0, 1, calculateAngle(0, 1)]);
angles.push([0, -1, calculateAngle(0, -1)]);

for (let i = 1; i <= maxDim; i++) {
  for (let j = 1; j <= maxDim; j++) {
    if (isCoprime(i, j)) {
      angles.push([i, j, calculateAngle(i, j)]);
      angles.push([i, -j, calculateAngle(i, -j)]);
      angles.push([-i, j, calculateAngle(-i, j)]);
      angles.push([-i, -j, calculateAngle(-i, -j)]);
    }
  }
}

angles.sort((a, b) => a[2] - b[2]);

let part1 = 0;
let base = [];

for (let row = 0; row < height; row++) {
  for (let col = 0; col < width; col++) {
    if (input[row][col] != '#') {
      continue;
    }
    let count = 0;
    for (const [a, b, angle] of angles) {
      for (let i = 1; (a * i + col) >= 0 && (b * i + row) >= 0 && (a * i + col) < width && (b * i + row) < height; i++) {
        if (input[b * i + row][a * i + col] === '#') {
          count++;
          break;
        }
      }
    }
    if (count > part1) {
      base = [col, row];
      part1 = count;
    }
  }
}

console.log(part1);

const [col, row] = base;

const asteroids = [];

for (const [a, b, angle] of angles) {
  const currAngle = [];
  for (let i = 1; (a * i + col) >= 0 && (b * i + row) >= 0 && (a * i + col) < width && (b * i + row) < height; i++) {
    if (input[b * i + row][a * i + col] === '#') {
      currAngle.push([a * i + col, b * i + row]);
    }
  }
  asteroids.push(currAngle);
}

let i = 0;
let j = 0;

while (i <= 200) {
  for (const angle of asteroids) {
    if (j < angle.length) {
      i++;
      if (i === 200) {
        console.log(angle[j][0] * 100 + angle[j][1]);
        break;
      }
    }
  }
  j++;
}