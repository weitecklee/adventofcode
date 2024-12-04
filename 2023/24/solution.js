const fs = require("fs");
const path = require("path");
const mathjs = require("mathjs");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split(" @ ").map((b) => b.split(", ").map(Number)));

class Hailstone {
  constructor(pos, vel) {
    this.pos = pos;
    this.vel = vel;
    this.m = vel[1] / vel[0];
    this.b = pos[1] - this.m * pos[0];
  }
  findIntersection(hailstone) {
    const x = (hailstone.b - this.b) / (this.m - hailstone.m);
    const y = this.m * x + this.b;
    const t1 = (x - this.pos[0]) / this.vel[0];
    const t2 = (x - hailstone.pos[0]) / hailstone.vel[0];
    return [x, y, t1, t2];
  }
}

const hailstones = input.map(([pos, vel]) => new Hailstone(pos, vel));
let part1 = 0;
for (let i = 0; i < hailstones.length; i++) {
  for (let j = i + 1; j < hailstones.length; j++) {
    const intersection = hailstones[i].findIntersection(hailstones[j]);
    if (
      intersection[2] > 0 &&
      intersection[3] > 0 &&
      intersection[0] >= 200000000000000 &&
      intersection[0] <= 400000000000000 &&
      intersection[1] >= 200000000000000 &&
      intersection[1] <= 400000000000000
    )
      part1++;
  }
}

console.log(part1);

/*
  Find some pos p and vel v such that for each hailstone with pos p_i and vel v_i, there exists a time t_i > 0 such that:
    p[0] + v[0] * t_i = p_i[0] + v_i[0] * t_i
    p[1] + v[1] * t_i = p_i[1] + v_i[1] * t_i
    p[2] + v[2] * t_i = p_i[2] + v_i[2] * t_i

  Rewrite to eliminate t_i:
  t_i = (p[0] - p_i[0]) / (v_i[0] - v[0]) = (p[1] - p_i[1]) / (v_i[1] - v[1]) = (p[2] - p_i[2]) / (v_i[2] - v[2])
  (p[0] - p_i[0]) * (v_i[1] - v[1]) = (p[1] - p_i[1]) * (v_i[0] - v[0])
  (p[0] - p_i[0]) * (v_i[2] - v[2]) = (p[2] - p_i[2]) * (v_i[0] - v[0])

  Expand:
  p[0] * v_i[1] - p[0] * v[1] - p_i[0] * v_i[1] + p_i[0] * v[1] = p[1] * v_i[0] - p[1] * v[0] - p_i[1] * v_i[0] + p_i[1] * v[0]  // equation (a)
  p[0] * v_i[2] - p[0] * v[2] - p_i[0] * v_i[2] + p_i[0] * v[2] = p[2] * v_i[0] - p[2] * v[0] - p_i[2] * v_i[0] + p_i[2] * v[0]  // equation (b)

  If we introduce another hailstone with pos p_j and vel v_j, we get similarly:
  p[0] * v_j[1] - p[0] * v[1] - p_j[0] * v_j[1] + p_j[0] * v[1] = p[1] * v_j[0] - p[1] * v[0] - p_j[1] * v_j[0] + p_j[1] * v[0]  // equation (c)
  p[0] * v_j[2] - p[0] * v[2] - p_j[0] * v_j[2] + p_j[0] * v[2] = p[2] * v_j[0] - p[2] * v[0] - p_j[2] * v_j[0] + p_j[2] * v[0]  // equation (d)

  Subtracting the pairs of equations (a/c and b/d) from each other eliminates the nonlinear terms (e.g., p[0] * v[1])
  p[0] * v_i[1] - p[0] * v_j[1] - p_i[0] * v_i[1] + p_j[0] * v_j[1] + p_i[0] * v[1] - p_j[0] * v[1] = p[1] * v_i[0] - p[1] * v_j[0] - p_i[1] * v_i[0] + p_j[1] * v_j[0] + p_i[1] * v[0] - p_j[1] * v[0]
  p[0] * v_i[2] - p[0] * v_j[2] - p_i[0] * v_i[2] + p_j[0] * v_j[2] + p_i[0] * v[2] - p_j[0] * v[2] = p[2] * v_i[0] - p[2] * v_j[0] - p_i[2] * v_i[0] + p_j[2] * v_j[0] + p_i[2] * v[0] - p_j[2] * v[0]

  Rearrange to put unknowns on the left and knowns on the right:
  p[0] * (v_i[1] - v_j[1]) + p[1] * (v_j[0] - v_i[0]) + v[0] * (p_j[1] - p_i[1]) + v[1] * (p_i[0] - p_j[0]) = p_i[0] * v_i[1] - p_j[0] * v_j[1] - p_i[1] * v_i[0] + p_j[1] * v_j[0]
  p[0] * (v_i[2] - v_j[2]) + p[2] * (v_j[0] - v_i[0]) + v[0] * (p_j[2] - p_i[2]) + v[2] * (p_i[0] - p_j[0]) = p_i[0] * v_i[2] - p_j[0] * v_j[2] - p_i[2] * v_i[0] + p_j[2] * v_j[0]

  Use four hailstone pairs to get six equations (i = 0, j = 1, 2, 3):
  p[0] * (v_0[1] - v_1[1]) + p[1] * (v_1[0] - v_0[0]) + v[0] * (p_1[1] - p_0[1]) + v[1] * (p_0[0] - p_1[0]) = p_0[0] * v_0[1] - p_1[0] * v_1[1] - p_0[1] * v_0[0] + p_1[1] * v_1[0]
  p[0] * (v_0[2] - v_1[2]) + p[2] * (v_1[0] - v_0[0]) + v[0] * (p_1[2] - p_0[2]) + v[2] * (p_0[0] - p_1[0]) = p_0[0] * v_0[2] - p_1[0] * v_1[2] - p_0[2] * v_0[0] + p_1[2] * v_1[0]
  p[0] * (v_0[1] - v_2[1]) + p[1] * (v_2[0] - v_0[0]) + v[0] * (p_2[1] - p_0[1]) + v[1] * (p_0[0] - p_2[0]) = p_0[0] * v_0[1] - p_2[0] * v_2[1] - p_0[1] * v_0[0] + p_2[1] * v_2[0]
  p[0] * (v_0[2] - v_2[2]) + p[2] * (v_2[0] - v_0[0]) + v[0] * (p_2[2] - p_0[2]) + v[2] * (p_0[0] - p_2[0]) = p_0[0] * v_0[2] - p_2[0] * v_2[2] - p_0[2] * v_0[0] + p_2[2] * v_2[0]
  p[0] * (v_0[1] - v_3[1]) + p[1] * (v_3[0] - v_0[0]) + v[0] * (p_3[1] - p_0[1]) + v[1] * (p_0[0] - p_3[0]) = p_0[0] * v_0[1] - p_3[0] * v_3[1] - p_0[1] * v_0[0] + p_3[1] * v_3[0]
  p[0] * (v_0[2] - v_3[2]) + p[2] * (v_3[0] - v_0[0]) + v[0] * (p_3[2] - p_0[2]) + v[2] * (p_0[0] - p_3[0]) = p_0[0] * v_0[2] - p_3[0] * v_3[2] - p_0[2] * v_0[0] + p_3[2] * v_3[0]

  With that, we have six equations and six unknowns, no nonlinear terms so it can be solved with Cramer's rule.
  Ax = b
  x = [p[0], p[1], p[2], v[0], v[1], v[2]]T
  A = [
    [v_0[1] - v_1[1], v_1[0] - v_0[0],               0, p_1[1] - p_0[1], p_0[0] - p_1[0],               0],
    [v_0[2] - v_1[2],               0, v_1[0] - v_0[0], p_1[2] - p_0[2],               0, p_0[0] - p_1[0]],
    [v_0[1] - v_2[1], v_2[0] - v_0[0],               0, p_2[1] - p_0[1], p_0[0] - p_2[0],               0],
    [v_0[2] - v_2[2],               0, v_2[0] - v_0[0], p_2[2] - p_0[2],               0, p_0[0] - p_2[0]],
    [v_0[1] - v_3[1], v_3[0] - v_0[0],               0, p_3[1] - p_0[1], p_0[0] - p_3[0],               0],
    [v_0[2] - v_3[2],               0, v_3[0] - v_0[0], p_3[2] - p_0[2],               0, p_0[0] - p_3[0]]
  ]

  b = [
    p_0[0] * v_0[1] - p_1[0] * v_1[1] - p_0[1] * v_0[0] + p_1[1] * v_1[0],
    p_0[0] * v_0[2] - p_1[0] * v_1[2] - p_0[2] * v_0[0] + p_1[2] * v_1[0],
    p_0[0] * v_0[1] - p_2[0] * v_2[1] - p_0[1] * v_0[0] + p_2[1] * v_2[0],
    p_0[0] * v_0[2] - p_2[0] * v_2[2] - p_0[2] * v_0[0] + p_2[2] * v_2[0],
    p_0[0] * v_0[1] - p_3[0] * v_3[1] - p_0[1] * v_0[0] + p_3[1] * v_3[0],
    p_0[0] * v_0[2] - p_3[0] * v_3[2] - p_0[2] * v_0[0] + p_3[2] * v_3[0]
  ]
  x = A^-1 * b
*/

const p_0 = hailstones[0].pos;
const v_0 = hailstones[0].vel;
const p_1 = hailstones[1].pos;
const v_1 = hailstones[1].vel;
const p_2 = hailstones[2].pos;
const v_2 = hailstones[2].vel;
const p_3 = hailstones[3].pos;
const v_3 = hailstones[3].vel;

const A = [
  [v_0[1] - v_1[1], v_1[0] - v_0[0], 0, p_1[1] - p_0[1], p_0[0] - p_1[0], 0],
  [v_0[2] - v_1[2], 0, v_1[0] - v_0[0], p_1[2] - p_0[2], 0, p_0[0] - p_1[0]],
  [v_0[1] - v_2[1], v_2[0] - v_0[0], 0, p_2[1] - p_0[1], p_0[0] - p_2[0], 0],
  [v_0[2] - v_2[2], 0, v_2[0] - v_0[0], p_2[2] - p_0[2], 0, p_0[0] - p_2[0]],
  [v_0[1] - v_3[1], v_3[0] - v_0[0], 0, p_3[1] - p_0[1], p_0[0] - p_3[0], 0],
  [v_0[2] - v_3[2], 0, v_3[0] - v_0[0], p_3[2] - p_0[2], 0, p_0[0] - p_3[0]],
];

const b = [
  p_0[0] * v_0[1] - p_1[0] * v_1[1] - p_0[1] * v_0[0] + p_1[1] * v_1[0],
  p_0[0] * v_0[2] - p_1[0] * v_1[2] - p_0[2] * v_0[0] + p_1[2] * v_1[0],
  p_0[0] * v_0[1] - p_2[0] * v_2[1] - p_0[1] * v_0[0] + p_2[1] * v_2[0],
  p_0[0] * v_0[2] - p_2[0] * v_2[2] - p_0[2] * v_0[0] + p_2[2] * v_2[0],
  p_0[0] * v_0[1] - p_3[0] * v_3[1] - p_0[1] * v_0[0] + p_3[1] * v_3[0],
  p_0[0] * v_0[2] - p_3[0] * v_3[2] - p_0[2] * v_0[0] + p_3[2] * v_3[0],
];

const x = mathjs.lusolve(mathjs.matrix(A), mathjs.matrix(b));
console.log(x._data.slice(0, 3).reduce((a, b) => a + Math.round(Number(b)), 0));
