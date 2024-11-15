const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n");

const timestamp = Number(input[0]);
const buses = input[1]
  .split(",")
  .filter((a) => a !== "x")
  .map(Number);

let minWait = Infinity;
let part1 = 0;
for (const bus of buses) {
  const wait = Math.ceil(timestamp / bus) * bus - timestamp;
  if (wait < minWait) {
    minWait = wait;
    part1 = bus * wait;
  }
}
console.log(part1);

// Chinese remainder theorem

const schedule = input[1]
  .split(",")
  .map((a, i) => (a === "x" ? null : [Number(a), i]))
  .filter(Boolean);

let part2 = 0;
let period = schedule[0][0];
for (let i = 1; i < schedule.length; i++) {
  const [bus, offset] = schedule[i];
  while ((part2 + offset) % bus !== 0) {
    part2 += period;
  }
  period *= bus;
}
console.log(part2);
