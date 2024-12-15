const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n");

let sum = 0;
for (const line of input) {
  let i = 0;
  let n = 0;
  while (i < line.length) {
    switch (line[line.length - 1 - i]) {
      case "2":
        n += 2 * 5 ** i;
        break;
      case "1":
        n += 5 ** i;
        break;
      case "-":
        n -= 5 ** i;
        break;
      case "=":
        n -= 2 * 5 ** i;
        break;
    }
    i++;
  }

  sum += n;
}

const snafu = ["0", ...sum.toString(5).split("")].map(Number);

for (let i = snafu.length - 1; i > 0; i--) {
  snafu[i - 1] += Math.floor(snafu[i] / 5);
  snafu[i] %= 5;
  if (snafu[i] >= 3) {
    snafu[i - 1] += 1;
    snafu[i] = snafu[i] === 3 ? "=" : "-";
  }
}
if (snafu[0] === 0) snafu.shift();

console.log(snafu.join(""));
