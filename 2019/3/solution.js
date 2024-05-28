const fs = require('fs');
const path = require('path');

const input = fs.readFileSync(path.join(__dirname, 'input.txt'), 'utf-8', (err, data) => {
  if (err) {
    console.log(err)
  } else {
    return data;
  }
}).split('\n').map((a) => a.split(','));

const pos2String = (pos) => pos[0] + ',' + pos[1];

class Wire {
  constructor(input) {
    this.input = input;
    this.set = new Set();
    this.steps = new Map();
    this.mapWire();
  }

  mapWire() {
    const pos = [0, 0];
    let step = 0;
    for (const part of this.input) {
      const direction = part[0];
      const distance = Number(part.slice(1));
      const increment = [0, 0];
      switch (direction) {
        case 'U':
          increment[1] = -1;
          break;
        case 'D':
          increment[1] = 1;
          break;
        case 'L':
          increment[0] = -1;
          break;
        case 'R':
          increment[0] = 1;
      }
      for (let i = 0; i < distance; i++) {
        pos[0] += increment[0];
        pos[1] += increment[1];
        step++;
        const posString = pos2String(pos);
        this.set.add(posString);
        if (!this.steps.has(posString)) {
          this.steps.set(posString, step);
        }
      }
    }
  }
}

const wire1 = new Wire(input[0]);
const wire2 = new Wire(input[1]);

const intersections = wire1.set.intersection(wire2.set);

let part1 = Infinity;
let part2 = Infinity;

for (const pos of intersections) {
  const manhattanDistance = pos.split(',').map(Number).map(Math.abs).reduce((a, b) => a + b);
  const combinedSteps = wire1.steps.get(pos) + wire2.steps.get(pos);
  part1 = Math.min(part1, manhattanDistance);
  part2 = Math.min(part2, combinedSteps);
}

console.log(part1);
console.log(part2);