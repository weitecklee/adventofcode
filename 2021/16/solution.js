const fs = require("fs");
const path = require("path");
const MinHeap = require("../../utils/MinHeap");

const input = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8");

const message = input
  .split("")
  .map((c) => parseInt(c, 16).toString(2).padStart(4, "0"))
  .join("");

let part1 = 0;

let i = 0;

function decodeMessage() {
  const packetVersion = parseInt(message.slice(i, i + 3), 2);
  part1 += packetVersion;
  i += 3;
  const packetTypeID = parseInt(message.slice(i, i + 3), 2);
  i += 3;
  if (packetTypeID === 4) {
    while (message[i] !== "0") {
      i += 5;
    }
    i += 5;
  } else {
    if (message[i] === "0") {
      i += 1;
      const totalLength = parseInt(message.slice(i, i + 15), 2);
      i += 15;
      let tmp = i;
      while (i < tmp + totalLength) {
        decodeMessage();
      }
    } else {
      i += 1;
      const numSubPackets = parseInt(message.slice(i, i + 11), 2);
      i += 11;
      for (let j = 0; j < numSubPackets; j++) {
        decodeMessage();
      }
    }
  }
}

decodeMessage();
console.log(part1);
