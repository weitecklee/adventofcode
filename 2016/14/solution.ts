import * as fs from "fs";
import * as path from "path";

const puzzleInput = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8");

import * as crypto from "crypto";

function hasher(i: number): string {
  return crypto
    .createHash("md5")
    .update(puzzleInput + i)
    .digest("hex");
}

const tripletRegex = /(\w)\1{2}/;
const quintupletRegex = /(\w)\1{4}/g;

function part1(): number {
  let i = 0;
  let n = 0;
  const triplets: [number, string][] = [];
  const quintuplets: Map<string, number[]> = new Map(
    "1234567890abcdef".split("").map((c) => [c, []])
  );

  const keyIndices: number[] = [];
  while (keyIndices.length < 64) {
    const h = hasher(i);
    let match: RegExpMatchArray | null;
    if ((match = h.match(tripletRegex))) {
      triplets.push([i, match[1]]);
    }
    let matches: RegExpMatchArray | null;
    if ((matches = h.match(quintupletRegex))) {
      for (const match of matches) {
        const c = match[0];
        quintuplets.get(c)!.push(i);
      }
    }
    if (triplets.length && i >= triplets[0][0] + 1000) {
      const [j, c] = triplets.shift() as [number, string];
      const arr = quintuplets.get(c)!;
      for (const k of arr) {
        if (k > j && k <= j + 1000) {
          keyIndices.push(j);
          break;
        }
        if (k > j + 1000) break;
      }
    }
    i++;
  }
  return keyIndices[63];
}

console.log(part1());
