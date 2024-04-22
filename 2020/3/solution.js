const fs = require('fs');
const path = require('path');

const input = fs.readFileSync(path.join(__dirname, 'input.txt'), 'utf-8', (err, data) => {
  if (err) {
    console.log(err)
  } else {
    return data;
  }
}).split('\n');

const colMax = input[0].length;
const rowMax = input.length;

function countTrees([right, down]) {
  let row = 0;
  let col = 0;
  let trees = 0;
  while (row < rowMax) {
    if (input[row][col] === '#') {
      trees++;
    }
    row += down;
    col += right;
    col %= colMax;
  }

  return trees;
}

const part1 = countTrees([3, 1]);
console.log(part1);

const part2Slopes = [[1, 1], [5, 1], [7, 1], [1, 2]];
const part2 = part1 * part2Slopes.reduce((a, b) => a * countTrees(b), 1);
console.log(part2);