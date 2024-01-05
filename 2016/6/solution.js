const fs = require('fs');
const path = require('path');

const input = fs.readFileSync(path.join(__dirname, 'input.txt'), 'utf-8', (err, data) => {
  if (err) {
    console.log(err)
  } else {
    return data;
  }
}).split('\n');

const frequencyMaps = [];
for (let i = 0; i < input[0].length; i++) {
  frequencyMaps.push(new Map());
}

for (const line of input) {
  for (let i = 0; i < line.length; i++) {
    if (!frequencyMaps[i].has(line[i])) {
      frequencyMaps[i].set(line[i], 0);
    }
    frequencyMaps[i].set(line[i], frequencyMaps[i].get(line[i]) + 1);
  }
}

let res = '';
let res2 = '';

for (const freqMap of frequencyMaps) {
  const freqArray = [];
  for (const [letter, n] of freqMap) {
    freqArray.push([letter, n]);
  }
  freqArray.sort((a, b) => b[1] - a[1]);
  res += freqArray[0][0];
  res2 += freqArray[freqArray.length - 1][0];
}

console.log(res);
console.log(res2);