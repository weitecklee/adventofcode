const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split(" = "));

const memory = new Map();
const masks = [];

class Mask {
  constructor(mask) {
    this.mask = mask;
    this.writes = [];
  }

  addWrite(line) {
    const memoryAddress = Number(line[0].slice(4, -1));
    const value = Number(line[1]);
    this.writes.push([memoryAddress, value]);
  }

  writeValues1(memory) {
    for (const [memoryAddress, value] of this.writes) {
      const binaryValue = value.toString(2).padStart(36, "0");
      const maskedValue = binaryValue
        .split("")
        .map((a, i) => (this.mask[i] === "X" ? a : this.mask[i]))
        .join("");
      memory.set(memoryAddress, parseInt(maskedValue, 2));
    }
  }

  writeValues2(memory) {
    for (const [memoryAddress, value] of this.writes) {
      const binaryAddress = memoryAddress.toString(2).padStart(36, "0");
      const maskedAddress = binaryAddress
        .split("")
        .map((a, i) =>
          this.mask[i] === "0" ? a : this.mask[i] === "1" ? "1" : "X"
        )
        .join("");
      let addresses = [[]];
      for (const bit of maskedAddress) {
        if (bit === "X") {
          const newAddresses = [];
          for (const address of addresses) {
            newAddresses.push([...address, 0]);
            newAddresses.push([...address, 1]);
          }
          addresses = newAddresses;
        } else {
          for (const address of addresses) {
            address.push(bit);
          }
        }
      }
      for (const address of addresses) {
        memory.set(parseInt(address.join(""), 2), value);
      }
    }
  }
}

let currentMask;

for (const line of input) {
  if (line[0] === "mask") {
    currentMask = new Mask(line[1]);
    masks.push(currentMask);
  } else {
    currentMask.addWrite(line);
  }
}

for (const mask of masks) {
  mask.writeValues1(memory);
}
const part1 = Array.from(memory.values()).reduce((a, b) => a + b);
console.log(part1);

memory.clear();

for (const mask of masks) {
  mask.writeValues2(memory);
}
const part2 = Array.from(memory.values()).reduce((a, b) => a + b);
console.log(part2);
