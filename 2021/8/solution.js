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

let count = 0;
let outputs = 0;

const master = [
  "abcefg",
  "cf",
  "acdeg",
  "acdfg",
  "bcdf",
  "abdfg",
  "abdefg",
  "acf",
  "abcdefg",
  "abcdfg",
];

for (const line of input) {
  const s = line.match(/\w+/g);
  const patterns = s.slice(0, 10);
  const values = s.slice(10);
  const lengths = new Map();
  for (let i = 0; i < 10; i++) {
    patterns[i] = patterns[i].split("").sort();
    if (!lengths.has(patterns[i].length)) {
      lengths.set(patterns[i].length, []);
    }
    lengths.get(patterns[i].length).push(patterns[i]);
  }
  const decode = new Map();
  for (const s of lengths.get(3)[0]) {
    if (!lengths.get(2)[0].includes(s)) {
      decode.set(s, "a");
      break;
    }
  }
  for (let i = 0; i < 2; i++) {
    for (const s2 of lengths.get(6)) {
      if (!s2.includes(lengths.get(2)[0][i])) {
        decode.set(lengths.get(2)[0][i], "c");
        decode.set(lengths.get(2)[0][1 - i], "f");
        break;
      }
    }
  }
  for (let i = 0; i < 4; i++) {
    if (decode.has(lengths.get(4)[0][i])) {
      continue;
    }
    for (const s2 of lengths.get(6)) {
      if (!s2.includes(lengths.get(4)[0][i])) {
        decode.set(lengths.get(4)[0][i], "d");
        break;
      }
    }
  }
  for (const s of lengths.get(4)[0]) {
    if (!decode.has(s)) {
      decode.set(s, "b");
      break;
    }
  }
  for (const p of lengths.get(6)) {
    let n = 0;
    let odd;
    for (const s of p) {
      if (!decode.has(s)) {
        n++;
        odd = s;
      }
    }
    if (n === 1) {
      decode.set(odd, "g");
      break;
    }
  }
  for (const s of lengths.get(7)[0]) {
    if (!decode.has(s)) {
      decode.set(s, "e");
      break;
    }
  }
  const encode = new Map();
  for (const [a, b] of decode) {
    encode.set(b, a);
  }
  const digits = new Map();
  for (let i = 0; i < 10; i++) {
    const pattern = [];
    for (const c of master[i]) {
      pattern.push(encode.get(c));
    }
    digits.set(pattern.sort().join(""), i);
  }
  let output = "";
  for (const value of values) {
    if (
      value.length === 2 ||
      value.length === 3 ||
      value.length === 4 ||
      value.length === 7
    ) {
      count++;
    }
    const str = value.split("").sort().join("");
    output += digits.get(str);
  }
  outputs += Number(output);
}

console.log(count);
console.log(outputs);
