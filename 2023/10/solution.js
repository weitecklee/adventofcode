const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split(""));

const start = [0, 0];
let startFound = false;

for (let r = 0; r < input.length; r++) {
  for (let c = 0; c < input[0].length; c++) {
    if (input[r][c] === "S") {
      [start[0], start[1]] = [c, r];
      break;
    }
  }
  if (startFound) break;
}

const directions = [
  [0, 1], // down
  [1, 0], // right
  [-1, 0], // left
  [0, -1], // up
];

const queue = [];
// figure out valid directions from S
for (let i = 0; i < directions.length; i++) {
  const [dx, dy] = directions[i];
  let x = start[0] + dx;
  let y = start[1] + dy;
  if (x < 0 || x >= input[0].length || y < 0 || y >= input.length) continue;
  let turn = i;
  if (i === 0) {
    // going down
    switch (input[y][x]) {
      case "L": // turn right
        turn = 1;
        break;
      case "J": // turn left
        turn = 2;
        break;
      case "|": // continue down
        break;
      default:
        continue;
    }
  } else if (i === 1) {
    // going right
    switch (input[y][x]) {
      case "J": // turn up
        turn = 3;
        break;
      case "7": // turn down
        turn = 0;
        break;
      case "-": // continue right
        break;
      default:
        continue;
    }
  } else if (i === 2) {
    // going left
    switch (input[y][x]) {
      case "L": // turn up
        turn = 3;
        break;
      case "F": // turn down
        turn = 0;
        break;
      case "-": // continue left
        break;
      default:
        continue;
    }
  } else {
    // going up
    switch (input[y][x]) {
      case "F": // turn right
        turn = 1;
        break;
      case "7": // turn left
        turn = 2;
        break;
      case "|": // continue up
        break;
      default:
        continue;
    }
  }
  queue.push([x, y, turn]);
}

// figure out what pipe S is
const dirFromStart = [queue.pop()[2], queue.pop()[2]];
dirFromStart.sort((a, b) => a - b);
if (dirFromStart[0] === 0) {
  if (dirFromStart[1] === 1) {
    // down and right
    input[start[1]][start[0]] = "F";
  } else if (dirFromStart[1] === 2) {
    // down and left
    input[start[1]][start[0]] = "7";
  } else {
    // down and up
    input[start[1]][start[0]] = "|";
  }
} else if (dirFromStart[0] === 1) {
  if (dirFromStart[1] === 2) {
    // right and left
    input[start[1]][start[0]] = "-";
  } else {
    // right and up
    input[start[1]][start[0]] = "L";
  }
} else {
  // left and up
  input[start[1]][start[0]] = "J";
}

let [x, y] = start;
let dir = dirFromStart[0];
const pipePath = [[...start, input[start[1]][start[0]]]];
while (true) {
  const [dx, dy] = directions[dir];
  x += dx;
  y += dy;
  if (x === start[0] && y === start[1]) break;

  switch (input[y][x]) {
    case "F":
      dir = dir === 3 ? 1 : 0;
      break;
    case "7":
      dir = dir === 3 ? 2 : 0;
      break;
    case "L":
      dir = dir === 0 ? 1 : 3;
      break;
    case "J":
      dir = dir === 0 ? 2 : 3;
      break;
    case "|":
    case "-":
      break;
    default:
      throw new Error("Invalid pipe: " + input[y][x]);
  }
  pipePath.push([x, y, input[y][x]]);
}

console.log(pipePath.length / 2);

// Combination of Shoelace formula and Pick's theorem
// Shoelace formula calculates area of polygon given vertex coordinates
// Pick's theorem calculates area of polygon given number of boundary points and number of interior points
// We have vertices and boundary points in pipePath, and ultimately want to find sum of interior points and boundary points
// Shoelace formula
// https://en.wikipedia.org/wiki/Shoelace_formula
// Pick's theorem
// https://en.wikipedia.org/wiki/Pick%27s_theorem

const boundaryPoints = pipePath.length;
const vertices = pipePath.filter((a) => a[2] !== "|" && a[2] !== "-");

let sum = 0;
for (let i = 0; i < vertices.length; i++) {
  const [x1, y1] = vertices[i];
  const [x2, y2] = vertices[(i + 1) % vertices.length];
  sum += x1 * y2 - x2 * y1;
}

const area = Math.abs(sum) / 2; // from Shoelace formula
const interiorPoints = area - boundaryPoints / 2 + 1; // Pick's theorem rearranged
console.log(interiorPoints);
