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

const red = 12;
const green = 13;
const blue = 14;

function Game(record) {
  const parts = record.split(": ");
  this.id = Number(parts[0].split(" ")[1]);
  this.subsets = [];
  for (const subsetString of parts[1].split("; ")) {
    const subset = new Map();
    for (const cubes of subsetString.split(", ")) {
      const parts2 = cubes.split(" ");
      subset.set(parts2[1], Number(parts2[0]));
    }
    this.subsets.push(subset);
  }
}

const games = [];

for (const line of input) {
  games.push(new Game(line));
}

let part1 = 0;
let part2 = 0;

for (const game of games) {
  let isPossible = true;
  let maxRed = 0;
  let maxGreen = 0;
  let maxBlue = 0;
  for (const subset of game.subsets) {
    if (
      isPossible &&
      ((subset.get("red") ?? 0) > red ||
        (subset.get("green") ?? 0) > green ||
        (subset.get("blue") ?? 0) > blue)
    ) {
      isPossible = false;
    }
    maxRed = Math.max(maxRed, subset.get("red") ?? 0);
    maxGreen = Math.max(maxGreen, subset.get("green") ?? 0);
    maxBlue = Math.max(maxBlue, subset.get("blue") ?? 0);
  }
  if (isPossible) {
    part1 += game.id;
  }
  part2 += maxRed * maxGreen * maxBlue;
}

console.log(part1);
console.log(part2);
