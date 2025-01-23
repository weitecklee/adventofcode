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

const targetColor = "shiny gold";
const bags = new Map();

class Bag {
  constructor(color, containedBags) {
    this.color = color;
    this.containedBags = containedBags;
    this.checked = false;
    this.containsTarget = false;
    this.numberOfBagsInside = 0;
  }

  check() {
    if (this.checked) {
      return;
    }
    for (const [containedBagColor, containedBagCount] of this.containedBags) {
      const currentBag = bags.get(containedBagColor);
      if (!currentBag.checked) {
        currentBag.check();
      }

      if (containedBagColor === targetColor || currentBag.containsTarget) {
        this.containsTarget = true;
      }
      this.numberOfBagsInside +=
        containedBagCount * (currentBag.numberOfBagsInside + 1);
    }
    this.checked = true;
  }
}

const colorPattern = /^.*?(?= bag)/;
const bagPattern = /(\d+) (.*?) bag/g;

for (const line of input) {
  const color = line.match(colorPattern)[0];
  const bagMatches = line.matchAll(bagPattern);
  const tempBags = new Map();

  for (const bagMatch of bagMatches) {
    tempBags.set(bagMatch[2], Number(bagMatch[1]));
  }

  bags.set(color, new Bag(color, tempBags));
}

let part1 = 0;
for (const [color, bag] of bags) {
  bag.check();
  part1 += bag.containsTarget;
}

console.log(part1);
const targetBag = bags.get(targetColor);
console.log(targetBag.numberOfBagsInside);
