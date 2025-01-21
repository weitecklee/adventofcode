import * as fs from "fs";
import * as path from "path";

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n");

const directionRegex = /e|se|sw|w|ne|nw/g;
const directions: Map<string, number[]> = new Map([
  ["e", [2, 0]],
  ["se", [1, -1]],
  ["sw", [-1, -1]],
  ["w", [-2, 0]],
  ["ne", [1, 1]],
  ["nw", [-1, 1]],
]);

let tiles: Set<string> = new Set();

for (const line of input) {
  const moves = line.match(directionRegex)!;
  const loc = [0, 0];
  for (const move of moves) {
    const d = directions.get(move)!;
    loc[0] += d[0];
    loc[1] += d[1];
  }
  const coord = loc.join(",");
  if (tiles.has(coord)) tiles.delete(coord);
  else tiles.add(coord);
}

console.log(tiles.size);

for (let i = 0; i < 100; i++) {
  const adjacents: Map<string, number> = new Map();
  for (const coord of tiles) {
    const [x, y] = coord.split(",").map(Number);
    for (const [_, [dx, dy]] of directions) {
      const coord2 = [x + dx, y + dy].join(",");
      adjacents.set(coord2, (adjacents.get(coord2) || 0) + 1);
    }
  }

  const tiles2: Set<string> = new Set();
  for (const coord of tiles) {
    if (adjacents.has(coord) && adjacents.get(coord)! <= 2) tiles2.add(coord);
  }
  for (const [coord, n] of adjacents) {
    if (!tiles.has(coord) && n === 2) tiles2.add(coord);
  }

  tiles = tiles2;
}

console.log(tiles.size);
