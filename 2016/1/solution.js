const fs = require('fs');
const path = require('path');

const input = fs.readFileSync(path.join(__dirname, 'input.txt'), 'utf-8', (err, data) => {
  if (err) {
    console.log(err)
  } else {
    return data;
  }
});

const pos = [0, 0];
const dir = [0, 1];

const steps = input.split(', ')

class Instruction {
  constructor(s) {
    this.turn = s[0];
    this.dist = Number(s.slice(1));
  }
}

const instructions = [];

for (const step of steps) {
  instructions.push(new Instruction(step));
}

for (const instruction of instructions) {
  if (instruction.turn === "L") {
    [dir[0], dir[1]] = [-dir[1], dir[0]];
  } else {
    [dir[0], dir[1]] = [dir[1], -dir[0]];
  }
  pos[0] += instruction.dist * dir[0];
  pos[1] += instruction.dist * dir[1];
}

console.log(Math.abs(pos[0]) + Math.abs(pos[1]))

const seen = new Set();
dir[0] = 0;
dir[1] = 1;
pos[0] = 0;
pos[1] = 0;

const coordinate = (p) => (p[0] + ',' + p[1]);
seen.add(coordinate(pos));

for (const instruction of instructions) {
  if (instruction.turn === "L") {
    [dir[0], dir[1]] = [-dir[1], dir[0]];
  } else {
    [dir[0], dir[1]] = [dir[1], -dir[0]];
  }
  for (let i = 0; i < instruction.dist; i++) {
    pos[0] += dir[0];
    pos[1] += dir[1];
    const coord = coordinate(pos);
    if (seen.has(coord)) {
      console.log(Math.abs(pos[0]) + Math.abs(pos[1]));
      return;
    }
    seen.add(coord);
  }
}