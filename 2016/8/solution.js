const fs = require('fs');
const path = require('path');

const input = fs.readFileSync(path.join(__dirname, 'input.txt'), 'utf-8', (err, data) => {
  if (err) {
    console.log(err)
  } else {
    return data;
  }
}).split('\n');

const screen = new Array(6).fill(null).map((a) => new Array(50).fill(' '));

const re = /\d+/g;

for (const line of input) {
  const nums = line.match(re).map(Number);
  if (line.includes('rect')) {
    for (let j = 0; j < nums[0]; j++) {
      for (let i = 0; i < nums[1]; i++) {
        screen[i][j] = '#';
      }
    }
  } else if (line.includes('row')) {
    screen[nums[0]] = screen[nums[0]].slice(50 - nums[1]).concat(screen[nums[0]].slice(0, 50 - nums[1]))
  } else {
    const column = screen.map((a) => a[nums[0]]);
    for (let i = nums[1]; i < 6; i++) {
      screen[i][nums[0]] = column[i - nums[1]];
    }
    for (let i = 0; i < nums[1]; i++) {
      screen[i][nums[0]] = column[i + 6 - nums[1]];
    }
  }
}

let count = 0;
for (const row of screen) {
  count += row.reduce((a, b) => a + (b === '#' ? 1 : 0), 0);
  console.log(row.join(''));
}

console.log(count);