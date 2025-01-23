const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8", (err, data) => {
    if (err) {
      console.log(err);
    } else {
      return data;
    }
  })
  .split("\n");

function Card(winningNumbers, myNumbers) {
  this.winningNumbers = new Set(winningNumbers);
  this.myNumbers = myNumbers;
  this.matches = 0;
  for (const num of this.myNumbers) {
    this.matches += this.winningNumbers.has(num);
  }
  this.copies = 1;
}

const numberPattern = /\d+/g;
const cards = [];

for (const line of input) {
  const numbers = line.split(":")[1].split("|");
  const winningNumbers = numbers[0].match(numberPattern).map(Number);
  const myNumbers = numbers[1].match(numberPattern).map(Number);
  cards.push(new Card(winningNumbers, myNumbers));
}

const part1 = cards.reduce(
  (a, b) => a + (b.matches ? 2 ** (b.matches - 1) : 0),
  0
);
console.log(part1);

for (let i = 0; i < cards.length; i++) {
  for (let j = 1; j <= cards[i].matches; j++) {
    cards[i + j].copies += cards[i].copies;
  }
}

const part2 = cards.reduce((a, b) => a + b.copies, 0);
console.log(part2);
