import * as fs from "fs";
import * as path from "path";
import MinHeap from "../../utils/MinHeap";

const puzzleInput = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n");

const rMax = puzzleInput.length - 1;
const cMax = Math.max(...puzzleInput.map((r) => r.length)) - 1;

const paddedInput = puzzleInput.map((r) => r.padEnd(cMax + 1, " "));

class Node {
  name: string;
  coords: number[];
  coordString: string;
  neighbors: Map<Node, number>;
  isOuter: boolean;

  constructor(name: string, coords: number[]) {
    this.name = name;
    this.coords = coords;
    this.coordString = coords.join(",");
    this.neighbors = new Map();
    this.isOuter =
      coords[0] === 2 ||
      coords[0] === rMax - 2 ||
      coords[1] === 2 ||
      coords[1] === cMax - 2;
  }
}

const portalMap: Map<string, Node[]> = new Map();
const nodeMap: Map<string, Node> = new Map();

for (let r = 0; r <= rMax; r++) {
  for (let c = 0; c <= cMax; c++) {
    if (/\w/.test(paddedInput[r][c])) {
      if (r + 1 <= rMax && /\w/.test(paddedInput[r + 1][c])) {
        const portalName = paddedInput[r][c] + paddedInput[r + 1][c];
        if (!portalMap.has(portalName)) {
          portalMap.set(portalName, []);
        }
        const portal = portalMap.get(portalName)!;
        let coords = [r + 2, c];
        if (r - 1 >= 0 && paddedInput[r - 1][c] === ".") {
          coords = [r - 1, c];
        }
        const node = new Node(portalName, coords);
        portal.push(node);
        nodeMap.set(node.coordString, node);
      } else if (c + 1 <= cMax && /\w/.test(paddedInput[r][c + 1])) {
        const portalName = paddedInput[r][c] + paddedInput[r][c + 1];
        if (!portalMap.has(portalName)) {
          portalMap.set(portalName, []);
        }
        const portal = portalMap.get(portalName)!;
        let coords = [r, c + 2];
        if (c - 1 >= 0 && paddedInput[r][c - 1] === ".") {
          coords = [r, c - 1];
        }
        const node = new Node(portalName, coords);
        portal.push(node);
        nodeMap.set(node.coordString, node);
      }
    }
  }
}

const directions = [
  [-1, 0],
  [1, 0],
  [0, -1],
  [0, 1],
];

for (const node of nodeMap.values()) {
  const queue = [[0, ...node.coords]];
  const visited: Set<string> = new Set([node.coordString]);
  while (queue.length) {
    const [d, r, c] = queue.pop()!;
    for (const [dr, dc] of directions) {
      const [r2, c2] = [r + dr, c + dc];
      if (paddedInput[r2][c2] !== ".") continue;
      const coordString = `${r2},${c2}`;
      if (visited.has(coordString)) continue;
      visited.add(coordString);
      if (nodeMap.has(coordString)) {
        const node2 = nodeMap.get(coordString)!;
        node.neighbors.set(node2, d + 1);
        continue;
      }
      queue.push([d + 1, r2, c2]);
    }
  }
}

for (const arr of portalMap.values()) {
  if (arr.length === 1) continue;
  const [node1, node2] = arr;
  node1.neighbors.set(node2, 1);
  node2.neighbors.set(node1, 1);
}

const nodeAA = portalMap.get("AA")![0];
const nodeZZ = portalMap.get("ZZ")![0];

const queue: [number, Node, Set<Node>][] = [[0, nodeAA, new Set([nodeAA])]];

while (queue.length) {
  const [steps, currentNode, visitedNodes] = MinHeap.pop(queue)!;
  if (currentNode === nodeZZ) {
    console.log(steps);
    break;
  }
  for (const [neighbor, d] of currentNode.neighbors) {
    if (visitedNodes.has(neighbor)) continue;
    const visited2 = new Set(visitedNodes);
    visited2.add(neighbor);
    MinHeap.push(queue, [steps + d, neighbor, visited2]);
  }
}

const queue2: [number, number, Node][] = [[0, 0, nodeAA]];
const visited2 = new Set(["0AA1"]);

while (queue2.length) {
  const [steps, level, currentNode] = MinHeap.pop(queue2)!;
  if (currentNode === nodeZZ) {
    console.log(steps);
    break;
  }
  for (const [neighbor, d] of currentNode.neighbors) {
    if ((neighbor === nodeZZ || neighbor === nodeAA) && level !== 0) continue;
    let level2 = level;
    if (neighbor.name === currentNode.name) {
      if (currentNode.isOuter) level2--;
      else level2++;
    }
    if (level2 < 0) continue;
    const neighborKey = `${level2}${neighbor.name}${neighbor.isOuter ? 1 : 2}`;
    if (visited2.has(neighborKey)) continue;
    visited2.add(neighborKey);
    MinHeap.push(queue2, [steps + d, level2, neighbor]);
  }
}
