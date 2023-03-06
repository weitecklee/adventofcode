const stacks = new Map();

const fs = require('fs');
const path = require('path');
let input = fs.readFileSync(path.resolve(__dirname, 'input.txt'), 'utf-8', (err, data) => {
  if (err) {
    console.log(err)
  } else {
    return data;
  }
});

input = input.split('\n');

let stackRows = 0;
while (input[stackRows][0] === '[') {
  stackRows++;
}

for (let i = stackRows - 1; i >= 0; i--) {
  for (let j = 1; j < input[i].length; j += 4) {
    if (input[i][j] != ' ') {
      const stack = Math.floor(j / 4);
      if (!stacks.has(stack)) {
        stacks.set(stack, []);
      }
      stacks.get(stack).push(input[i][j]);
    }
  }
}

for (let i = stackRows + 2; i < input.length; i++) {
  const [n, from, to] = input[i].match(/\d+/g);
  for (let i = 0; i < n; i++) {
    stacks.get(to - 1).push(stacks.get(from - 1).pop());
  }
}

let res = '';
for (const [n, stack] of stacks) {
  res += stack.pop();
}
console.log(res);

const stacks2 = new Map();
for (let i = stackRows - 1; i >= 0; i--) {
  for (let j = 1; j < input[i].length; j += 4) {
    if (input[i][j] != ' ') {
      const stack = Math.floor(j / 4);
      if (!stacks2.has(stack)) {
        stacks2.set(stack, []);
      }
      stacks2.get(stack).push(input[i][j]);
    }
  }
}

for (let i = stackRows + 2; i < input.length; i++) {
  const [n, from, to] = input[i].match(/\d+/g);
  stacks2.get(to - 1).push(...stacks2.get(from - 1).slice(stacks2.get(from - 1).length - n));
  stacks2.set(from - 1, stacks2.get(from - 1).slice(0, stacks2.get(from - 1).length - n));
}

let res2 = '';
for (const [n, stack] of stacks2) {
  res2 += stack.pop();
}
console.log(res2);