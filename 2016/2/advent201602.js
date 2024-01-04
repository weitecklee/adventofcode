const fs = require('fs');

const input = fs.readFileSync('input.txt', 'utf-8', (err, data) => {
  if (err) {
    console.log(err)
  } else {
    return data;
  }
}).split('\n');

const pos = [1, 1];
let res = 0;

for (const line of input) {
  for (const c of line) {
    if (c === "U") {
      pos[1]--;
      if (pos[1] < 0) {
        pos[1] = 0
      }
    } else if (c === "D") {
      pos[1]++;
      if (pos[1] > 2) {
        pos[1] = 2;
      }
    } else if (c === "L") {
      pos[0]--;
      if (pos[0] < 0) {
        pos[0] = 0
      }

    } else {
      pos[0]++;
      if (pos[0] > 2) {
        pos[0] = 2;
      }
    }
  }
  res = res * 10 + (pos[1]) * 3 + pos[0] + 1;
}

console.log(res);

pos[0] = 0;
pos[1] = 2;
res = '';
const keypad = [
  '__1__',
  '_234_',
  '56789',
  '_ABC_',
  '__D__',
]

for (const line of input) {
  for (const c of line) {
    if (c === "U") {
      if ((pos[0] === 1 || pos[0] === 3) && pos[1] > 1) {
        pos[1]--;
      } else if (pos[0] === 2 && pos[1] > 0) {
        pos[1]--;
      }
    } else if (c === "D") {
      if ((pos[0] === 1 || pos[0] === 3) && pos[1] < 3) {
        pos[1]++;
      } else if (pos[0] === 2 && pos[1] < 4) {
        pos[1]++;
      }
    } else if (c === "L") {
      if ((pos[1] === 1 || pos[1] === 3) && pos[0] > 1) {
        pos[0]--;
      } else if (pos[1] === 2 && pos[0] > 0) {
        pos[0]--;
      }
    } else {
      if ((pos[1] === 1 || pos[1] === 3) && pos[0] < 3) {
        pos[0]++;
      } else if (pos[1] === 2 && pos[0] < 4) {
        pos[0]++;
      }
    }
  }
  res += keypad[pos[1]][pos[0]];
}

console.log(res);