const fs = require('fs');

const input = fs.readFileSync('input.txt', 'utf-8', (err, data) => {
  if (err) {
    console.log(err)
  } else {
    return data;
  }
}).split('\n');

const re1 = /[aeiou]/g;
const re2 = /(\w)\1/;
const re3 = /ab|cd|pq|xy/;
const re4 = /(\w\w).*\1/;
const re5 = /(\w)\w\1/;

let count = 0;
let count2 = 0;

for (const line of input) {
  if (re1.test(line) && line.match(re1).length >= 3 && re2.test(line) && !re3.test(line)) {
    count++;
  }
  if (re4.test(line) && re5.test(line)) {
    count2++;
  }
}

console.log(count);
console.log(count2);
