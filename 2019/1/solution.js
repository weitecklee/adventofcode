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
  .split("\n")
  .map(Number);

const calculateFuelRequirement = (mass) => Math.floor(mass / 3) - 2;

const part1 = input.reduce((a, b) => a + calculateFuelRequirement(b), 0);

const calculateFuelRequirement2 = (mass) => {
  const fuel = calculateFuelRequirement(mass);
  return fuel > 0 ? fuel + calculateFuelRequirement2(fuel) : 0;
};

const part2 = input.reduce((a, b) => a + calculateFuelRequirement2(b), 0);

console.log(part1);
console.log(part2);
