const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n");

class Instruction {
  constructor(line) {
    const [a, b, c] = line.split(" ");
    this.dir = a;
    this.dist = Number(b);
    this.color = c.slice(1, -1);
    this.dir2 = ["R", "D", "L", "U"][Number(this.color[this.color.length - 1])];
    this.dist2 = parseInt(this.color.slice(1, 6), 16);
  }
}

const instructions = input.map((line) => new Instruction(line));

function calculateSolution(instructions) {
  // instructions: []{str dir, int dist}
  // Combination of Shoelace formula and Pick's theorem
  // Shoelace formula calculates area of polygon given vertex coordinates
  // Pick's theorem calculates area of polygon given number of boundary points and number of interior points
  // We have vertices (and boundary points) from input, and ultimately want to find sum of interior points and boundary points
  // Shoelace formula
  // https://en.wikipedia.org/wiki/Shoelace_formula
  // Pick's theorem
  // https://en.wikipedia.org/wiki/Pick%27s_theorem

  let x0 = 0; // vertex coordinates (start from origin)
  let y0 = 0;
  let x1, y1;
  let sum = 0;
  let boundaryPoints = 0;

  for (const { dir, dist } of instructions) {
    [x1, y1] = [x0, y0];
    switch (dir) {
      case "U":
        y1 = y0 - dist;
        break;
      case "D":
        y1 = y0 + dist;
        break;
      case "L":
        x1 = x0 - dist;
        break;
      case "R":
        x1 = x0 + dist;
        break;
    }
    boundaryPoints += dist; // number of boundary points is sum of distances
    sum += x0 * y1 - x1 * y0;
    x0 = x1;
    y0 = y1;
  }

  const area = Math.abs(sum) / 2; // from Shoelace formula
  const interiorPoints = area - boundaryPoints / 2 + 1; // Pick's theorem rearranged

  return interiorPoints + boundaryPoints;
}

const part1 = calculateSolution(instructions);
console.log(part1);
const part2 = calculateSolution(
  instructions.map((i) => ({ dir: i.dir2, dist: i.dist2 }))
);
console.log(part2);
