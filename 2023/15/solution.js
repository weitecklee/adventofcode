const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split(",");

function executeHASH(s) {
  let curr = 0;
  for (const c of s) {
    curr += c.charCodeAt(0);
    curr *= 17;
    curr %= 256;
  }
  return curr;
}

const part1 = input.reduce((a, b) => a + executeHASH(b), 0);
console.log(part1);

class Boxes {
  constructor() {
    this.boxes = new Map();
    for (let i = 0; i < 256; i++) {
      // use Map instead of Array for better insertion/deletion performance,
      // iteration happens in insertion order which satisfies requirements
      this.boxes.set(i, new Map());
    }
  }

  remove(label, boxNumber) {
    const box = this.boxes.get(boxNumber);
    box.delete(label);
  }

  add(label, boxNumber, focalLength) {
    const box = this.boxes.get(boxNumber);
    box.set(label, { focalLength: Number(focalLength) });
  }

  focusingPower() {
    // let res = 0;
    // for (const [i, box] of this.boxes.entries()) {
    //   for (let j = 0; j < box.length; j++) {
    //     res += (i + 1) * (j + 1) * box[j].focalLength;
    //   }
    // }
    // return res;
    return Array.from(this.boxes.values()).reduce(
      (a, box, i) =>
        a +
        Array.from(box.values()).reduce(
          (b, lens, j) => b + (i + 1) * (j + 1) * lens.focalLength,
          0
        ),
      0
    );
  }
}

const boxes = new Boxes();

for (const s of input) {
  const [label, focalLength] = s.split(/[=\-]/);
  const boxNumber = executeHASH(label);
  if (focalLength) {
    boxes.add(label, boxNumber, focalLength);
  } else {
    boxes.remove(label, boxNumber);
  }
}

console.log(boxes.focusingPower());
