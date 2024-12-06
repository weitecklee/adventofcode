const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split(" "))
  .map((a) => [a[0], a[1].split(",").map(Number)]);

function countArrangements(pattern, groupings) {
  const regex = new RegExp(
    "^\\.*" + groupings.map((a) => `#{${a}}`).join("\\.+") + "\\.*$"
  );
  let permutations = [[]];
  let count = 0;
  for (const c of pattern) {
    if (c === "?") {
      const newPermutations = [];
      for (const p of permutations) {
        newPermutations.push([...p, "."]);
        newPermutations.push([...p, "#"]);
      }
      permutations = newPermutations;
    } else {
      permutations.forEach((p) => p.push(c));
    }
  }
  for (const p of permutations) {
    if (regex.test(p.join(""))) count++;
  }

  return count;
}

const part1 = input.reduce(
  (acc, [pattern, groupings]) => acc + countArrangements(pattern, groupings),
  0
);
console.log(part1);
