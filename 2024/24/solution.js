const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n\n");

const wires = new Map();

class Wire {
  constructor(name, input) {
    this.name = name;
    this.input = input;
  }

  get output() {
    const res = Number(this.input);
    if (!isNaN(res)) {
      return res;
    }
    const [wire1, op, wire2] = this.input.split(" ");
    switch (op) {
      case "AND":
        return wires.get(wire1).output & wires.get(wire2).output;
      case "OR":
        return wires.get(wire1).output | wires.get(wire2).output;
      case "XOR":
        return wires.get(wire1).output ^ wires.get(wire2).output;
      default:
        throw new Error("Unknown operation: ", op);
    }
  }
}

for (const line of input[0].split("\n")) {
  const [name, input] = line.split(": ");
  wires.set(name, new Wire(name, input));
}

for (const line of input[1].split("\n")) {
  const [input, name] = line.split(" -> ");
  wires.set(name, new Wire(name, input));
}

let part1 = "";
let i = 0;
while (true) {
  const wireName = `z${i.toString().padStart(2, "0")}`;
  if (!wires.has(wireName)) break;
  part1 += wires.get(wireName).output.toString();
  i++;
}
console.log(parseInt(part1.split("").reverse().join(""), 2));
