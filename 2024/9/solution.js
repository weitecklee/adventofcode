const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("")
  .map(Number);

let left = 0;
let right = (input.length - 1) % 2 ? input.length - 2 : input.length - 1;
let pos = 0;
let part1 = 0;
let waitingToBeSlotted = input[right];
while (left < right) {
  if (left % 2) {
    // empty space
    for (let i = 0; i < input[left]; i++) {
      if (waitingToBeSlotted--) {
        part1 += (right / 2) * pos++;
      } else {
        right -= 2;
        waitingToBeSlotted = input[right] - 1;
        part1 += (right / 2) * pos++;
      }
    }
  } else {
    // file space
    for (let i = 0; i < input[left]; i++) {
      part1 += (left / 2) * pos++;
    }
  }
  left++;
}
// take care of any still waiting to be slotted
for (let i = 0; i < waitingToBeSlotted; i++) {
  part1 += (right / 2) * pos++;
}

console.log(part1);

const input2 = input.slice();
const posMap = new Map(); // map of (file/empty) block to position, position points to start of block
pos = 0;
for (let i = 0; i < input2.length; i++) {
  posMap.set(i, pos);
  pos += input2[i];
}

let part2 = 0;

for (
  let right = (input2.length - 1) % 2 ? input2.length - 2 : input2.length - 1;
  right >= 0;
  right -= 2
) {
  // find first empty space that fits
  let pos = 1;
  while (pos < right && input2[pos] < input2[right]) {
    pos += 2;
  }
  if (pos < right) {
    // if found, slot file in
    // have to account for if other file(s) already slotted in previously
    // so must compare old input[pos] and new input[pos]
    for (let j = 0; j < input2[right]; j++) {
      part2 += (input[pos] - input2[pos] + posMap.get(pos) + j) * (right / 2);
    }
    input2[pos] -= input2[right];
    input2[right] = 0;
  }
}

// add up files that were not moved
for (let i = 0; i < input2.length; i += 2) {
  for (let j = 0; j < input2[i]; j++) {
    part2 += (posMap.get(i) + j) * (i / 2);
  }
}
console.log(part2);
