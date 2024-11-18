const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split(",")
  .map(Number);

// const gameHistory = new Map();

// input.forEach((num, i) => gameHistory.set(num, i + 1));

// let turn = input.length + 1;
// let currentNumber = 0;

// while (turn < 2020) {
//   let nextNumber = 0;
//   if (gameHistory.has(currentNumber)) {
//     nextNumber = turn - gameHistory.get(currentNumber);
//   }
//   gameHistory.set(currentNumber, turn);
//   currentNumber = nextNumber;
//   turn++;
// }

// console.log(currentNumber);

// while (turn < 30000000) {
//   let nextNumber = 0;
//   if (gameHistory.has(currentNumber)) {
//     nextNumber = turn - gameHistory.get(currentNumber);
//   }
//   gameHistory.set(currentNumber, turn);
//   currentNumber = nextNumber;
//   turn++;
// }

// console.log(currentNumber);

function playMemoryGame(input, turns) {
  const gameHistory = new Map();

  input.forEach((num, i) => gameHistory.set(num, i + 1));

  let turn = input.length + 1;
  let currentNumber = 0;

  while (turn < turns) {
    let nextNumber = 0;
    if (gameHistory.has(currentNumber)) {
      nextNumber = turn - gameHistory.get(currentNumber);
    }
    gameHistory.set(currentNumber, turn);
    currentNumber = nextNumber;
    turn++;
  }
  return currentNumber;
}

console.log(playMemoryGame(input, 2020));
console.log(playMemoryGame(input, 30000000));
