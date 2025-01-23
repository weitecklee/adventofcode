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

let part1 = 0;

for (const line of input) {
  let i = 0;
  while (i < line.length) {
    const tmp = Number(line[i]);
    if (!isNaN(tmp)) {
      part1 += tmp * 10;
      break;
    }
    i++;
  }

  i = line.length - 1;
  while (i >= 0) {
    const tmp = Number(line[i]);
    if (!isNaN(tmp)) {
      part1 += tmp;
      break;
    }
    i--;
  }
}

console.log(part1);

let part2 = 0;

const digits = new Map();
digits.set("one", 1);
digits.set("two", 2);
digits.set("three", 3);
digits.set("four", 4);
digits.set("five", 5);
digits.set("six", 6);
digits.set("seven", 7);
digits.set("eight", 8);
digits.set("nine", 9);

for (const line of input) {
  let i = 0;
  while (i < line.length) {
    const tmp = Number(line[i]);
    if (!isNaN(tmp)) {
      part2 += tmp * 10;
      break;
    } else {
      const tmp3 = line.slice(i, i + 3);
      if (digits.has(tmp3)) {
        part2 += digits.get(tmp3) * 10;
        break;
      }
      const tmp4 = line.slice(i, i + 4);
      if (digits.has(tmp4)) {
        part2 += digits.get(tmp4) * 10;
        break;
      }
      const tmp5 = line.slice(i, i + 5);
      if (digits.has(tmp5)) {
        part2 += digits.get(tmp5) * 10;
        break;
      }
    }
    i++;
  }

  i = line.length - 1;
  while (i >= 0) {
    const tmp = Number(line[i]);
    if (!isNaN(tmp)) {
      part2 += tmp;
      break;
    } else {
      const tmp3 = line.slice(i - 2, i + 1);
      if (digits.has(tmp3)) {
        part2 += digits.get(tmp3);
        break;
      }
      const tmp4 = line.slice(i - 3, i + 1);
      if (digits.has(tmp4)) {
        part2 += digits.get(tmp4);
        break;
      }
      const tmp5 = line.slice(i - 4, i + 1);
      if (digits.has(tmp5)) {
        part2 += digits.get(tmp5);
        break;
      }
    }
    i--;
  }
}

console.log(part2);
