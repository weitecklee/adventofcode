const fs = require('fs');
const path = require('path');

const input = fs.readFileSync(path.join(__dirname, 'input.txt'), 'utf-8', (err, data) => {
  if (err) {
    console.log(err)
  } else {
    return data;
  }
}).split('\n');

function convertString(str) {
  const [n, chem] = str.split(' ');
  return [chem, Number(n)];
}

class DefaultDict {
  constructor() {
    this.map = new Map();
  }

  has(key) {
    return this.map.has(key);
  }

  get(key) {
    if (!this.map.has(key)) {
      this.map.set(key, 0);
    }
    return this.map.get(key);
  }

  set(key, value) {
    this.map.set(key, value);
  }
}

function Recipe(quantity, inputs) {
  this.quantity = quantity;
  this.inputs = inputs.map(convertString);
}

function parseInput(input) {
  const recipes = new Map();

  for (const line of input) {
    const [inputs, output] = line.split(' => ');
    const [chemical, quantity] = convertString(output);
    recipes.set(chemical, new Recipe(quantity, inputs.split(', ')))
  }

  return recipes;
}

function calculateOreReq(chemical, quantity, recipes, surplus) {
  if (chemical === 'ORE') {
    return quantity;
  }

  const available = surplus.get(chemical);
  const used = Math.min(surplus.get(chemical), quantity);
  surplus.set(chemical, available - used);
  quantity -= used;
  if (quantity === 0) {
    return 0;
  }

  const recipe = recipes.get(chemical);
  const batches = Math.ceil(quantity / recipe.quantity);
  surplus.set(chemical, surplus.get(chemical) + batches * recipe.quantity - quantity);
  let ore = 0;

  for (const [inputChemical, inputQuantity] of recipe.inputs) {
    ore += calculateOreReq(inputChemical, inputQuantity * batches, recipes, surplus);
  }

  return ore;
}

function part1(recipes) {
  const surplus = new DefaultDict();
  return calculateOreReq('FUEL', 1, recipes, surplus);
}

function part2(recipes, oreAvailable) {
  let lo = 0;
  let hi = oreAvailable;

  while (lo < hi) {
    const mid = Math.floor((lo + hi + 1) / 2);
    const surplus = new DefaultDict();
    const ore = calculateOreReq('FUEL', mid, recipes, surplus);
    if (ore > oreAvailable) {
      hi = mid - 1;
    } else {
      lo = mid;
    }
  }

  return lo;
}

const recipes = parseInput(input);
console.log(part1(recipes));
console.log(part2(recipes, 1000000000000));
