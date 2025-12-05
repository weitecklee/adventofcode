import * as fs from "fs";
import * as path from "path";

const puzzleInput = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n");

function toCoords(s: string): number[] {
  return s.split(",").map(Number);
}

function parseInput(data: string[]): Map<string, number> {
  const papers: Map<string, number> = new Map();
  for (const [i, row] of data.entries()) {
    for (const [j, ch] of row.split("").entries()) {
      if (ch === "@") {
        papers.set([i, j].toString(), 0);
      }
    }
  }
  for (const [k, _] of papers) {
    const coords = toCoords(k);
    for (let r = coords[0] - 1; r <= coords[0] + 1; r++) {
      for (let c = coords[1] - 1; c <= coords[1] + 1; c++) {
        if (r === coords[0] && c === coords[1]) continue;
        const tmp = [r, c].toString();
        if (papers.has(tmp)) {
          papers.set(tmp, papers.get(tmp)! + 1);
        }
      }
    }
  }
  return papers;
}

function removePaper(papers: Map<string, number>): number {
  let res = 0;
  for (const v of papers.values()) {
    if (v < 4) res++;
  }
  return res;
}

function removePaperContinuous(papers: Map<string, number>): number {
  let res = 0;
  while (true) {
    let removed = 0;
    for (const [k, v] of papers.entries()) {
      if (v < 4) {
        removed++;
        papers.delete(k);
        const coords = toCoords(k);
        for (let r = coords[0] - 1; r <= coords[0] + 1; r++) {
          for (let c = coords[1] - 1; c <= coords[1] + 1; c++) {
            if (r === coords[0] && c === coords[1]) continue;
            const tmp = [r, c].toString();
            if (papers.has(tmp)) {
              papers.set(tmp, papers.get(tmp)! - 1);
            }
          }
        }
      }
    }
    if (removed === 0) break;
    res += removed;
  }
  return res;
}

function part1(papers: Map<string, number>): number {
  return removePaper(papers);
}

function part2(papers: Map<string, number>): number {
  return removePaperContinuous(papers);
}

const paperMap = parseInput(puzzleInput);

console.log(part1(paperMap));
console.log(part2(paperMap));
