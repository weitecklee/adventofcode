const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n");

// let part1 = 0;

// for (let c = 0; c < input[0].length; c++) {
//   let rowAfterTilt = 0;
//   for (let r = 0; r < input.length; r++) {
//     if (input[r][c] === "#") {
//       rowAfterTilt = r + 1;
//     } else if (input[r][c] === "O") {
//       part1 += input.length - rowAfterTilt;
//       rowAfterTilt++;
//     }
//   }
// }

// console.log(part1);

class Platform {
  constructor(input) {
    this.platform = input;
    this.blankPlatform = input.map((row) => row.replace(/O/g, "."));
  }

  tiltNorth() {
    let result = this.blankPlatform.map((a) => a.split(""));
    for (let c = 0; c < this.platform[0].length; c++) {
      let rowAfterTilt = 0;
      for (let r = 0; r < this.platform.length; r++) {
        if (this.platform[r][c] === "#") {
          rowAfterTilt = r + 1;
        } else if (this.platform[r][c] === "O") {
          result[rowAfterTilt][c] = "O";
          rowAfterTilt++;
        }
      }
    }
    this.platform = result.map((a) => a.join(""));
  }

  tiltWest() {
    let result = this.blankPlatform.map((a) => a.split(""));
    for (let r = 0; r < this.platform.length; r++) {
      let colAfterTilt = 0;
      for (let c = 0; c < this.platform[0].length; c++) {
        if (this.platform[r][c] === "#") {
          colAfterTilt = c + 1;
        } else if (this.platform[r][c] === "O") {
          result[r][colAfterTilt] = "O";
          colAfterTilt++;
        }
      }
    }
    this.platform = result.map((a) => a.join(""));
  }

  tiltSouth() {
    let result = this.blankPlatform.map((a) => a.split(""));
    for (let c = 0; c < this.platform[0].length; c++) {
      let rowAfterTilt = this.platform.length - 1;
      for (let r = this.platform.length - 1; r >= 0; r--) {
        if (this.platform[r][c] === "#") {
          rowAfterTilt = r - 1;
        } else if (this.platform[r][c] === "O") {
          result[rowAfterTilt][c] = "O";
          rowAfterTilt--;
        }
      }
    }
    this.platform = result.map((a) => a.join(""));
  }

  tiltEast() {
    let result = this.blankPlatform.map((a) => a.split(""));

    for (let r = 0; r < this.platform.length; r++) {
      let colAfterTilt = this.platform[0].length - 1;
      for (let c = this.platform[0].length - 1; c >= 0; c--) {
        if (this.platform[r][c] === "#") {
          colAfterTilt = c - 1;
        } else if (this.platform[r][c] === "O") {
          result[r][colAfterTilt] = "O";
          colAfterTilt--;
        }
      }
    }
    this.platform = result.map((a) => a.join(""));
  }

  cycle() {
    this.tiltNorth();
    this.tiltWest();
    this.tiltSouth();
    this.tiltEast();
  }

  calculateLoad() {
    return this.platform.reduce((acc, row, i) => {
      return (
        acc +
        row.split("").filter((a) => a === "O").length *
          (this.platform.length - i)
      );
    }, 0);
  }
}

const platform = new Platform(input);
platform.tiltNorth();
console.log(platform.calculateLoad());

// function printPlatform(platform) {
//   for (let row of platform) {
//     console.log(row);
//   }
//   console.log("---");
// }

const platform2 = new Platform(input);
const outputHistory = new Map();
const loadHistory = new Map();
for (let i = 1; i < 1000; i++) {
  platform2.cycle();
  const load = platform2.calculateLoad();
  const platformString = platform2.platform.join("");
  if (outputHistory.has(platformString)) {
    const initialOffset = outputHistory.get(platformString);
    const cycleLength = i - initialOffset;
    const index = ((1000000000 - initialOffset) % cycleLength) + initialOffset;
    console.log(loadHistory.get(index));
    break;
  }
  outputHistory.set(platformString, i);
  loadHistory.set(i, load);
}
