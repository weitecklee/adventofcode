const fs = require('fs');

let input = fs.readFileSync('input.txt', 'utf-8', (err, data) => {
  if (err) {
    console.log(err)
  } else {
    return data;
  }
});

input = input.split('\n');
input = input.map((a) => a.split('').map(Number));

const tracker = [];
for (const line of input) {
  tracker.push(new Array(line.length).fill(0));
}

for (let i = 0; i < input.length; i++) {
  let max = -1;
  for (let j = 0; j < input[i].length; j++) {
    if (input[i][j] > max) {
      tracker[i][j] = 1;
      max = input[i][j];
    }
  }
  max = -1;
  for (let j = input[i].length - 1; j >=0 ; j--) {
    if (input[i][j] > max) {
      tracker[i][j] = 1;
      max = input[i][j];
    }
  }
}

for (let j = 0; j < input[0].length; j++) {
  let max = -1;
  for (let i = 0; i < input.length; i++) {
    if (input[i][j] > max) {
      tracker[i][j] = 1;
      max = input[i][j];
    }
  }
  max = -1;
  for (let i = input.length - 1; i >= 0; i--) {
    if (input[i][j] > max) {
      tracker[i][j] = 1;
      max = input[i][j];
    }
  }
}

let count = 0;
for (let i = 0; i < tracker.length; i++) {
  for (let j = 0; j < tracker[i].length; j++) {
    if (tracker[i][j] > 0) {
      count++;
    }
  }
}

console.log(count);

let max = 0;

const viewer = (a, b) => {
  if (a === 0 || b === 0 || a === input.length - 1 || b === input[0].length - 1) {
    return 0;
  }
  const ht = input[a][b];

  let left = 0;
  for (x = a - 1; x >= 0; x--) {
    left++;
    if (input[x][b] >= ht) {
      break;
    }
  }
  let right = 0;
  for (x = a + 1; x < input.length; x++) {
    right++;
    if (input[x][b] >= ht) {
      break;
    }
  }

  let up = 0;
  for (y = b - 1; y >= 0; y--) {
    up++;
    if (input[a][y] >= ht) {
      break;
    }
  }
  let down = 0;
  for (y = b + 1; y < input[0].length; y++) {
    down++;
    if (input[a][y] >= ht) {
      break;
    }
  }
  return left * right * up * down;
}

for (let i = 0; i < input.length; i++) {
  for (let j = 0; j < input[i].length; j++) {
    const score = viewer(i, j);
    max = Math.max(score, max);
  }
}

console.log(max);