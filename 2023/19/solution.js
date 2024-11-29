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
        const test =
          ruleMatch[2] === ">"
            ? (partRating) => partRating[category] > num
            : (partRating) => partRating[category] < num;
        return {
          rule,
          test,
          next: ruleMatch[4],
        };
      } else {
        return {
          rule,
          test: (partRating) => true,
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
