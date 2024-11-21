const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n");

const sequence = input[0].split("").map((a) => (a === "L" ? "left" : "right"));

function Node(left, right) {
  this.left = left;
  this.right = right;
}

const nodeMap = new Map();

for (let i = 2; i < input.length; i++) {
  const [node, left, right] = input[i].match(/\w+/g);
  nodeMap.set(node, new Node(left, right));
}

let part1 = 0;
let curr = "AAA";

while (curr != "ZZZ") {
  curr = nodeMap.get(curr)[sequence[part1 % sequence.length]];
  part1++;
}

console.log(part1);

const currentNodes = Array.from(nodeMap.keys()).filter((a) => a[2] === "A");

function findLoops(node, sequence) {
  let steps = 0;
  let curr = node;
  while (curr[2] !== "Z") {
    curr = nodeMap.get(curr)[sequence[steps % sequence.length]];
    steps++;
  }
  const loopNode = curr;
  const stepsToZ = steps;
  do {
    curr = nodeMap.get(curr)[sequence[steps % sequence.length]];
    steps++;
  } while (curr != loopNode);
  // return steps to node ending with Z, then steps to loop back to that node
  return [stepsToZ, steps - stepsToZ];
}

const stepsToZ = [];
const loops = [];

for (const node of currentNodes) {
  const [a, b] = findLoops(node, sequence);
  stepsToZ.push(a);
  loops.push(b);
}

// console.log(stepsToZ);
// console.log(loops);

// Examination of stepsToZ and loops arrays show that stepsToZ[i] = loops[i], so the answer is just finding lcm of either array
function gcd(a, b) {
  if (b === 0) return a;
  return gcd(b, a % b);
}
function lcm(a, b) {
  return (a * b) / gcd(a, b);
}

const part2 = stepsToZ.reduce(lcm);
console.log(part2);
