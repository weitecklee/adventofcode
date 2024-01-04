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

function Dir(parent) {
  this.parent = parent;
  this.subdirs = new Map();
  this.size = 0;
}

const main = new Dir(null);
let curr = main;

for (const line of input) {
  if (line[0] === '$' && line !== '$ cd \/') { // line starting with $
    const path = /\$ cd (.*)/.exec(line); // line starting with $ cd, others ignored
    if (path) {
      if (path[1] === '..') {
        curr = curr.parent;
      } else {
        curr = curr.subdirs.get(path[1]);
      }
    }
  } else if (/\d/.test(line[0])) { // line starting with number
    const size = /^\d+/.exec(line)[0];
    curr.size += Number(size);
  } else { // line starting with dir
    const path = /dir (.*)/.exec(line);
    if (path) {
      const newDir = new Dir(curr);
      curr.subdirs.set(path[1], newDir);
    }
  }
}

let sum = 0;
const recur = (dir) => {
  let size = dir.size;
  for (const [path, subdir] of dir.subdirs) {
    size += recur(subdir);
  }
  if (size <= 100000) {
    sum += size;
  }
  return size;
}

const total = recur(main);
console.log(sum);

const spaceToBeFreed = total - 40000000;
let min = total;
const recur2 = (dir) => {
  let size = dir.size;
  for (const [path, subdir] of dir.subdirs) {
    size += recur2(subdir);
  }
  if (size >= spaceToBeFreed && size <= min) {
    min = size;
  }
  return size;
}
recur2(main);
console.log(min);
