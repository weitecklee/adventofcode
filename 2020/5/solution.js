const fs = require('fs');
const path = require('path');

const input = fs.readFileSync(path.join(__dirname, 'input.txt'), 'utf-8', (err, data) => {
  if (err) {
    console.log(err)
  } else {
    return data;
  }
}).split('\n');

function convertRow(row) {
  return parseInt(row.replaceAll('F', '0').replaceAll('B', '1'), 2);
}

function convertCol(col) {
  return parseInt(col.replaceAll('L', '0').replaceAll('R', '1'), 2);
}

const seats = new Set();
let part1 = 0;
for (const line of input) {
  const rowNum = convertRow(line.slice(0, 7));
  const colNum = convertCol(line.slice(7));
  const seatID = rowNum * 8 + colNum;
  part1 = Math.max(part1, seatID);
  seats.add(seatID);
}
console.log(part1);

for (const seat of seats) {
  if (seats.has(seat + 2) && !(seats.has(seat + 1))) {
    console.log(seat + 1);
    break;
  }
  if (seats.has(seat - 2) && !(seats.has(seat - 1))) {
    console.log(seat - 1);
    break;
  }
}