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

const draws = input[0].split(',').map(Number);

const boards = new Set();

const Board = function() {
  this.spaces = new Map();
  this.rows = new Array(5).fill(0);
  this.cols = new Array(5).fill(0);
};

for (let i = 2; i < input.length; i += 6) {
  const board = new Board();
  for (let r = 0; r < 5; r++) {
    const line = input[i + r].match(/\d+/g).map(Number);
    for (let c = 0; c < 5; c++) {
      board.spaces.set(line[c], [r, c]);
    }
  }
  boards.add(board);
}

let bingo = false;
let winningBoard = null;
let partOneDone = false;

for (const draw of draws) {
  for (const board of boards) {
    if (board.spaces.has(draw)) {
      const [r, c] = board.spaces.get(draw);
      board.spaces.delete(draw);
      board.rows[r]++;
      board.cols[c]++;
      if (board.rows[r] === 5 || board.cols[c] === 5) {
        bingo = true;
        winningBoard = board;
        if (boards.size > 1) {
          boards.delete(board);
        } else {
          let sum = 0;
          for (const [n, space] of winningBoard.spaces) {
            sum += n;
          }
          console.log(sum * draw);
          return;
        }
      }
    }
  }
  if (bingo && !partOneDone) {
    let sum = 0;
    for (const [n, space] of winningBoard.spaces) {
      sum += n;
    }
    console.log(sum * draw);
    partOneDone = true;
  }
}


