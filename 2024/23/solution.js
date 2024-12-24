const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split("-"));

class Node {
  constructor(name) {
    this.name = name;
    this.neighbors = new Set();
  }
}

const nodeMap = new Map();

for (const [nameA, nameB] of input) {
  if (!nodeMap.has(nameA)) nodeMap.set(nameA, new Node(nameA));
  if (!nodeMap.has(nameB)) nodeMap.set(nameB, new Node(nameB));
  const nodeA = nodeMap.get(nameA);
  const nodeB = nodeMap.get(nameB);
  nodeA.neighbors.add(nodeB);
  nodeB.neighbors.add(nodeA);
}

const trios = new Set();
for (const [name, node] of nodeMap) {
  if (name[0] !== "t") continue;
  for (const neighbor of node.neighbors) {
    const intersection = node.neighbors.intersection(neighbor.neighbors);
    for (const third of intersection) {
      const trio = [name, neighbor.name, third.name];
      trio.sort();
      trios.add(trio.join("-"));
    }
  }
}

console.log(trios.size);
