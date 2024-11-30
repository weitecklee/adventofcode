const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n\n");

const workflowRegex = /(.+){(.+)}/;
const ruleRegex = /(\w+)([><])(\d+):(\w+)/;

class Workflow {
  constructor(line) {
    const regexMatch = line.match(workflowRegex);
    this.name = regexMatch[1];
    this.rules = regexMatch[2].split(",").map((rule) => {
      const ruleMatch = rule.match(ruleRegex);
      if (ruleMatch) {
        const category = ruleMatch[1];
        const num = Number(ruleMatch[3]);
        const sign = ruleMatch[2];
        const test =
          sign === ">"
            ? (part) => part[category] > num
            : (part) => part[category] < num;
        return {
          category,
          num,
          rule,
          sign,
          test,
          next: ruleMatch[4],
        };
      } else {
        return {
          category: "?",
          num: 0,
          sign: "=",
          rule,
          test: (part) => true,
          next: rule,
        };
      }
    });
  }

  evaluatePart(part) {
    for (const rule of this.rules) {
      if (rule.test(part)) {
        return rule.next;
      }
    }
  }
}

const workflowMap = new Map();
input[0].split("\n").forEach((line) => {
  const workflow = new Workflow(line);
  workflowMap.set(workflow.name, workflow);
});
const ratings = input[1].split("\n").map((line) => {
  return line
    .slice(1, -1)
    .split(",")
    .reduce((acc, part) => {
      const [category, num] = part.split("=");
      acc[category] = Number(num);
      return acc;
    }, {});
});

let part1 = 0;
for (const rating of ratings) {
  let curr = "in";
  while (curr !== "A" && curr !== "R") {
    curr = workflowMap.get(curr).evaluatePart(rating);
  }
  if (curr === "A") {
    part1 += Object.values(rating).reduce((acc, val) => acc + val, 0);
  }
}
console.log(part1);

function reverseRule(rule) {
  const newRule = { ...rule };
  newRule.sign = rule.sign === ">" ? "<" : ">";
  newRule.rule = rule.rule.replace(rule.sign, newRule.sign);
  newRule.num = rule.num + (rule.sign === ">" ? 1 : -1);
  newRule.test =
    rule.sign === ">"
      ? (part) => part[rule.category] < newRule.num
      : (part) => part[rule.category] > newRule.num;
  return newRule;
}

const queue = [{ rules: [], currentWorkflow: workflowMap.get("in") }];
const combinations = [];
while (queue.length) {
  const { rules, currentWorkflow } = queue.pop();
  // for each rule, reverse all rules before it and add current rule
  for (let i = 0; i < currentWorkflow.rules.length; i++) {
    const currentRules = rules.slice();
    for (let j = 0; j < i; j++) {
      currentRules.push(reverseRule(currentWorkflow.rules[j]));
    }
    currentRules.push(currentWorkflow.rules[i]);
    // if next is "A", add to combinations
    // if next is "R", do nothing (discard)
    // otherwise, add to queue
    if (currentWorkflow.rules[i].next === "A") {
      combinations.push(currentRules);
    } else if (currentWorkflow.rules[i].next !== "R") {
      queue.push({
        rules: currentRules,
        currentWorkflow: workflowMap.get(currentWorkflow.rules[i].next),
      });
    }
  }
}

let part2 = 0;
for (const combo of combinations) {
  const approved = { x: [1, 4000], m: [1, 4000], a: [1, 4000], s: [1, 4000] };
  for (const rule of combo) {
    if (rule.sign === "=") continue;
    if (rule.sign === ">") {
      approved[rule.category][0] = Math.max(
        approved[rule.category][0],
        rule.num + 1
      );
    } else {
      approved[rule.category][1] = Math.min(
        approved[rule.category][1],
        rule.num - 1
      );
    }
  }
  if (!Object.values(approved).every(([min, max]) => min <= max)) continue;

  part2 += Object.values(approved).reduce(
    (acc, [min, max]) => acc * (max - min + 1),
    1
  );
}

console.log(part2);
