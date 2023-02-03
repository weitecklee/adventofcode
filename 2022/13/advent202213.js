const fs = require('fs');

let input = fs.readFileSync('input202213.txt', 'utf-8', (err, data) => {
  if (err) {
    console.log(err)
  } else {
    return data;
  }
});

input = input.split('\n');

const check = (a, b) => {
  if (typeof a === 'number' && typeof b === 'number') {
    if (a < b) {
      return true;
    }
    if (a > b) {
      return false;
    }
    return;
  }
  if (typeof a === 'number') {
    return check([a], b);
  }
  if (typeof b === 'number') {
    return check(a, [b]);
  }
  for (let i = 0; i < a.length; i++) {
    if (b[i] === undefined) {
      return false;
    }
    const c = check(a[i], b[i]);
    if (c !== undefined) {
      return c;
    }
  }
  if (b.length > a.length) {
    return true;
  }
  return;
}

let ans = 0;
const input2 = [];
for (let i = 0; i < input.length; i += 3) {
  const left = JSON.parse(input[i]);
  const right = JSON.parse(input[i + 1]);
  input2.push(left, right);
  if (check(left, right)) {
    ans += i / 3 + 1;
  }
}

console.log(ans);

input2.push([[2]]);
input2.push([[6]]);

input2.sort((a, b) => {
  if (check(a, b)) {
    return -1;
  }
  return 1;
})

let decoder = 1;
for (let i = 0; i < input2.length; i++) {
  if (input2[i].length === 1 && input2[i][0].length === 1 && (input2[i][0][0] === 2 || input2[i][0][0] === 6)) {
    decoder *= i + 1;
  }
}

console.log(decoder);