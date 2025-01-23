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
  .split(",")
  .map(Number);

function intcode(input, noun, verb) {
  input[1] = noun;
  input[2] = verb;
  let pos = 0;
  while (true) {
    pos %= input.length;
    const opcode = input[pos];
    switch (opcode) {
      case 1:
      case 2:
        const inputNum1 = input[input[pos + 1]];
        const inputNum2 = input[input[pos + 2]];
        const outputPos = input[pos + 3];
        if (opcode === 1) {
          input[outputPos] = inputNum1 + inputNum2;
        } else {
          input[outputPos] = inputNum1 * inputNum2;
        }
        pos += 4;
        break;
      case 99:
        return input[0];
        break;
      default:
        console.error("Unknown opcode: ", opcode);
    }
  }
}

function part1() {
  return intcode(input.slice(), 12, 2);
}

function part2(target) {
  for (let i = 0; i < 100; i++) {
    for (let j = 0; j < 100; j++) {
      if (intcode(input.slice(), i, j) === target) {
        return i * 100 + j;
      }
    }
  }
}

console.log(part1());
console.log(part2(19690720));
