import * as fs from "fs";
import * as path from "path";

const puzzleInput = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split(",").map(Number));

class Point {
  pos: number[];
  constellation: Set<Point>;
  constructor(pos: number[]) {
    this.pos = pos;
    this.constellation = new Set([this]);
  }
}

function calcDist(p1: Point, p2: Point): number {
  return p1.pos.reduce((a, b, i) => a + Math.abs(b - p2.pos[i]), 0);
}

const points = puzzleInput.map((a) => new Point(a));
let constellations: Set<Set<Point>> = new Set();

for (let i = 0; i < points.length; i++) {
  constellations.add(points[i].constellation);
  for (let j = i + 1; j < points.length; j++) {
    if (calcDist(points[i], points[j]) <= 3) {
      for (const p of points[j].constellation) {
        points[i].constellation.add(p);
      }
      constellations.delete(points[j].constellation);
      points[j].constellation = points[i].constellation;
    }
  }
}

let nConstellations = 0;
while (nConstellations !== constellations.size) {
  nConstellations = constellations.size;
  const constellations2: Set<Set<Point>> = new Set();
  for (const con1 of constellations) {
    let noIntersections = true;
    for (const con2 of constellations2) {
      if (!con1.isDisjointFrom(con2)) {
        con1.forEach((p) => {
          con2.add(p);
        });
        noIntersections = false;
        break;
      }
    }
    if (noIntersections) constellations2.add(con1);
  }
  constellations = constellations2;
}

console.log(nConstellations);
