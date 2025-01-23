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

class Group {
  constructor(answers) {
    this.answers = answers.map((a) => a.split(""));
    this.part1 = 0;
    this.part2 = 0;
    this.calculatePart1();
    this.calculatePart2();
  }

  calculatePart1() {
    const tempSet = new Set();
    for (const answer of this.answers) {
      for (const c of answer) {
        tempSet.add(c);
      }
    }
    this.part1 = tempSet.size;
  }

  calculatePart2() {
    const mainSet = new Set(this.answers[0]);
    for (let i = 1; i < this.answers.length; i++) {
      const tempSet = new Set(this.answers[i]);
      for (const v of mainSet.values()) {
        if (!tempSet.has(v)) {
          mainSet.delete(v);
        }
      }
    }
    this.part2 = mainSet.size;
  }
}

const groups = [];
let tempAnswers = [];
let i = 0;

while (i < input.length) {
  if (input[i].length) {
    tempAnswers.push(input[i]);
  } else {
    groups.push(new Group(tempAnswers));
    tempAnswers = [];
  }
  i++;
}

groups.push(new Group(tempAnswers));

const part1 = groups.reduce((a, b) => a + b.part1, 0);
const part2 = groups.reduce((a, b) => a + b.part2, 0);
console.log(part1);
console.log(part2);
