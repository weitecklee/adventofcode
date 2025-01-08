const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n\n");

class Image {
  constructor(algo, image) {
    this.algo = algo;
    this.image = image;
    this.enhanceCount = 0;
  }

  pad() {
    const newImage = [];
    const chr = this.enhanceCount % 2 ? "#" : ".";
    for (let i = 0; i < 2; i++) {
      newImage.push(chr.repeat(this.image[0].length + 4));
    }
    for (const row of this.image) {
      newImage.push(chr + chr + row + chr + chr);
    }
    for (let i = 0; i < 2; i++) {
      newImage.push(chr.repeat(this.image[0].length + 4));
    }
    this.image = newImage;
  }

  enhance() {
    this.pad();
    const newImage = [];
    for (let r = 0; r < this.image.length; r++) {
      const row = [];
      for (let c = 0; c < this.image[0].length; c++) {
        const bin = [];
        for (let i = r - 1; i <= r + 1; i++) {
          for (let j = c - 1; j <= c + 1; j++) {
            bin.push(this.pixelAt(i, j));
          }
        }
        const idx = parseInt(bin.join(""), 2);
        row.push(this.algo[idx]);
      }
      newImage.push(row.join(""));
    }
    this.image = newImage;
    this.enhanceCount++;
  }

  pixelAt(r, c) {
    if (
      r < 0 ||
      c < 0 ||
      r > this.image.length - 1 ||
      c > this.image[0].length - 1
    ) {
      if (this.enhanceCount % 2) return "1";
      return "0";
    }
    return this.image[r][c] === "#" ? "1" : "0";
  }
  get litPixels() {
    let count = 0;
    for (const row of this.image) {
      count += row.replaceAll(".", "").length;
    }
    return count;
  }

  print() {
    for (const row of this.image) {
      console.log(row);
    }
    console.log("");
  }
}

const image = new Image(input[0], input[1].split("\n"));

for (let i = 0; i < 2; i++) {
  image.enhance();
}

console.log(image.litPixels);

for (let i = 2; i < 50; i++) {
  image.enhance();
}

console.log(image.litPixels);
