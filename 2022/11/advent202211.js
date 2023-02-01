const monkeys = [];

function Monkey() {
  this.items = null;
  this.operation = null;
  this.test = null;
  this.ifTrue = null;
  this.ifFalse = null;
  this.inspected = 0;
};
const fs = require('fs');

let input = fs.readFileSync('input202211.txt', 'utf-8', (err, data) => {
  if (err) {
    console.log(err)
  } else {
    return data;
  }
});

input = input.split('\n');

let prime = 1;

let curr;
for (const line of input) {
  if (/Monkey/.test(line)) {
    curr = new Monkey();
    monkeys.push(curr);
  } else if (/Starting/.test(line)) {
    const items = line.match(/\d+/g).map(Number);
    curr.items = items;
  } else if (/Operation/.test(line)) {
    const parse = line.match(/old (.) (.+)/);
    const oper = parse[1];
    const arg = parse[2];
    if (arg === 'old') {
      if (oper === '+') {
        curr.operation = (a) => (a + a);
      } else {
        curr.operation = (a) => (a * a);
      }
    } else {
      if (oper === '+') {
        curr.operation = (a) => (a + Number(arg));
      } else {
        curr.operation = (a) => (a * Number(arg));
      }
    }
  } else if (/Test/.test(line)) {
    const divisor = Number(line.match(/\d+/)[0]);
    prime *= divisor;
    curr.test = (a) => (a % divisor === 0);
  } else if (/If true/.test(line)) {
    const ifTrue = Number(line.match(/\d+/)[0]);
    curr.ifTrue = ifTrue;
  } else if (/If false/.test(line)) {
    const ifFalse = Number(line.match(/\d+/)[0]);
    curr.ifFalse = ifFalse;
  }
}

for (let i = 0; i < 10000; i++) {
  for (const monkey of monkeys) {
    for (let item of monkey.items) {
      item = monkey.operation(item) % prime;
      if (monkey.test(item)) {
        monkeys[monkey.ifTrue].items.push(item);
      } else {
        monkeys[monkey.ifFalse].items.push(item);
      }
      monkey.inspected++;
    }
    monkey.items = [];
  }
}

const inspected = [];
for (const monkey of monkeys) {
  inspected.push(monkey.inspected);
}
inspected.sort((a, b) => b - a);
console.log(inspected[0] * inspected[1]);

