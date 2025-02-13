import * as fs from "fs";
import * as path from "path";

const puzzleInput = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n");

let cubes: Set<string> = new Set();

for (let x = 0; x < puzzleInput.length; x++) {
  for (let y = 0; y < puzzleInput[0].length; y++) {
    if (puzzleInput[x][y] === "#") {
      cubes.add(`${x},${-y},0,0`);
    }
  }
}

function simulate(cubes: Set<string>, isPart2: boolean = false): number {
  let currCubes = new Set(cubes);

  for (let i = 0; i < 6; i++) {
    const activeNeighbors: Map<string, number> = new Map();

    for (const cube of currCubes) {
      const [x, y, z, w] = cube.split(",").map(Number);
      for (let dx = -1; dx <= 1; dx++) {
        for (let dy = -1; dy <= 1; dy++) {
          for (let dz = -1; dz <= 1; dz++) {
            for (let dw = isPart2 ? -1 : 0; dw <= (isPart2 ? 1 : 0); dw++) {
              if (dx === 0 && dy === 0 && dz === 0 && dw === 0) continue;
              const neighborCoord = [x + dx, y + dy, z + dz, w + dw].join(",");
              activeNeighbors.set(
                neighborCoord,
                (activeNeighbors.get(neighborCoord) || 0) + 1
              );
            }
          }
        }
      }
    }

    const newCubes: Set<string> = new Set();

    for (const cube of currCubes) {
      if (
        activeNeighbors.has(cube) &&
        (activeNeighbors.get(cube)! === 2 || activeNeighbors.get(cube)! === 3)
      ) {
        newCubes.add(cube);
      }
      activeNeighbors.delete(cube);
    }

    for (const [cube, nActive] of activeNeighbors) {
      if (nActive === 3) {
        newCubes.add(cube);
      }
    }
    currCubes = newCubes;
  }

  return currCubes.size;
}

console.log(simulate(cubes));
console.log(simulate(cubes, true));
