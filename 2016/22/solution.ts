import * as fs from "fs";
import * as path from "path";

const puzzleInput = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n");

const nodeRegex = /^(\S+)\s+(\d+)\w\s+(\d+)\w\s+(\d+)\w\s+(\d+)%$/;

class Node {
  name: string;
  size: number;
  used: number;
  avail: number;
  constructor(line: string) {
    const parts = line.match(nodeRegex)!;
    this.name = parts[1];
    this.size = Number(parts[2]);
    this.used = Number(parts[3]);
    this.avail = Number(parts[4]);
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

console.log(part1());
