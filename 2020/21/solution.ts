import * as fs from "fs";
import * as path from "path";

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n");

class Ingredient {
  name: string;
  appearances: number;

  constructor(name: string) {
    this.name = name;
    this.appearances = 0;
  }
}

class Allergen {
  name: string;
  ingredientCandidates: Set<string>;
  candidateSets: Set<string>[];
  ingredient: string;

  constructor(name: string) {
    this.name = name;
    this.ingredientCandidates = new Set();
    this.candidateSets = [];
    this.ingredient = "";
  }
}

const ingredientMap: Map<string, Ingredient> = new Map();
const allergenMap: Map<string, Allergen> = new Map();

const foodRegex = /^(.+) \(contains (.+)\)$/;

for (const line of input) {
  const match = line.match(foodRegex);
  if (match) {
    const ingreds = new Set(match[1].split(" "));
    const allergs = match[2].split(", ");
    for (const allerg of allergs) {
      if (!allergenMap.has(allerg))
        allergenMap.set(allerg, new Allergen(allerg));
      const allergen = allergenMap.get(allerg)!;
      allergen.candidateSets.push(ingreds);
    }
    for (const ingred of ingreds) {
      if (!ingredientMap.has(ingred))
        ingredientMap.set(ingred, new Ingredient(ingred));
      const ingredient = ingredientMap.get(ingred)!;
      ingredient.appearances++;
    }
  }
}

const candidates: Set<string> = new Set();

for (const [_, allergen] of allergenMap) {
  allergen.ingredientCandidates = new Set(allergen.candidateSets[0]);
  for (let i = 1; i < allergen.candidateSets.length; i++) {
    for (const candidate of allergen.ingredientCandidates) {
      if (!allergen.candidateSets[i].has(candidate))
        allergen.ingredientCandidates.delete(candidate);
    }
  }
  for (const candidate of allergen.ingredientCandidates) {
    candidates.add(candidate);
  }
}

let part1 = 0;
for (const [name, ingredient] of ingredientMap) {
  if (!candidates.has(name)) part1 += ingredient.appearances;
}

console.log(part1);

while (candidates.size) {
  for (const [_, allergen] of allergenMap) {
    if (allergen.ingredient) continue;
    if (allergen.ingredientCandidates.size === 1) {
      allergen.ingredient = allergen.ingredientCandidates
        .values()
        .next().value!;
      candidates.delete(allergen.ingredient);
    } else {
      for (const candidate of allergen.ingredientCandidates) {
        if (!candidates.has(candidate))
          allergen.ingredientCandidates.delete(candidate);
      }
    }
  }
}

console.log(
  Array.from(allergenMap.values())
    .sort((a, b) => {
      if (a.name < b.name) return -1;
      return 1;
    })
    .map((a) => a.ingredient)
    .join(",")
);
