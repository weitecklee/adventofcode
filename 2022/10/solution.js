const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8", (err, data) => {
    if (err) {
      console.log(err);
    } else {
      return data;
    }
  })
  .split("\n");

const signal = [1];

for (const line of input) {
  const last = signal[signal.length - 1];
  signal.push(last);
  if (line[0] !== "n") {
    const oper = line.split(" ");
    signal.push(last + Number(oper[1]));
  }
}

const qs = [20, 60, 100, 140, 180, 220];
let sum = 0;
for (const q of qs) {
  sum += q * signal[q - 1];
}

console.log(sum);

const screen = new Array(6);
for (let i = 0; i < 6; i++) {
  screen[i] = " ".repeat(40).split("");
}

for (let i = 0; i < 240; i++) {
  const x = i % 40;
  const y = Math.floor(i / 40);
  if (Math.abs(signal[i] - x) <= 1) {
    screen[y][x] = "#";
  }
}
for (let i = 0; i < 6; i++) {
  screen[i] = screen[i].join("");
  console.log(screen[i]);
}
