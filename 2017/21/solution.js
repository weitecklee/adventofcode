const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split(" => "));

const rules = new Map();
for (const [pattern, res] of input) {
  rules.set(pattern, res);
}

let patterns = [`.#.\/..#\/###`];

function rotate(pattern) {
  if (pattern.length === 5) {
    return pattern[3] + pattern[0] + "/" + pattern[4] + pattern[1];
  }
  return (
    pattern[8] +
    pattern[4] +
    pattern[0] +
    "/" +
    pattern[9] +
    pattern[5] +
    pattern[1] +
    "/" +
    pattern[10] +
    pattern[6] +
    pattern[2]
  );
}

function flip(pattern) {
  return pattern
    .split("/")
    .map((a) => a.split("").reverse().join(""))
    .join("/");
}

function breakup(patterns) {
  // turn each 4x4 to four 2x2's
  const res = [];
  for (const pattern of patterns) {
    for (let i = 0; i < 2; i++) {
      for (let j = 0; j < 2; j++) {
        res.push(
          pattern[2 * j + 10 * i] +
            pattern[2 * j + 10 * i + 1] +
            "/" +
            pattern[2 * j + 10 * i + 5] +
            pattern[2 * j + 10 * i + 6]
        );
      }
    }
  }
  return res;
}

function combine(patterns) {
  // turn each 6x6 to nine 2x2's
  const res = [];
  for (let i = 0; i < patterns.length; i += 4) {
    // yeah I'm just gonna manually type all this out
    // it does the job well enough
    res.push(
      patterns[i][0] + patterns[i][1] + "/" + patterns[i][4] + patterns[i][5]
    );
    res.push(
      patterns[i][2] +
        patterns[i + 1][0] +
        "/" +
        patterns[i][6] +
        patterns[i + 1][4]
    );
    res.push(
      patterns[i + 1][1] +
        patterns[i + 1][2] +
        "/" +
        patterns[i + 1][5] +
        patterns[i + 1][6]
    );
    res.push(
      patterns[i][8] +
        patterns[i][9] +
        "/" +
        patterns[i + 2][0] +
        patterns[i + 2][1]
    );
    res.push(
      patterns[i][10] +
        patterns[i + 1][8] +
        "/" +
        patterns[i + 2][2] +
        patterns[i + 3][0]
    );
    res.push(
      patterns[i + 1][9] +
        patterns[i + 1][10] +
        "/" +
        patterns[i + 3][1] +
        patterns[i + 3][2]
    );
    res.push(
      patterns[i + 2][4] +
        patterns[i + 2][5] +
        "/" +
        patterns[i + 2][8] +
        patterns[i + 2][9]
    );
    res.push(
      patterns[i + 2][6] +
        patterns[i + 3][4] +
        "/" +
        patterns[i + 2][10] +
        patterns[i + 3][8]
    );
    res.push(
      patterns[i + 3][5] +
        patterns[i + 3][6] +
        "/" +
        patterns[i + 3][9] +
        patterns[i + 3][10]
    );
  }
  return res;
}

function iterate(patterns, i) {
  // when i % 3 === 0, pattern is divisible by 3 and not 2, no further action needed
  // when i % 3 === 1, pattern is made up of 4x4's, break them up to 2x2's
  // when i % 3 === 2, pattern is made up of 3x3's BUT pattern size is divisible by 2
  //                   combine neighboring 3x3's into 6x6's, then break them up to 2x2's.
  if (i % 3 === 1) {
    patterns = breakup(patterns);
  } else if (i % 3 === 2) {
    patterns = combine(patterns);
  }
  let tmp = [];
  for (let pattern of patterns) {
    let rot = 0;
    while (!rules.has(pattern)) {
      pattern = rotate(pattern);
      rot++;
      if (rot === 4) {
        pattern = flip(pattern);
      }
      if (rot === 8) {
        throw new Error("No matching pattern found");
      }
    }
    let res = rules.get(pattern);
    tmp.push(res);
  }
  return tmp;
}

for (let i = 0; i < 5; i++) {
  patterns = iterate(patterns, i);
}

console.log(
  patterns.reduce((a, b) => a + b.split("").filter((c) => c === "#").length, 0)
);

for (let i = 5; i < 18; i++) {
  patterns = iterate(patterns, i);
}

console.log(
  patterns.reduce((a, b) => a + b.split("").filter((c) => c === "#").length, 0)
);
