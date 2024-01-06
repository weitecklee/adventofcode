const fs = require('fs');
const path = require('path');

const input = fs.readFileSync(path.join(__dirname, 'input.txt'), 'utf-8', (err, data) => {
  if (err) {
    console.log(err)
  } else {
    return data;
  }
}).split('\n');

const gearMap = new Map();
const symbolSet = new Set();
const symbolPattern = /[^0-9.]/g;
const numberPattern = /\d+/g;
const gearPattern = /\*/g;
const coord = (a, b) => a + ',' + b;

function Gear() {
  this.ratio = 1;
  this.parts = 0;
}

for (let i = 0; i < input.length; i++) {
  const matches = input[i].matchAll(symbolPattern);
  for (const match of matches) {
    symbolSet.add(coord(match.index, i));
  }
  const gears = input[i].matchAll(gearPattern);
  for (const gear of gears) {
    gearMap.set(coord(gear.index, i), new Gear());
  }
}

let part1 = 0;

for (let i = 0; i < input.length; i++) {
  const matches = input[i].matchAll(numberPattern);
  for (const match of matches) {
    let isPart = false;
    for (let j = -1; j < match[0].length + 1; j++) {
      for (let k = -1; k < 2; k++) {
      const curr = coord(match.index + j, i + k);
      if (symbolSet.has(curr)) {
        isPart = true;
      }
      if (gearMap.has(curr)) {
        const gear = gearMap.get(curr);
        gear.parts++;
        gear.ratio *= Number(match[0]);
      }
      }
    }
    if (isPart) {
      part1 += Number(match[0]);
    }
  }
}

console.log(part1);

let part2 = 0;

for (const gear of gearMap.values()) {
  if (gear.parts === 2) {
    part2 += gear.ratio;
  }
}

console.log(part2);