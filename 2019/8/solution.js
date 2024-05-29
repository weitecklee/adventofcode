const fs = require('fs');
const path = require('path');

const input = fs.readFileSync(path.join(__dirname, 'input.txt'), 'utf-8', (err, data) => {
  if (err) {
    console.log(err)
  } else {
    return data;
  }
});

const width = 25;
const height = 6;
const image = [];

for (let i = 0; i < input.length; i += width * height) {
  const layer = [];
  for (let j = 0; j < height; j++) {
    layer.push(input.slice(i + width * j, i + (width) * (j + 1)));
  }
  image.push(layer);
}

let minZeros = Infinity;
let part1 = 0;

for (const layer of image) {
  let count0 = 0;
  let count1 = 0;
  let count2 = 0;
  for (const row of layer) {
    for (const pixel of row) {
      switch (pixel) {
        case '0':
          count0++;
          break;
        case '1':
          count1++;
          break;
        case '2':
          count2++;
      }
    }
  }
  if (count0 < minZeros) {
    minZeros = count0;
    part1 = count1 * count2;
  }
}

console.log(part1);

const message = [];
for (let row = 0; row < height; row++) {
  const line = [];
  for (let col = 0; col < width; col++) {
    for (const layer of image) {
      if (layer[row][col] === '1') {
        line.push('@');
        break;
      }
      if (layer[row][col] === '0') {
        line.push(' ');
        break;
      }
    }
  }
  message.push(line.join(''));
}

for (const line of message) {
  console.log(line);
}