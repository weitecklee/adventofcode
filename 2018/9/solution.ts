import * as fs from "fs";
import * as path from "path";

const puzzleInput = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8");

const [nPlayers, nMarbles] = puzzleInput.match(/\d+/g)!.map(Number);

class Marble {
  value: number;
  next: Marble;
  prev: Marble;
  constructor(value: number) {
    this.value = value;
    this.next = this;
    this.prev = this;
  }
}

class Circle {
  currentMarble: Marble;
  constructor() {
    this.currentMarble = new Marble(0);
  }

  addMarble(n: number): number {
    if (n % 23 === 0) {
      let curr = this.currentMarble;
      for (let i = 0; i < 7; i++) {
        curr = curr.prev;
      }
      curr.next.prev = curr.prev;
      curr.prev.next = curr.next;
      this.currentMarble = curr.next;
      return curr.value + n;
    }

    const newMarble = new Marble(n);
    const place = this.currentMarble.next;
    place.next.prev = newMarble;
    newMarble.next = place.next;
    newMarble.prev = place;
    place.next = newMarble;
    this.currentMarble = newMarble;
    return 0;
  }
}

function play(nPlayers: number, nMarbles: number): number {
  const circle = new Circle();
  const scores = Array(nPlayers).fill(0);
  for (let i = 1; i <= nMarbles; i++) {
    scores[(i - 1) % nPlayers] += circle.addMarble(i);
  }
  return Math.max(...scores);
}

console.log(play(nPlayers, nMarbles));
console.log(play(nPlayers, nMarbles * 100));
