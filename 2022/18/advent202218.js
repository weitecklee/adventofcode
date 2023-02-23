const fs = require('fs');

let input = fs.readFileSync('input202218.txt', 'utf-8', (err, data) => {
  if (err) {
    console.log(err)
  } else {
    return data;
  }
});

input = input.split('\n');

const lines = new Set(input);

let area = 0;

const q = input.slice();

while (q.length) {
  const check = q.pop();
  const parse = check.split(',').map(Number);
  const toCheck = [];
  toCheck.push([parse[0] - 1, parse[1], parse[2]].join(','));
  toCheck.push([parse[0] + 1, parse[1], parse[2]].join(','));
  toCheck.push([parse[0], parse[1] - 1, parse[2]].join(','));
  toCheck.push([parse[0], parse[1] + 1, parse[2]].join(','));
  toCheck.push([parse[0], parse[1], parse[2] - 1].join(','));
  toCheck.push([parse[0], parse[1], parse[2] + 1].join(','));
  for (const entry of toCheck) {
    if (!lines.has(entry)) {
      area++;
    }
  }
}

console.log(area);

