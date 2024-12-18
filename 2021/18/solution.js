const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n");

const twoDigitRegex = /\d{2,}/;

function reduce(fish) {
  while (true) {
    // explode
    let nests = 0;
    let explosion = false;
    for (let i = 0; i < fish.length; i++) {
      if (fish[i] === ",") continue;
      if (fish[i] === "[") nests++;
      else if (fish[i] === "]") nests--;
      else if (nests > 4) {
        let j = i;
        while (fish[j] !== "[" && fish[j] !== "]") j++;
        if (fish[j] === "[") continue;
        const nums = fish.slice(i, j).split(",").map(Number);
        fish = fish.slice(0, i - 1) + "0" + fish.slice(j + 1);
        let k = i - 2;
        while (k >= 0 && !/\d/.test(fish[k])) k--;
        if (k >= 0) {
          let k2 = k - 1;
          while (k2 >= 0 && /\d/.test(fish[k2])) k2--;
          const newLeftNum = Number(fish.slice(k2 + 1, k + 1)) + nums[0];
          fish = fish.slice(0, k2 + 1) + newLeftNum + fish.slice(k + 1);
          i += newLeftNum.toString().length + k2 - k;
        }
        k = i + 1;
        while (k < fish.length && !/\d/.test(fish[k])) k++;
        if (k < fish.length) {
          let k2 = k + 1;
          while (k2 < fish.length && /\d/.test(fish[k2])) k2++;
          const newRightNum = Number(fish.slice(k, k2)) + nums[1];
          fish = fish.slice(0, k) + newRightNum + fish.slice(k2);
        }
        explosion = true;
        break;
      }
    }
    if (explosion) continue;
    // split
    if (twoDigitRegex.test(fish)) {
      fish = fish.replace(twoDigitRegex, (match) => {
        const n = Number(match);
        return `[${Math.floor(n / 2)},${Math.ceil(n / 2)}]`;
      });
      continue;
    }
    return fish;
  }
}

function add(fish1, fish2) {
  return reduce(`[${fish1},${fish2}]`);
}

function magnitude(fish) {
  if (typeof fish === "string") fish = JSON.parse(fish);
  if (typeof fish === "number") return fish;
  return 3 * magnitude(fish[0]) + 2 * magnitude(fish[1]);
}

const part1 = input.reduce((a, b) => add(a, b));
console.log(magnitude(part1));

let max = 0;
for (const fish1 of input) {
  for (const fish2 of input) {
    if (fish1 === fish2) continue;
    max = Math.max(max, magnitude(add(fish1, fish2)));
  }
}

console.log(max);
