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

// figure out what pipe S is
const validDirections = [];
for (let i = 0; i < directions.length; i++) {
  const [dx, dy] = directions[i];
  let x = start[0] + dx;
  let y = start[1] + dy;
  if (x < 0 || x >= input[0].length || y < 0 || y >= input.length) continue;
  if (i === 0) {
    switch (input[y][x]) {
      case "L":
      case "J":
      case "|":
        validDirections.push(i);
        break;
      default:
        continue;
    }
  } else if (i === 1) {
    switch (input[y][x]) {
      case "J":
      case "7":
      case "-":
        validDirections.push(i);
        break;
      default:
        continue;
    }
  } else if (i === 2) {
    switch (input[y][x]) {
      case "L":
      case "F":
      case "-":
        validDirections.push(i);
        break;
      default:
        continue;
    }
  } else {
    switch (input[y][x]) {
      case "F":
      case "7":
      case "|":
        validDirections.push(i);
        break;
      default:
        continue;
    }
  }
}

validDirections.sort((a, b) => a - b);
if (validDirections[0] === 0) {
  if (validDirections[1] === 1) {
    input[start[1]][start[0]] = "F";
  } else if (validDirections[1] === 2) {
    input[start[1]][start[0]] = "7";
  } else {
    input[start[1]][start[0]] = "|";
  }
} else if (validDirections[0] === 1) {
  if (validDirections[1] === 2) {
    input[start[1]][start[0]] = "-";
  } else {
    input[start[1]][start[0]] = "L";
  }
} else {
  input[start[1]][start[0]] = "J";
}

let [x, y] = start;
let dir = validDirections[0];
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
