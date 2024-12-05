const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n\n");

const rules = input[0].split("\n").map((a) => a.split("|").map(Number));
const updates = input[1].split("\n").map((a) => a.split(",").map(Number));

class Rule {
  constructor(num) {
    this.num = num;
    this.afterNum = new Set();
  }

  addRule(n) {
    this.afterNum.add(n);
  }
}

const ruleMap = new Map();

for (const [n1, n2] of rules) {
  if (!ruleMap.has(n1)) {
    ruleMap.set(n1, new Rule(n1));
  }
  ruleMap.get(n1).addRule(n2);
}

function sortingFunction(a, b) {
  if (ruleMap.has(a) && ruleMap.get(a).afterNum.has(b)) return -1;
  if (ruleMap.has(b) && ruleMap.get(b).afterNum.has(a)) return 1;
  return 0;
}

let part1 = 0;
let part2 = 0;

for (const update of updates) {
  const order = new Map();
  for (let i = 0; i < update.length; i++) {
    order.set(update[i], i);
  }
  let isCorrectOrder = true;
  for (let i = 0; i < update.length; i++) {
    if (ruleMap.has(update[i])) {
      const rule = ruleMap.get(update[i]);
      for (const n of rule.afterNum) {
        if (order.has(n) && order.get(n) < i) {
          isCorrectOrder = false;
          break;
        }
      }
    }
    if (!isCorrectOrder) break;
  }
  if (isCorrectOrder) {
    part1 += update[Math.floor(update.length / 2)];
  } else {
    update.sort(sortingFunction);
    part2 += update[Math.floor(update.length / 2)];
  }
}

console.log(part1);
console.log(part2);
