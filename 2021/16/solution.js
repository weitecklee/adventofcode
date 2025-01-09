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
  const packetVersion = parseInt(message.slice(i, (i += 3)), 2);
  part1 += packetVersion;
  const packetTypeID = parseInt(message.slice(i, (i += 3)), 2);
  if (packetTypeID === 4) {
    const value = [];
    while (message[i] !== "0") {
      value.push(message.slice(++i, (i += 4)));
    }
    value.push(message.slice(++i, (i += 4)));
    return parseInt(value.join(""), 2);
  }
  const subPackets = [];
  if (message[i++] === "0") {
    const totalLength = parseInt(message.slice(i, (i += 15)), 2);
    let tmp = i;
    while (i < tmp + totalLength) {
      subPackets.push(decodeMessage());
    }
  } else {
    const numSubPackets = parseInt(message.slice(i, (i += 11)), 2);
    for (let j = 0; j < numSubPackets; j++) {
      subPackets.push(decodeMessage());
    }
  }
  switch (packetTypeID) {
    case 0:
      return subPackets.reduce((a, b) => a + b, 0);
    case 1:
      return subPackets.reduce((a, b) => a * b, 1);
    case 2:
      return Math.min(...subPackets);
    case 3:
      return Math.max(...subPackets);
    case 5:
      return subPackets[0] > subPackets[1] ? 1 : 0;
    case 6:
      return subPackets[0] < subPackets[1] ? 1 : 0;
    case 7:
      return subPackets[0] === subPackets[1] ? 1 : 0;
  }
}

const part2 = decodeMessage();
console.log(part1);
console.log(part2);
