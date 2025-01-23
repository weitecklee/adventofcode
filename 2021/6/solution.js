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
  .split(",")
  .map(Number);

const fish = new Array(9).fill(0);

for (const n of input) {
  fish[n]++;
}

for (let days = 0; days < 80; days++) {
  const newFish = fish[0];
  for (let i = 0; i < fish.length - 1; i++) {
    fish[i] = fish[i + 1];
  }
  fish[6] += newFish;
  fish[8] = newFish;
}

console.log(fish.reduce((a, b) => a + b));

for (let days = 80; days < 256; days++) {
  const newFish = fish[0];
  for (let i = 0; i < fish.length - 1; i++) {
    fish[i] = fish[i + 1];
  }
  fish[6] += newFish;
  fish[8] = newFish;
}

console.log(fish.reduce((a, b) => a + b));
