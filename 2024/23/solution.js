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

const nodes = Array.from(nodeMap.values());
let part2Set = new Set();

for (let i = 0; i < nodes.length; i++) {
  const checked = new Set();
  for (let j = i + 1; j < nodes.length; j++) {
    if (!nodes[i].neighbors.has(nodes[j])) continue;
    if (checked.has(nodes[j])) continue;
    checked.add(nodes[j]);
    const network = new Set([nodes[i], nodes[j]]);
    for (let k = j + 1; k < nodes.length; k++) {
      if (nodes[k].neighbors.isSupersetOf(network)) {
        checked.add(nodes[k]);
        network.add(nodes[k]);
      }
    }
    if (network.size > part2Set.size) {
      part2Set = network;
    }
  }
}

const part2 = Array.from(part2Set.keys()).map((a) => a.name);
part2.sort();
console.log(part2.join(","));
