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

const drops = new Set(input);
const air = new Set();

let area = 0;
let minX = Infinity;
let maxX = -Infinity;
let minY = Infinity;
let maxY = -Infinity;
let minZ = Infinity;
let maxZ = -Infinity;

const q = input.slice();

while (q.length) {
  const check = q.pop();
  const parse = check.split(",").map(Number);
  minX = Math.min(minX, parse[0]);
  maxX = Math.max(maxX, parse[0]);
  minY = Math.min(minY, parse[1]);
  maxY = Math.max(maxY, parse[1]);
  minZ = Math.min(minZ, parse[2]);
  maxZ = Math.max(maxZ, parse[2]);
  const toCheck = [];
  toCheck.push([parse[0] - 1, parse[1], parse[2]].join(","));
  toCheck.push([parse[0] + 1, parse[1], parse[2]].join(","));
  toCheck.push([parse[0], parse[1] - 1, parse[2]].join(","));
  toCheck.push([parse[0], parse[1] + 1, parse[2]].join(","));
  toCheck.push([parse[0], parse[1], parse[2] - 1].join(","));
  toCheck.push([parse[0], parse[1], parse[2] + 1].join(","));
  for (const entry of toCheck) {
    if (!drops.has(entry)) {
      area++;
    }
  }
}

console.log(area);

// Use outer bounds to find surrounding air, add to area when droplet cube is encountered

let area2 = 0;
const checked = new Set();
const first = [minX - 1, minY - 1, minZ - 1].join(",");
checked.add(first);
q.push(first);

while (q.length) {
  const check = q.pop();
  const parse = check.split(",").map(Number);
  const toCheck = [];
  if (parse[0] > minX - 1) {
    toCheck.push([parse[0] - 1, parse[1], parse[2]].join(","));
  }
  if (parse[0] < maxX + 1) {
    toCheck.push([parse[0] + 1, parse[1], parse[2]].join(","));
  }
  if (parse[1] > minY - 1) {
    toCheck.push([parse[0], parse[1] - 1, parse[2]].join(","));
  }
  if (parse[1] < maxY + 1) {
    toCheck.push([parse[0], parse[1] + 1, parse[2]].join(","));
  }
  if (parse[2] > minZ - 1) {
    toCheck.push([parse[0], parse[1], parse[2] - 1].join(","));
  }
  if (parse[2] < maxZ + 1) {
    toCheck.push([parse[0], parse[1], parse[2] + 1].join(","));
  }
  for (const entry of toCheck) {
    if (drops.has(entry)) {
      area2++;
    } else if (!checked.has(entry)) {
      q.push(entry);
      checked.add(entry);
    }
  }
}

console.log(area2);
