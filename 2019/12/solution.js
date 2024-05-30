const fs = require('fs');
const path = require('path');

const input = fs.readFileSync(path.join(__dirname, 'input.txt'), 'utf-8', (err, data) => {
  if (err) {
    console.log(err)
  } else {
    return data;
  }
}).split('\n');

const posStringRegex = /<x=(.+), y=(.+), z=(.+)>/;

class Moon {
  constructor(posString) {
    const match = posString.match(posStringRegex);
    this.pos = [Number(match[1]), Number(match[2]), Number(match[3])];
    this.vel = [0, 0, 0];
  }

  move() {
    for (let i = 0; i < 3; i++) {
      this.pos[i] += this.vel[i];
    }
  }

  gravitate(moon) {
    for (let i = 0; i < 3; i++) {
      const diff = this.pos[i] - moon.pos[i];
      if (diff > 0) {
        this.vel[i]--;
        moon.vel[i]++;
      } else if (diff < 0) {
        this.vel[i]++;
        moon.vel[i]--;
      }
    }
  }

  calculateEnergy() {
    return this.pos.reduce((a, b) => a + Math.abs(b), 0) * this.vel.reduce((a, b) => a + Math.abs(b), 0);
  }
}

function moveMoons(moons) {
  for (const moon of moons) {
    moon.move();
  }
}

function simulateMoons(moons) {
  for (let i = 0; i < moons.length; i++) {
    for (let j = i + 1; j < moons.length; j++) {
      moons[i].gravitate(moons[j]);
    }
  }
  moveMoons(moons);
}

function part1(input) {
  const moons = input.map((a) => new Moon(a));

  for (let i = 0; i < 1000; i++) {
    simulateMoons(moons);
  }

  return moons.reduce((a, b) => a + b.calculateEnergy(), 0);

}

function printStates(moons) {
  const res = [];
  for (let i = 0; i < 3; i++) {
    const state = [];
    for (const moon of moons) {
      state.push(moon.pos[i], moon.vel[i]);
    }
    res.push(state.join(','));
  }
  return res;
}

function gcd(a, b) {
  while (b != 0) {
    const tmp = b;
    b = a % b;
    a = tmp;
  }
  return a;
}

function lcm(a, b) {
  return a * b / gcd(a, b);
}

function part2(input) {
  const moons = input.map((a) => new Moon(a));
  const repeated = [null, null, null];
  const states = [new Map(), new Map(), new Map()];
  let step = 0;
  while (repeated.some((a) => !a)) {
    const currStates = printStates(moons);
    for (let i = 0; i < 3; i++) {
      if (repeated[i]) {
        continue;
      }
      if (states[i].has(currStates[i])) {
        repeated[i] = [states[i].get(currStates[i]), step];
      } else {
        states[i].set(currStates[i], step);
      }
    }
    simulateMoons(moons);
    step++;
  }
  return repeated.reduce((a, b) => lcm(a, b[1]), 1);
}

console.log(part1(input));
console.log(part2(input));