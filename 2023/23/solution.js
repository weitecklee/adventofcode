const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split(""));

const start = [input[0].indexOf("."), 0];
const end = [input[input.length - 1].indexOf("."), input.length - 1];

const queue = [[...start, new Set()]];

const directions = [
  [0, 1],
  [1, 0],
  [-1, 0],
  [0, -1],
];
const slopeTiles = ["v", ">", "<", "^"];

// console.time("part1");
let part1 = 0;

while (queue.length) {
  const [x, y, visited] = queue.pop();
  if (x === end[0] && y === end[1]) {
    part1 = Math.max(part1, visited.size);
    continue;
  }
  visited.add(`${x},${y}`);
  const slopeIndex = slopeTiles.indexOf(input[y][x]);
  if (slopeIndex >= 0) {
    const [dx, dy] = directions[slopeIndex];
    const x2 = x + dx;
    const y2 = y + dy;
    queue.push([x2, y2, new Set(visited)]);
    continue;
  }
  for (let i = 0; i < directions.length; i++) {
    const x2 = x + directions[i][0];
    const y2 = y + directions[i][1];
    if (x2 < 0 || x2 >= input[0].length || y2 < 0 || y2 >= input.length)
      continue;
    if (input[y2][x2] === "#") continue;
    if (slopeTiles.indexOf(input[y2][x2]) === 3 - i) continue;
    if (visited.has(`${x2},${y2}`)) continue;
    queue.push([x2, y2, new Set(visited)]);
  }
}

// console.timeEnd("part1");
console.log(part1);

/*
  Part 1 uses simply BFS to find longest path (while obeying slopes).
  Without slopes, grid becomes far too open to use above method (takes forever).
  So first we go through grid and find intersections (nodes) and construct a graph,
  keeping track of distances between neighbor nodes.
  Then we use BFS to find longest path between start and end nodes.
  (My input had 36 nodes)
*/

// console.time("part2");
function Node(x, y) {
  this.addr = `${x},${y}`;
  this.x = x;
  this.y = y;
  this.neighbors = new Map();
}

const nodes = new Map([
  [`${start[0]},${start[1]}`, new Node(...start)],
  [`${end[0]},${end[1]}`, new Node(...end)],
]);
const visited = new Set([`${start[0]},${start[1]}`]);
const queue2 = [[start[0], start[1]]];

// Find nodes
while (queue2.length) {
  const [x, y, distance, origNode] = queue2.pop();
  const next = [];
  for (let i = 0; i < directions.length; i++) {
    const x2 = x + directions[i][0];
    const y2 = y + directions[i][1];
    if (x2 < 0 || x2 >= input[0].length || y2 < 0 || y2 >= input.length)
      continue;
    if (input[y2][x2] === "#") continue;
    if (visited.has(`${x2},${y2}`)) continue;
    visited.add(`${x2},${y2}`);
    next.push([x2, y2]);
  }
  if (next.length > 1) {
    nodes.set(`${x},${y}`, new Node(x, y));
  }
  queue2.push(...next);
}

visited.clear();

queue2.push([start[0], start[1], nodes.get(`${start[0]},${start[1]}`)]);

// Construct graph
// From each node, BFS to find neighbors and distances
while (queue2.length) {
  // queue2 is queue of nodes
  const [x, y, origNode] = queue2.shift();
  const queue3 = [];
  for (const [dx, dy] of directions) {
    const x2 = x + dx;
    const y2 = y + dy;
    if (x2 < 0 || x2 >= input[0].length || y2 < 0 || y2 >= input.length)
      continue;
    if (input[y2][x2] === "#") continue;
    queue3.push([x2, y2, 1]);
  }
  // queue3 is BFS from node to neighbors
  // When neighbor is found, add distance data to node/neighbor
  // Then add neighbor to queue2
  while (queue3.length) {
    const [x2, y2, distance] = queue3.shift();
    for (const [dx, dy] of directions) {
      const x3 = x2 + dx;
      const y3 = y2 + dy;
      if (x3 < 0 || x3 >= input[0].length || y3 < 0 || y3 >= input.length)
        continue;
      if (input[y3][x3] === "#") continue;
      if (nodes.has(`${x3},${y3}`)) {
        if (origNode.addr === `${x3},${y3}`) continue;
        origNode.neighbors.set(`${x3},${y3}`, distance + 1);
        nodes.get(`${x3},${y3}`).neighbors.set(origNode.addr, distance + 1);
        queue2.push([x3, y3, nodes.get(`${x3},${y3}`)]);
      } else {
        if (visited.has(`${x3},${y3}`)) continue;
        visited.add(`${x3},${y3}`);
        queue3.push([x3, y3, distance + 1]);
      }
    }
  }
}

queue2.push([
  nodes.get(`${start[0]},${start[1]}`),
  0,
  new Set([`${start[0]},${start[1]}`]),
]);

let part2 = 0;

// DFS through graph to find longest path
while (queue2.length) {
  const [node, distance, visited] = queue2.pop();
  if (node.addr === `${end[0]},${end[1]}`) {
    part2 = Math.max(part2, distance);
    continue;
  }
  for (const [addr, dist] of node.neighbors) {
    if (visited.has(addr)) continue;
    const visited2 = new Set(visited);
    visited2.add(addr);
    queue2.push([nodes.get(addr), distance + dist, visited2]);
  }
}

// console.timeEnd("part2");
console.log(part2);

/*

part1: 9.111s
part2: 33.332s

Wow this is super sloppy, maybe I'll get around to cleaning it up later.
Woof.

*/
