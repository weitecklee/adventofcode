const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n");

const instructions = input.map((a) => [a.slice(0, 1), Number(a.slice(1))]);

class Ship {
  constructor() {
    this.position = [0, 0];
    this.direction = [1, 0];
  }
  navigate(instruction) {
    const [action, value] = instruction;
    switch (action) {
      case "N":
        this.position[1] += value;
        break;
      case "S":
        this.position[1] -= value;
        break;
      case "E":
        this.position[0] += value;
        break;
      case "W":
        this.position[0] -= value;
        break;
      case "L":
        this.rotate(360 - value);
        break;
      case "R":
        this.rotate(value);
        break;
      case "F":
        this.position[0] += this.direction[0] * value;
        this.position[1] += this.direction[1] * value;
        break;
      default:
        throw new Error("Unknown action: " + action);
    }
  }
  rotate(degrees) {
    if (degrees === 180) {
      this.direction[0] *= -1;
      this.direction[1] *= -1;
    } else if (degrees === 90) {
      this.direction = [this.direction[1], -this.direction[0]];
    } else {
      this.direction = [-this.direction[1], this.direction[0]];
    }
  }
  distance() {
    return Math.abs(this.position[0]) + Math.abs(this.position[1]);
  }
}

class Ship2 extends Ship {
  constructor() {
    super();
    this.waypoint = [10, 1];
  }
  navigate(instruction) {
    const [action, value] = instruction;
    switch (action) {
      case "N":
        this.waypoint[1] += value;
        break;
      case "S":
        this.waypoint[1] -= value;
        break;
      case "E":
        this.waypoint[0] += value;
        break;
      case "W":
        this.waypoint[0] -= value;
        break;
      case "L":
        this.rotate(360 - value);
        break;
      case "R":
        this.rotate(value);
        break;
      case "F":
        this.position[0] += this.waypoint[0] * value;
        this.position[1] += this.waypoint[1] * value;
        break;
      default:
        throw new Error("Unknown action: " + action);
    }
  }
  rotate(degrees) {
    if (degrees === 180) {
      this.waypoint[0] *= -1;
      this.waypoint[1] *= -1;
    } else if (degrees === 90) {
      this.waypoint = [this.waypoint[1], -this.waypoint[0]];
    } else {
      this.waypoint = [-this.waypoint[1], this.waypoint[0]];
    }
  }
}

const ship1 = new Ship();

for (const instruction of instructions) {
  ship1.navigate(instruction);
}

console.log(ship1.distance());

const ship2 = new Ship2();

for (const instruction of instructions) {
  ship2.navigate(instruction);
}

console.log(ship2.distance());
