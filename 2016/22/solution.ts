import * as fs from "fs";
import * as path from "path";
import MinHeap from "../../utils/MinHeap";

const puzzleInput = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n");

const nodeRegex = /^(\S+)\s+(\d+)\w\s+(\d+)\w\s+(\d+)\w\s+(\d+)%$/;

class Node {
  name: string;
  size: number;
  used: number;
  avail: number;
  pos: number[];
  coords: string;
  constructor(line: string) {
    const parts = line.match(nodeRegex)!;
    this.name = parts[1];
    this.size = Number(parts[2]);
    this.used = Number(parts[3]);
    this.avail = Number(parts[4]);
    const coords = this.name.match(/\d+/g);
    this.pos = coords!.map(Number);
    this.coords = this.pos.join(",");
  }
}

const nodes = puzzleInput.slice(2).map((a) => new Node(a));

function part1(): number {
  let res = 0;
  for (let i = 0; i < nodes.length; i++) {
    for (let j = i + 1; j < nodes.length; j++) {
      if (nodes[i].used > 0 && nodes[i].used <= nodes[j].avail) res++;
      if (nodes[j].used > 0 && nodes[j].used <= nodes[i].avail) res++;
    }
  }
  return res;
}

// The solution method is hinted at by the problem description.
// All nodes (except for one empty node) have more data than there is space
// available on any other node so you cannot combine data from multiple nodes
// in one node. The only way is to keep swapping the empty node with another.
// There are some outlier nodes, which are very large and very full. Those
// are effectively walls and must be maneuvered around.
// As shown in the problem description, locate the empty node, move it to
// the left of the target node, then move the target node by swapping and
// shifting the empty node around it, all while avoiding the wall nodes.
// Calculate shortest route from empty node to left of target node, swap with
// target node, then start the following process:
// - repeatedly move to left of [current node holding target data] by moving
//   around it and then swap, as illustrated in problem description
// (This assumes there are no wall nodes that interfere in first two rows)
// Each cycle of this process takes 5 steps.
// Resulting number of steps is sum of:
// - shortest route from empty node to left of target node
// - 1 for first swap
// - 5 * (maximum x coordinate - 1)   (because it just got swapped)

function part2(): number {
  let steps = 0;
  const nodeMap: Map<string, Node> = nodes.reduce((a, b) => {
    a.set(b.coords, b);
    return a;
  }, new Map());
  const xMax = Math.max(...nodes.map((n) => n.pos[0]));
  const yMax = Math.max(...nodes.map((n) => n.pos[1]));

  // confirm there are no wall nodes in top two rows(y = 0 and y = 1)
  for (let x = 0; x <= xMax; x++) {
    if (nodeMap.get(`${x},0`)!.used > 100) return -1;
    if (nodeMap.get(`${x},1`)!.used > 100) return -1;
  }

  let emptyPos: number[] = [];
  for (const node of nodes) {
    if (node.used === 0) {
      emptyPos = node.pos;
      break;
    }
  }

  const directions = [
    [-1, 0],
    [1, 0],
    [0, 1],
    [0, -1],
  ];

  function heuristic(x: number, y: number): number {
    return Math.abs(xMax - 1 - x) + y;
  }

  const queue: [number, ...number[]][] = [[0, 0, emptyPos[0], emptyPos[1]]];
  const visited: Map<string, number> = new Map();
  while (queue.length) {
    let [_, curr, x, y] = MinHeap.pop(queue) as number[];
    if (x === xMax - 1 && y === 0) {
      steps = curr;
      break;
    }
    curr++;
    for (const [dx, dy] of directions) {
      const [x2, y2] = [x + dx, y + dy];
      if (x2 < 0 || y2 < 0 || x2 > xMax || y2 > yMax) continue;
      const node = nodeMap.get(`${x2},${y2}`)!;
      if (node.used > 100) continue;
      if (visited.has(node.coords) && visited.get(node.coords)! <= curr)
        continue;
      visited.set(node.coords, curr);
      MinHeap.push(queue, [curr + heuristic(x2, y2), curr, x2, y2]);
    }
  }

  return steps + 1 + (xMax - 1) * 5;
}

console.log(part1());
console.log(part2());
