const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8", (err, data) => {
    if (err) {
      console.log(err);
    } else {
      return data;
    }
  })
  .split("\n");

class Knot {
  constructor() {
    this.pos = [0, 0];
    this.tail = null;
  }
  move(vec) {
    this.pos[0] += vec[0];
    this.pos[1] += vec[1];
    if (this.tail) {
      const dif0 = this.pos[0] - this.tail.pos[0];
      const dif1 = this.pos[1] - this.tail.pos[1];
      if (Math.abs(dif0) > 1 || Math.abs(dif1) > 1) {
        this.tail.move([Math.sign(dif0), Math.sign(dif1)]);
      }
    }
  }
}

class Tail extends Knot {
  constructor() {
    super();
    this.visited = new Set();
    this.visited.add("0.0");
  }
  move(vec) {
    this.pos[0] += vec[0];
    this.pos[1] += vec[1];
    this.visited.add(this.pos.join("."));
    if (this.tail) {
      const dif0 = this.pos[0] - this.tail.pos[0];
      const dif1 = this.pos[1] - this.tail.pos[1];
      if (Math.abs(dif0) > 1 || Math.abs(dif1) > 1) {
        this.tail.move([Math.sign(dif0), Math.sign(dif1)]);
      }
    }
  }
}

const knots = new Array(10);
knots[0] = new Knot();
knots[1] = new Tail();
for (let i = 2; i < 9; i++) {
  knots[i] = new Knot();
}
knots[9] = new Tail();
for (let i = 0; i < 9; i++) {
  knots[i].tail = knots[i + 1];
}

for (const line of input) {
  const [dir, dist] = line.split(" ");
  const vec = [0, 0];
  if (dir === "U") {
    vec[1] = 1;
  } else if (dir === "D") {
    vec[1] = -1;
  } else if (dir === "L") {
    vec[0] = -1;
  } else {
    vec[0] = 1;
  }
  for (let i = 0; i < Number(dist); i++) {
    knots[0].move(vec);
  }
}

console.log(knots[1].visited.size);
console.log(knots[9].visited.size);
