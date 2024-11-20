const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n\n");

class Tile {
  constructor(data) {
    const tileLines = data.split("\n");
    this.id = Number(tileLines[0].match(/\d+/)[0]);
    this.imageData = tileLines.slice(1);
    this.neighbors = new Set();
    this.edgeSides = [];
    this.placed = false;
  }

  getTopSide() {
    return this.imageData[0];
  }

  getBottomSide() {
    return this.imageData[this.imageData.length - 1];
  }

  getLeftSide() {
    return this.imageData.map((line) => line[0]).join("");
  }

  getRightSide() {
    return this.imageData.map((line) => line[line.length - 1]).join("");
  }

  getSides() {
    return [
      this.getTopSide(),
      this.getRightSide(),
      this.getBottomSide(),
      this.getLeftSide(),
    ];
  }

  reverse() {
    this.imageData = this.imageData.map((line) =>
      line.split("").reverse().join("")
    );
  }

  rotate() {
    this.imageData = this.imageData.map((_, i) =>
      this.imageData
        .map((line) => line[i])
        .reverse()
        .join("")
    );
  }

  print() {
    console.log(`Tile ${this.id}:`);
    for (const line of this.imageData) {
      console.log(line);
    }
  }
}

const tiles = new Map();

for (const tileData of input) {
  const tile = new Tile(tileData);
  tiles.set(tile.id, tile);
}

const sides = new Map();

for (const tile of tiles.values()) {
  for (const side of tile.getSides()) {
    const reversedSide = side.split("").reverse().join("");
    if (sides.has(side)) {
      sides.get(side).push(tile.id);
    } else if (sides.has(reversedSide)) {
      sides.get(reversedSide).push(tile.id);
    } else {
      sides.set(side, [tile.id]);
    }
  }
}

for (const [side, tileIDs] of sides.entries()) {
  if (tileIDs.length === 1) {
    const tile = tiles.get(tileIDs[0]);
    const sides = tile.getSides();
    for (let i = 0; i < sides.length; i++) {
      if (sides[i] === side) {
        tile.edgeSides.push(i);
        break;
      }
    }
  } else {
    const [tile1, tile2] = tileIDs.map((id) => tiles.get(id));
    tile1.neighbors.add(tile2);
    tile2.neighbors.add(tile1);
  }
}

const corners = Array.from(tiles.values()).filter(
  (tile) => tile.edgeSides.length === 2
);

const part1 = corners.reduce((acc, tile) => acc * tile.id, 1);
console.log(part1);

const image = [[corners[0]]];
let currentTile = corners[0];

const edgeSidesString = JSON.stringify(currentTile.edgeSides);

// orient corner tile so that edge sides are on top and left
if (edgeSidesString === "[0,1]") {
  currentTile.reverse();
} else if (edgeSidesString === "[1,2]") {
  currentTile.rotate();
  currentTile.rotate();
} else if (edgeSidesString === "[2,3]") {
  currentTile.reverse();
  currentTile.rotate();
}
currentTile.placed = true;

let currentSide = currentTile.getRightSide();
let row = 0;
const tilesPerEdge = Math.sqrt(tiles.size);

function printImage() {
  for (const row of image) {
    for (let i = 0; i < row[0].imageData.length; i++) {
      let line = "";
      for (const tile of row) {
        line += tile.imageData[i] + "|";
      }
      console.log(line);
    }
    console.log();
  }
}

while (true) {
  while (true) {
    let neighborFound = false;
    for (const neighbor of currentTile.neighbors) {
      if (neighbor.placed) {
        currentTile.neighbors.delete(neighbor);
        neighbor.neighbors.delete(currentTile);
        continue;
      }
      let rotations = 0;
      while (currentSide !== neighbor.getLeftSide()) {
        neighbor.rotate();
        rotations++;
        if (rotations === 4) {
          neighbor.reverse();
        } else if (rotations === 8) {
          break;
        }
      }
      if (rotations != 8) {
        neighbor.placed = true;
        currentTile.neighbors.delete(neighbor);
        neighbor.neighbors.delete(currentTile);
        image[row].push(neighbor);
        currentTile = neighbor;
        currentSide = currentTile.getRightSide();
        neighborFound = true;
        break;
      }
    }
    if (!neighborFound) {
      currentTile.print();
      currentTile.neighbors.forEach((neighbor) => neighbor.print());
      throw new Error("No neighbor found");
    }
    if (image[row].length === tilesPerEdge) {
      break;
    }
  }

  if (image.length === tilesPerEdge) {
    break;
  }
  currentTile = image[row][0];
  currentSide = currentTile.getBottomSide();
  row++;
  for (const neighbor of currentTile.neighbors) {
    let rotations = 0;
    while (currentSide !== neighbor.getTopSide()) {
      neighbor.rotate();
      rotations++;
      if (rotations === 4) {
        neighbor.reverse();
      } else if (rotations === 8) {
        break;
      }
    }
    if (rotations != 8) {
      neighbor.placed = true;
      image.push([neighbor]);
      currentTile = neighbor;
      currentSide = currentTile.getRightSide();
      break;
    }
  }
}

console.log("printImage");
printImage();
