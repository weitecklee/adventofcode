const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split(": "));

const register = new Map([
  ["A", Number(input[0][1])],
  ["B", Number(input[1][1])],
  ["C", Number(input[2][1])],
]);
const programString = input[4][1];
const program = input[4][1].split(",").map(Number);

function runProgram(registerA, registerB = 0, registerC = 0) {
  let i = 0;
  const output = [];
  while (i >= 0 && i < program.length) {
    const instruction = program[i];
    const literalOperand = program[i + 1];
    let comboOperand;
    switch (literalOperand) {
      case 0:
      case 1:
      case 2:
      case 3:
        comboOperand = literalOperand;
        break;
      case 4:
        comboOperand = registerA;
        break;
      case 5:
        comboOperand = registerB;
        break;
      case 6:
        comboOperand = registerC;
        break;
    }
    switch (instruction) {
      case 0:
        registerA = Math.trunc(registerA / 2 ** comboOperand);
        break;
      case 1:
        registerB = registerB ^ literalOperand;
        break;
      case 2:
        registerB = comboOperand & 7;
        break;
      case 3:
        if (registerA !== 0) {
          i = literalOperand;
          continue;
        }
        break;
      case 4:
        registerB = registerB ^ registerC;
        break;
      case 5:
        output.push(comboOperand & 7);
        break;
      case 6:
        registerB = Math.trunc(registerA / 2 ** comboOperand);
        break;
      case 7:
        registerC = Math.trunc(registerA / 2 ** comboOperand);
        break;
      default:
        throw new Error("Unknown instruction: ", instruction);
    }
    i += 2;
  }
  return output;
}

console.log(
  runProgram(register.get("A"), register.get("B"), register.get("C")).join(",")
);

let j = 0;
for (let i = program.length - 1; i >= 0; i--) {
  j *= 8;
  const currTarget = program.slice(i).join(",");
  while (true) {
    const curr = runProgram(j).join(",");
    if (curr === currTarget) {
      break;
    }
    j++;
  }
}
console.log(j);
