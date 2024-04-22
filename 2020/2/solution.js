const fs = require('fs');
const path = require('path');

const input = fs.readFileSync(path.join(__dirname, 'input.txt'), 'utf-8', (err, data) => {
  if (err) {
    console.log(err)
  } else {
    return data;
  }
}).split('\n');

let part1 = 0;
let part2 = 0;

const linePattern = /^(\d+)-(\d+) (\w): (\w+)$/;
for (const line of input) {
  const lineMatch = line.match(linePattern);
  const a = Number(lineMatch[1]);
  const b = Number(lineMatch[2]);
  const letter = lineMatch[3];
  const password = lineMatch[4];
  const count = password.split("").reduce((a, b) => a + (b === letter), 0);
  if (count >= a && count <= b) {
    part1++;
  }
  if ((password[a - 1] === letter) ^ (password[b - 1] === letter) ) {
    part2++;
  }
}

console.log(part1);
console.log(part2);
