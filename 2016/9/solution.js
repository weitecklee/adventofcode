const fs = require("fs");
const path = require("path");

const input = fs.readFileSync(
  path.join(__dirname, "input.txt"),
  "utf-8",
  (err, data) => {
    if (err) {
      console.log(err);
    } else {
      return data;
    }
  }
);

const decompress = (code) => {
  let i = 0;
  let res = 0;
  while (i < code.length) {
    while (i < code.length && code[i] !== "(") {
      res++;
      i++;
    }
    if (i >= code.length) {
      break;
    }
    let j = i;
    while (code[j] !== ")") {
      j++;
    }
    const marker = code.slice(i + 1, j);
    const nums = marker.split("x").map(Number);
    res += nums[0] * nums[1];
    i = j + 1 + nums[0];
  }
  return res;
};

const decompress2 = (code) => {
  let i = 0;
  let res = 0;
  while (i < code.length) {
    while (i < code.length && code[i] !== "(") {
      res++;
      i++;
    }
    if (i >= code.length) {
      break;
    }
    let j = i;
    let balance = 1;
    while (code[j] !== ")") {
      j++;
    }
    const marker = code.slice(i + 1, j);
    const nums = marker.split("x").map(Number);
    res += decompress2(code.slice(j + 1, j + 1 + nums[0])) * nums[1];
    i = j + 1 + nums[0];
  }
  return res;
};

console.log(decompress(input));
console.log(decompress2(input));
