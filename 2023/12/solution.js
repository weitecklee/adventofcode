const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split(" "))
  .map((a) => [a[0], a[1].split(",").map(Number)]);

const memo = new Map();
const memoKey = (pattern, groupings) => pattern + "|" + groupings.join(",");

/*
  Phew. This took a while. countArrangements took a fair bit of debugging
  to eventually get down (and it still looks pretty ugly). Even after
  I got it working, it looked like it was gonna take days to solve part 2.
  Thankfully, with memoization  it only takes half a second.
*/

function countArrangements(pattern, groupings) {
  const key = memoKey(pattern, groupings);
  if (memo.has(key)) return memo.get(key);

  // check if groupings can fit in pattern
  const groupingsLength = groupings.reduce((a, b) => a + b + 1);
  if (groupingsLength > pattern.length) return 0;

  let count = 0;

  for (let i = 0; i + groupings[0] <= pattern.length; i++) {
    // check if there are any #'s in front, break if so
    if (pattern.slice(0, i).includes("#")) break;
    let validArrangements = 1;
    // increment i until we find a group of #'s and ?'s
    // that fits our grouping size
    while (
      i + groupings[0] <= pattern.length &&
      /\./.test(pattern.slice(i, i + groupings[0]))
    ) {
      i++;
    }
    // check if there are any #'s in front, break if so
    if (pattern.slice(0, i).includes("#")) break;
    // check if there's a # immediately before or after the group
    if (i + groupings[0] < pattern.length && pattern[i + groupings[0]] === "#")
      continue;
    // check if out of bounds
    if (i + groupings[0] > pattern.length) break;
    if (groupings.length > 1) {
      // recur if there are groupings left
      validArrangements *= countArrangements(
        pattern.slice(i + groupings[0] + 1),
        groupings.slice(1)
      );
    } else if (pattern.slice(i + groupings[0] + 1).includes("#")) {
      // if no groupings left, check if there are any #'s left unaccounted for
      continue;
    }
    count += validArrangements;
  }

  memo.set(key, count);
  return count;
}

// console.time("part1");

const part1 = input.reduce(
  (acc, [pattern, groupings]) => acc + countArrangements(pattern, groupings),
  0
);
console.log(part1);
// console.timeEnd("part1");

function countArrangements2(pattern, groupings) {
  const pattern2 = Array(5).fill(pattern).join("?");
  const groupings2 = Array(5).fill(groupings).flat();

  return countArrangements(pattern2, groupings2);
}

// console.time("part2");

const part2 = input.reduce(
  (acc, [pattern, groupings]) => acc + countArrangements2(pattern, groupings),
  0
);
console.log(part2);
// console.timeEnd("part2");

// function countArrangementsBrute(pattern, groupings) {
//   const regex = new RegExp(
//     "^\\.*" + groupings.map((a) => `#{${a}}`).join("\\.+") + "\\.*$"
//   );
//   let permutations = [[]];
//   let count = 0;
//   for (const c of pattern) {
//     if (c === "?") {
//       const newPermutations = [];
//       for (const p of permutations) {
//         newPermutations.push([...p, "."]);
//         newPermutations.push([...p, "#"]);
//       }
//       permutations = newPermutations;
//     } else {
//       permutations.forEach((p) => p.push(c));
//     }
//   }
//   for (const p of permutations) {
//     if (regex.test(p.join(""))) count++;
//   }

//   return count;
// }
