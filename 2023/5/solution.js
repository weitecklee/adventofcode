const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n\n");

const seeds = input[0].split(" ").slice(1).map(Number);

class ConversionMap {
  constructor(input) {
    this.name = input[0];
    this.ranges = input.slice(1).map((a) => a.split(" ").map(Number));
  }

  convert(num) {
    for (const range of this.ranges) {
      if (num >= range[1] && num < range[1] + range[2]) {
        return num - range[1] + range[0];
      }
    }
    return num;
  }

  deconvert(num) {
    for (const range of this.ranges) {
      if (num >= range[0] && num < range[0] + range[2]) {
        return num - range[0] + range[1];
      }
    }
    return num;
  }
}

const conversionMaps = [];
for (let i = 1; i < input.length; i++) {
  conversionMaps.push(new ConversionMap(input[i].split("\n")));
}

const part1 = Math.min(
  ...seeds.map((a) => conversionMaps.reduce((a, b) => b.convert(a), a))
);
console.log(part1);

const seedRanges = [];
for (let i = 0; i < seeds.length; i += 2) {
  seedRanges.push([seeds[i], seeds[i + 1]]);
}

let part2 = 0;
while (true) {
  let num = part2;
  for (let i = conversionMaps.length - 1; i >= 0; i--) {
    num = conversionMaps[i].deconvert(num);
  }
  if (seedRanges.some((a) => num >= a[0] && num < a[0] + a[1])) {
    console.log(part2);
    break;
  }
  part2++;
}
