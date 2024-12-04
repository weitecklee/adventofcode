const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n");

const directions = [
  [1, 0],
  [0, 1],
  [-1, 0],
  [0, -1],
  [1, 1],
  [1, -1],
  [-1, 1],
  [-1, -1],
];

const magicWord = "XMAS";

let part1 = 0;

for (let i = 0; i < input.length; i++) {
  for (let j = 0; j < input[i].length; j++) {
    if (input[i][j] === magicWord[0]) {
      for (const [dx, dy] of directions) {
        let x = i + dx;
        let y = j + dy;
        let k = 1;
        while (
          x >= 0 &&
          x < input.length &&
          y >= 0 &&
          y < input[i].length &&
          k < magicWord.length &&
          input[x][y] === magicWord[k]
        ) {
          x += dx;
          y += dy;
          k++;
        }
        if (k === magicWord.length) {
          part1++;
        }
      }
    }
  }
}

console.log(part1);

let part2 = 0;

for (let i = 1; i < input.length - 1; i++) {
  for (let j = 1; j < input[i].length - 1; j++) {
    if (input[i][j] === "A") {
      const cornerLetters = [
        input[i - 1][j - 1],
        input[i - 1][j + 1],
        input[i + 1][j + 1],
        input[i + 1][j - 1],
      ].join("");
      if (
        cornerLetters === "MMSS" ||
        cornerLetters === "MSSM" ||
        cornerLetters === "SSMM" ||
        cornerLetters === "SMMS"
      ) {
        part2++;
      }
    }
  }
}

console.log(part2);
