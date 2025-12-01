const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split(": "));

const monkeys = new Map();

class Monkey {
  constructor(name, job) {
    this.name = name;
    this.type = "number";
    this.number;
    this.monkey1;
    this.monkey2;
    if (/\d/.test(job)) {
      this.number = Number(job);
    } else {
      const parts = job.split(" ");
      this.monkey1 = parts[0];
      this.type = parts[1];
      this.monkey2 = parts[2];
    }
  }

  yell() {
    switch (this.type) {
      case "number":
        return this.number;
      case "+":
        return (
          monkeys.get(this.monkey1).yell() + monkeys.get(this.monkey2).yell()
        );
      case "-":
        return (
          monkeys.get(this.monkey1).yell() - monkeys.get(this.monkey2).yell()
        );
      case "*":
        return (
          monkeys.get(this.monkey1).yell() * monkeys.get(this.monkey2).yell()
        );
      case "/":
        return (
          monkeys.get(this.monkey1).yell() / monkeys.get(this.monkey2).yell()
        );
    }
  }
}

for (const [name, job] of input) {
  monkeys.set(name, new Monkey(name, job));
}

console.log(monkeys.get("root").yell());

monkeys.get("root").type = "-";
let lo = 0;
let hi = 331120084396440 * 2;

while (lo < hi) {
  const mid = Math.floor(lo + (hi - lo) / 2);
  monkeys.get("humn").number = mid;
  const res = monkeys.get("root").yell();
  if (res === 0) {
    console.log(mid);
    break;
  }
  if (res > 0) {
    lo = mid + 1;
  } else {
    hi = mid - 1;
  }
}
