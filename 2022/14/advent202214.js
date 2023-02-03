const fs = require('fs');

let input = fs.readFileSync('input202214.txt', 'utf-8', (err, data) => {
  if (err) {
    console.log(err)
  } else {
    return data;
  }
});

input = input.split('\n');
let maxy = 0;
let minx = Infinity;
let maxx = 0;
const waterfall = new Set();
for (const line of input) {
  const parse = line.split(' -> ').map((a) => a.split(',').map(Number));
  for (const coord of parse) {
    maxy = Math.max(maxy, coord[1]);
    minx = Math.min(minx, coord[0]);
    maxx = Math.max(maxx, coord[0]);
  }
  for (let i = 1; i < parse.length; i++) {
    if (parse[i][0] < parse[i - 1][0]) {
      for (let x = parse[i][0]; x <= parse[i - 1][0]; x++) {
        const coord = x + ',' + parse[i][1];
        waterfall.add(coord);
      }
    } else if (parse[i][0] > parse[i - 1][0]) {
      for (let x = parse[i][0]; x >= parse[i - 1][0]; x--) {
        const coord = x + ',' + parse[i][1];
        waterfall.add(coord);
      }
    } else if (parse[i][1] < parse[i - 1][1]) {
      for (let y = parse[i][1]; y <= parse[i - 1][1]; y++) {
        const coord = parse[i][0] + ',' + y;
        waterfall.add(coord);
      }
    } else {
      for (let y = parse[i][1]; y >= parse[i - 1][1]; y--) {
        const coord = parse[i][0] + ',' + y;
        waterfall.add(coord);
      }
    }
  }
}

let sand = 0;
let rest = true;
while (rest) {
  let x = 500;
  let y = 0;
  rest = false;
  while (!rest && x >= minx && x <= maxx && y <= maxy) {
    y++;
    if (waterfall.has(x + ',' + y)) {
      x--;
      if (waterfall.has(x + ',' + y)) {
        x += 2;
        if (waterfall.has(x + ',' + y)) {
          waterfall.add((x - 1) + ',' + (y - 1));
          rest = true;
          sand++;
        }
      }
    }
  }
}

console.log(sand);

const ground = maxy + 2;

while (!waterfall.has('500,0')) {
  let x = 500;
  let y = 0;
  rest = false;
  while (!rest && y < ground) {
    y++;
    if (waterfall.has(x + ',' + y)) {
      x--;
      if (waterfall.has(x + ',' + y)) {
        x += 2;
        if (waterfall.has(x + ',' + y)) {
          waterfall.add((x - 1) + ',' + (y - 1));
          rest = true;
          sand++;
        }
      }
    }
  }
  if (y === ground) {
    waterfall.add(x + ',' + y);
  }
}

console.log(sand);