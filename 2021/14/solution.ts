import * as fs from "fs";
import * as path from "path";

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n\n");

class Polymer {
  elementCounts: Map<string, number>;
  pairCounts: Map<string, number>;
  constructor(template: string) {
    this.elementCounts = new Map();
    this.pairCounts = new Map();

    for (const elem of template) {
      this.elementCounts.set(elem, (this.elementCounts.get(elem) || 0) + 1);
    }

    for (let i = 1; i < template.length; i++) {
      const pair = template.slice(i - 1, i + 1);
      this.pairCounts.set(pair, (this.pairCounts.get(pair) || 0) + 1);
    }
  }

  polymerize(rules: string[][]) {
    const elementCounts2: Map<string, number> = new Map(this.elementCounts);
    const pairCounts2: Map<string, number> = new Map();
    for (const [pair, newElem] of rules) {
      if (this.pairCounts.has(pair)) {
        elementCounts2.set(
          newElem,
          (elementCounts2.get(newElem) || 0) + this.pairCounts.get(pair)!
        );
        const newPair1 = pair[0] + newElem;
        const newPair2 = newElem + pair[1];
        pairCounts2.set(
          newPair1,
          (pairCounts2.get(newPair1) || 0) + this.pairCounts.get(pair)!
        );
        pairCounts2.set(
          newPair2,
          (pairCounts2.get(newPair2) || 0) + this.pairCounts.get(pair)!
        );
      }
    }

    this.elementCounts = elementCounts2;
    this.pairCounts = pairCounts2;
  }

  get puzzleSolution(): number {
    const maxElemQuantity = Math.max(...this.elementCounts.values());
    const minElemQuantity = Math.min(...this.elementCounts.values());
    return maxElemQuantity - minElemQuantity;
  }
}

const polymer = new Polymer(input[0]);

const insertionRules = input[1].split("\n").map((a) => a.split(" -> "));

for (let i = 0; i < 10; i++) {
  polymer.polymerize(insertionRules);
}

console.log(polymer.puzzleSolution);

for (let i = 10; i < 40; i++) {
  polymer.polymerize(insertionRules);
}

console.log(polymer.puzzleSolution);
