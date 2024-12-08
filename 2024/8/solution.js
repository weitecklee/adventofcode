const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n");

const antennaMap = new Map();
for (let r = 0; r < input.length; r++) {
  for (let c = 0; c < input[0].length; c++) {
    if (input[r][c] !== ".") {
      if (!antennaMap.has(input[r][c])) {
        antennaMap.set(input[r][c], []);
      }
      antennaMap.get(input[r][c]).push([r, c]);
    }
  }
}

const hgt = input.length;
const wid = input[0].length;

const antinodes = new Set();
const antinodes2 = new Set();

for (const coords of antennaMap.values()) {
  for (let i = 0; i < coords.length; i++) {
    for (let j = i + 1; j < coords.length; j++) {
      const [r1, c1] = coords[i];
      const [r2, c2] = coords[j];
      const dr = r2 - r1;
      const dc = c2 - c1;
      if (r1 - dr >= 0 && r1 - dr < hgt && c1 - dc >= 0 && c1 - dc < wid)
        antinodes.add(`${r1 - dr},${c1 - dc}`);
      let k = 0;
      while (
        r1 - dr * k >= 0 &&
        r1 - dr * k < hgt &&
        c1 - dc * k >= 0 &&
        c1 - dc * k < wid
      ) {
        antinodes2.add(`${r1 - dr * k},${c1 - dc * k}`);
        k++;
      }
      if (r2 + dr >= 0 && r2 + dr < hgt && c2 + dc >= 0 && c2 + dc < wid)
        antinodes.add(`${r2 + dr},${c2 + dc}`);
      k = 0;
      while (
        r2 + dr * k >= 0 &&
        r2 + dr * k < hgt &&
        c2 + dc * k >= 0 &&
        c2 + dc * k < wid
      ) {
        antinodes2.add(`${r2 + dr * k},${c2 + dc * k}`);
        k++;
      }
    }
  }
}

console.log(antinodes.size);
console.log(antinodes2.size);
