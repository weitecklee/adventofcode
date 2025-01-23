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

let count = 0;
const parsedInput = [];

const checkTriangle = (arr) =>
  arr[0] + arr[1] > arr[2] &&
  arr[0] + arr[2] > arr[1] &&
  arr[1] + arr[2] > arr[0];

for (const line of input) {
  const matches = line.match(/\d+/g);
  const sides = [];
  for (const match of matches) {
    sides.push(Number(match));
  }
  parsedInput.push(sides);
  if (checkTriangle(sides)) {
    count++;
  }
}

console.log(count);

let count2 = 0;

for (let i = 0; i < parsedInput.length; i += 3) {
  for (let j = 0; j < 3; j++) {
    const sides = [
      parsedInput[i][j],
      parsedInput[i + 1][j],
      parsedInput[i + 2][j],
    ];
    if (checkTriangle(sides)) {
      count2++;
    }
  }
}

console.log(count2);
