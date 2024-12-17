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

let program = input[4][1].split(",").map(Number);

let i = 0;
const output = [];

try {
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
        comboOperand = register.get("A");
        break;
      case 5:
        comboOperand = register.get("B");
        break;
      case 6:
        comboOperand = register.get("C");
        break;
    }
    switch (instruction) {
      case 0:
        register.set("A", Math.trunc(register.get("A") / 2 ** comboOperand));
        break;
      case 1:
        register.set("B", register.get("B") ^ literalOperand);
        break;
      case 2:
        register.set("B", comboOperand % 8);
        break;
      case 3:
        if (register.get("A") !== 0) {
          i = literalOperand;
          continue;
        }
        break;
      case 4:
        register.set("B", register.get("B") ^ register.get("C"));
        break;
      case 5:
        output.push(comboOperand % 8);
        break;
      case 6:
        register.set("B", Math.trunc(register.get("A") / 2 ** comboOperand));
        break;
      case 7:
        register.set("C", Math.trunc(register.get("A") / 2 ** comboOperand));
        break;
      default:
        throw new Error("Unknown instruction: ", instruction);
    }
    i += 2;
  }
} catch (e) {
  console.error(e);
  console.log(register);
} finally {
  console.log(output.join(","));
}
