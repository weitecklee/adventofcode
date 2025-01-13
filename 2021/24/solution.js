const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split(" "));

/*
  Actually solved first by hand by going through input and noticing patterns.
  Input is comprised of 14 subprograms (one for each digit) that are
  largely similar except the 5th, 6th, and 16th lines. (1-index)
  5th line is always "div z N" where N is either 1 or 26.
  6th line is always "add x N" where N is either positive or negative number.
  Crucially, whenever 5th line is "div z 1", 6th line is always
  "add x A" where A is a positive number. Whenever 5th line is
  "div z 26", 6th line is always "add x B" where B is a negative number.
  16th line is always "add y N" where N is positive number.
  What ends up happening is that for z to end up equal to 0 at the end of the
  program, there will be constraints placed on pairs of digits (e.g.,
  digits[0] + 4 = digits[13], digits[1] + 8 = digits[12], etc.)
  Figure out the constraints, figure out the maximum and minimum values of
  the digits within the constraints (digits[i] != 0) and there's your answer.

  The magic happens at every "eql x w" line, which is always followed by an
  "eql x 0" line. We want the first line to end up 1 (x = w) so that the second
  line ends up 0. When this happens, x wil be set to some digits[i] + N and w will
  always be set to some digits[j], and this is where we find our constraints.
  Since we know a digit can only be 1 to 9, N must be within range [-8, 8] for the
  equality to have any chance of being true.

  I can only assume that other puzzle inputs have a similar setup with just the lines
  mentioned above differing in the values of the numbers.
  Otherwise this code probably only works for my puzzle input...
*/

function copy(a) {
  return JSON.parse(JSON.stringify(a));
}

let w = [],
  x = 0, // either 0, 1, or [a, b] where a = i (index of digits) and y = value added to digit
  y = 0, // [a, b] where a = i (index of digits) and y = value added to digit
  // actual equation would be Z = 26 * (26 * ... * (26 * (26 * z[0] + z[1]) + z[2]) + ... z[n-1]) + z[n]
  i = 0; // index of digits
const z = [[]]; // array of [a, b] where a = i (index of digits) and y = value added to digit
const constraints = [];
for (const [instruction, a, b] of input) {
  switch (instruction) {
    case "inp": // only ever "inp w"
      w = [i++, 0];
      break;
    case "mod": // only ever "mod x 26" and after x has been set equal to z
      x = x.pop(); // "mod x 26" is then same as setting it equal to last entry of array
      break;
    case "div": // only ever "div z N" where N is either 1 or 26
      if (b === "26") z.pop(); // if N = 1, nothing happens; if N = 26, pop off last entry of z
      break;
    case "mul":
      switch (a) {
        case "x": // only ever "mul x 0"
          x = 0;
          break;
        case "y": // either "mul y x" or "mul y 0", only care when x == 0
          if (b === "x" && x === 0) y = 0;
          break;
      } // ignore "mul z y" instructions
      break;
    case "add":
      switch (a) {
        case "x": // either "add x z" or "add x N"
          if (b === "z") x = copy(z);
          else x[1] += Number(b);
          break;
        case "y": // either "add y w" or "add y N"
          if (b === "w") y = copy(w);
          else typeof y === "number" ? (y += Number(b)) : (y[1] += Number(b));
          break;
        case "z": // only ever "add z y"
          if (y !== 0) z.push(copy(y));
          break;
      }
      break;
    case "eql": // either "eql x w" or "eql x 0", only care when it's "eql x 0"
      if (b === "0") {
        if (Math.abs(x[1]) < 9) {
          constraints.push([copy(w), copy(x)]);
          x = 0;
        }
      }
      break;
  }
}

const part1 = Array(14).fill(9);
const part2 = Array(14).fill(1);

for (let [[i, a], [j, b]] of constraints) {
  if (b < 0) {
    a = -b;
    b = 0;
  }
  part1[i] -= a;
  part1[j] -= b;
  part2[i] += b;
  part2[j] += a;
}

console.log(part1.join(""));
console.log(part2.join(""));
