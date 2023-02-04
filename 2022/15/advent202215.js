const fs = require('fs');

let input = fs.readFileSync('input202215.txt', 'utf-8', (err, data) => {
  if (err) {
    console.log(err)
  } else {
    return data;
  }
});

input = input.split('\n');

const y = 2000000;

const arr = [];
for (const line of input) {
  const pos = line.match(/-?\d+/g).map(Number);
  const mdist = Math.abs(pos[0] - pos[2]) + Math.abs(pos[1] - pos[3]);
  const dist = Math.abs(y - pos[1]);
  if (dist <= mdist) {
    if (y === pos[1]) {
      arr.push([pos[0] - (mdist - dist), pos[0] - 1]);
      arr.push([pos[0] + 1, pos[0] + (mdist - dist)]);
    } else if (y === pos[3]) {
      arr.push([pos[0] - (mdist - dist), pos[2] - 1]);
      arr.push([pos[2] + 1, pos[0] + (mdist - dist)]);
    } else {
      arr.push([pos[0] - (mdist - dist), pos[0] + (mdist - dist)]);
    }
  }
}

arr.sort((a, b) => a[0] - b[0]);
let ans = 0;
let tmp = arr[0];
for (i = 1; i < arr.length; i++) {
  if (arr[i][0] > tmp[1]) {
    ans += tmp[1] - tmp[0] + 1;
    tmp = arr[i];
  } else {
    tmp[1] = Math.max(tmp[1], arr[i][1]);
  }
}
ans += tmp[1] - tmp[0] + 1;
console.log(ans);

const max = 4000000;
const pos = input.map((line) => line.match(/-?\d+/g).map(Number))

for (let j = 0; j <= max; j++) {
  const arr2 = [];
  for (const line of pos) {
    const mdist = Math.abs(line[0] - line[2]) + Math.abs(line[1] - line[3]);
    const dist = Math.abs(j - line[1]);
    if (dist <= mdist) {
      arr2.push([line[0] - (mdist - dist), line[0] + (mdist - dist)]);
    }
  }

  arr2.sort((a, b) => a[0] - b[0]);
  let tmp = arr2[0];
  let k = 1;
  while (k < arr2.length) {
    if (arr2[k][0] > tmp[1]) {
      // console.log(j);
      // console.log(tmp);
      // console.log(arr2[k]);
      // console.log('Tuning Frequency: ');
      console.log(4000000 * (tmp[1] + 1) + j);
      return;
    } else {
      tmp[1] = Math.max(tmp[1], arr2[k][1]);
    }
    k++;
  }
}
