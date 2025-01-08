const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n\n");

const wireMap = new Map();

class Wire {
  constructor(name, input) {
    this.name = name;
    this.input = input;
    this.startingWire = false;
    this.initialize();
  }

  initialize() {
    const res = Number(this.input);
    if (isNaN(res)) {
      const [wire1, op, wire2] = this.input.split(" ");
      this.wire1 = wire1;
      this.op = op;
      this.wire2 = wire2;
    } else {
      this.startingWire = true;
      this.value = res;
    }
  }

  get output() {
    if (this.startingWire) return this.value;
    switch (this.op) {
      case "AND":
        return wireMap.get(this.wire1).output & wireMap.get(this.wire2).output;
      case "OR":
        return wireMap.get(this.wire1).output | wireMap.get(this.wire2).output;
      case "XOR":
        return wireMap.get(this.wire1).output ^ wireMap.get(this.wire2).output;
      default:
        throw new Error("Unknown operation: ", op);
    }
  }

  print() {
    console.log(this.name, this.input);
  }
}

for (const line of input[0].split("\n")) {
  const [name, input] = line.split(": ");
  wireMap.set(name, new Wire(name, input));
}

for (const line of input[1].split("\n")) {
  const [input, name] = line.split(" -> ");
  wireMap.set(name, new Wire(name, input));
  const [wire1, op, wire2] = input.split(" ");
}

function calculateBinaryNumber(wirePrefix) {
  // returns decimal value of binary number represented by wires
  let res = "";
  let i = 0;
  while (true) {
    const wireName = wirePrefix + i.toString().padStart(2, "0");
    if (!wireMap.has(wireName)) break;
    res = wireMap.get(wireName).output.toString() + res;
    i++;
  }
  return parseInt(res, 2);
}

console.log(calculateBinaryNumber("z"));

/*
  Mapping out the wires and gates:

  z00 = x00 XOR y00

  z01 = tcd XOR bwv      // tmp1 and tmp2 in code below
    tcd = x01 XOR y01
    bwv = x00 AND y00

  z02 = frj XOR hqq      // new tmp1 and tmp2 at end of each iteration
    frj = x02 XOR y02
    hqq = sgv OR wqt     // otherWire in code below
      sgv = tcd AND bwv  // use old tmp1 and tmp2 to find otherWire
      wqt = x01 AND y01

  z03 = ckv XOR bbh
    bbh = x03 XOR y03
    ckv = bkc OR wsq
      bkc = frj AND hqq
      wsq = x02 AND y02

  ...

  z44 = frk XOR wpk
    wpk = x44 XOR y44
    frk = mpm OR kpc
      kpc = vtf AND gdw
      mpm = x43 AND y43

  z45 = mkv OR hgp
    mkv = x44 AND y44
    hgp = frk AND wpk

  Assuming we're lucky and the misplaced gates are in the middle,
  we can build up the structure manually and figure out which wires
  and gates are wrong and swap them with the correct ones.
  Also lucky that each gate only has one wrong wire.

*/

// function printWireTree(wire) {
//   const queue = [wire];
//   for (let i = 0; i < queue.length; i++) {
//     const w = wireMap.get(queue[i]);
//     if (w.wire1) {
//       queue.push(w.wire1);
//     }
//     if (w.wire2) {
//       queue.push(w.wire2);
//     }
//     w.print();
//   }
// }

function findCorrectWire(wire1, op, wire2) {
  const input1 = [wire1, op, wire2].join(" ");
  const input2 = [wire2, op, wire1].join(" ");
  for (const wire of wireMap.values()) {
    if (wire.input === input1 || wire.input === input2) {
      return wire.name;
    }
  }
  return null;
}

function findOtherWire(xWirePrev, yWirePrev, tmp1, tmp2) {
  const wire1 = findCorrectWire(xWirePrev, "AND", yWirePrev);
  const wire2 = findCorrectWire(tmp1, "AND", tmp2);
  return findCorrectWire(wire1, "OR", wire2);
}

function swapWires(wire1, wire2) {
  // swap wire inputs and reinitialize each
  [wire1.input, wire2.input] = [wire2.input, wire1.input];
  wire1.initialize();
  wire2.initialize();
}

const wrongWires = [];

const z01 = wireMap.get("z01");
let tmp1 = z01.wire1;
let tmp2 = z01.wire2;
for (let i = 2; i < 45; i++) {
  const zWire = wireMap.get("z" + i.toString().padStart(2, "0"));
  const xWire = "x" + i.toString().padStart(2, "0");
  const yWire = "y" + i.toString().padStart(2, "0");
  const xWirePrev = "x" + (i - 1).toString().padStart(2, "0");
  const yWirePrev = "y" + (i - 1).toString().padStart(2, "0");
  if (zWire.op !== "XOR") {
    const wire1 = findCorrectWire(xWire, "XOR", yWire);
    const wire2 = findOtherWire(xWirePrev, yWirePrev, tmp1, tmp2);
    const correctZWire = findCorrectWire(wire1, "XOR", wire2);
    wrongWires.push(zWire.name, correctZWire);
    swapWires(zWire, wireMap.get(correctZWire));
  }
  const wire1 = wireMap.get(zWire.wire1);
  const wire2 = wireMap.get(zWire.wire2);
  const xorWire = findCorrectWire(xWire, "XOR", yWire);
  const otherWire = findOtherWire(xWirePrev, yWirePrev, tmp1, tmp2);
  if (wire1.name === xorWire) {
    if (wire2.name !== otherWire) {
      wrongWires.push(wire2.name, otherWire);
      swapWires(wire2, wireMap.get(otherWire));
    }
  } else if (wire2.name == xorWire) {
    if (wire1.name !== otherWire) {
      wrongWires.push(wire1.name, otherWire);
      swapWires(wire1, wireMap.get(otherWire));
    }
  } else {
    // assume not the case that both wires are wrong
    if (wire1.name === otherWire) {
      wrongWires.push(wire2.name, xorWire);
      swapWires(wire2, wireMap.get(xorWire));
    } else {
      wrongWires.push(wire1.name, xorWire);
      swapWires(wire1, wireMap.get(xorWire));
    }
  }
  tmp1 = zWire.wire1;
  tmp2 = zWire.wire2;
}

wrongWires.sort();
console.log(wrongWires.join(","));
