const fs = require("fs");
const path = require("path");
const mathjs = require("mathjs");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split(": "))
  .map((b) => [b[0], b[1].split(" ")]);

/*
  Strategy is to take 1000 random pairs of nodes and find
  the shortest path between them. We tally up all the
  node-to-node connections in each path and find the three
  most common connections. Presumably, these three are the
  most critical connections that we want to cut.
  Seems to work well enough, 1000 is a big enough number that
  the same three connections are found each time.
  For the record, my input had 1458 nodes, which makes
  1458 * 1457 / 2 = 1062153 pairs.
*/

class Node {
  constructor(name) {
    this.name = name;
    this.neighbors = new Set();
  }

  addNeighbor(neighbor) {
    this.neighbors.add(neighbor);
  }
  removeNeighbor(neighbor) {
    this.neighbors.delete(neighbor);
  }
}

const nodeMap = new Map();

for (const [name, neighbors] of input) {
  if (!nodeMap.has(name)) {
    nodeMap.set(name, new Node(name));
  }
  const currNode = nodeMap.get(name);
  for (const neighbor of neighbors) {
    if (!nodeMap.has(neighbor)) {
      nodeMap.set(neighbor, new Node(neighbor));
    }
    const neighborNode = nodeMap.get(neighbor);
    currNode.addNeighbor(neighborNode);
    neighborNode.addNeighbor(currNode);
  }
}

const nodes = Array.from(nodeMap.values());
const connectionCounts = new Map();

function addConnection(nodes) {
  nodes.sort();
  const connKey = nodes.join("-");
  connectionCounts.set(connKey, (connectionCounts.get(connKey) || 0) + 1);
}

function addPath(path) {
  for (let i = 0; i < path.length - 1; i++) {
    addConnection([path[i], path[i + 1]]);
  }
}

function findPath(nodeA, nodeB) {
  if (nodeA === nodeB) return [];
  if (nodeA.neighbors.has(nodeB)) return [nodeA.name, nodeB.name];
  const queue = [[nodeA, [nodeA.name]]];
  const visited = new Set();
  let i = 0;
  while (i < queue.length) {
    const [currNode, path] = queue[i];
    i++;
    if (currNode === nodeB) {
      return path;
    }
    if (visited.has(currNode)) continue;
    visited.add(currNode);
    for (const neighbor of currNode.neighbors) {
      queue.push([neighbor, [...path, neighbor.name]]);
    }
  }
  return [];
}

for (let i = 0; i < 1000; i++) {
  const [nodeA, nodeB] = mathjs.pickRandom(nodes, 2);
  addPath(findPath(nodeA, nodeB));
}

const counts = Array.from(connectionCounts.entries()).sort(
  (a, b) => b[1] - a[1]
);

const candidates = counts.slice(0, 3).map((a) => a[0].split("-"));

for (const [nameA, nameB] of candidates) {
  const nodeA = nodeMap.get(nameA);
  const nodeB = nodeMap.get(nameB);
  nodeA.removeNeighbor(nodeB);
  nodeB.removeNeighbor(nodeA);
}

const group = new Set();
const queue = [nodes[0]];

let i = 0;
while (i < queue.length) {
  const currNode = queue[i];
  i++;
  if (group.has(currNode)) continue;
  group.add(currNode);
  for (const neighbor of currNode.neighbors) {
    queue.push(neighbor);
  }
}
const sizeA = group.size;
const sizeB = nodes.length - sizeA;
console.log(sizeA * sizeB);
