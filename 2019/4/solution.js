const fs = require('fs');
const path = require('path');

const input = fs.readFileSync(path.join(__dirname, 'input.txt'), 'utf-8', (err, data) => {
  if (err) {
    console.log(err)
  } else {
    return data;
  }
}).split('-').map(Number);

const adjacentSameDigitsRegex = /(.)\1+/g;

class Password {
  constructor(password) {
    this.password = password.toString();
    this.part1 = true;
    this.part2 = true;
    this.check();
  }

  check() {
    if (this.password.length != 6) {
      this.part1 = false;
      this.part2 = false;
      return;
    }
    const passwordArray = this.password.split('').map(Number);
    for (let i = 1; i < passwordArray.length; i++) {
      if (passwordArray[i] < passwordArray[i - 1]) {
        this.part1 = false;
        this.part2 = false;
        return;
      }
    }
    const matches = this.password.match(adjacentSameDigitsRegex);
    if (matches === null) {
      this.part1 = false;
      this.part2 = false;
      return;
    }
    this.part2 = matches.some((a) => a.length === 2);
  }
}

let part1 = 0;
let part2 = 0;
for (let i = input[0]; i <= input[1]; i++) {
  const password = new Password(i);
  if (password.part1) {
    part1++;
  }
  if (password.part2) {
    part2++;
  }
}

console.log(part1);
console.log(part2);

