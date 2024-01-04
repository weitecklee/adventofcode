const fs = require('fs');

const input = fs.readFileSync('input.txt', 'utf-8', (err, data) => {
  if (err) {
    console.log(err)
  } else {
    return data;
  }
}).split('\n');


const instructions = [];

for (const line of input) {
  const parts = line.split(' ');
  for (let i = 0; i < parts.length; i++) {
    const n = Number(parts[i]);
    if (!isNaN(n)) {
      parts[i] = n;
    }
  }
  instructions.push(parts);
}

const execute = (c) => {
  const register = new Map();
  register.set('a', 0);
  register.set('b', 0);
  register.set('c', c);
  register.set('d', 0);
  let i = 0;
  while (i < input.length) {
    const instruction = instructions[i];
    if (instruction[0] === 'cpy') {
      if (typeof instruction[1] === 'string') {
        register.set(instruction[2], register.get(instruction[1]));
      } else {
        register.set(instruction[2], instruction[1]);
      }
    } else if (instruction[0] === 'inc') {
      register.set(instruction[1], register.get(instruction[1]) + 1);
    } else if (instruction[0] === 'dec') {
      register.set(instruction[1], register.get(instruction[1]) - 1);
    } else {
      let check;
      if (typeof instruction[1] === 'string') {
        check = register.get(instruction[1]);
      } else {
        check = instruction[1];
      }
      if (check !== 0) {
        i += instruction[2] - 1;
      }
    }
    i++;
  }
  return register.get('a');
}

console.log(execute(0));
console.log(execute(1));