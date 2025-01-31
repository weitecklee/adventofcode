import * as fs from "fs";
import * as path from "path";

const puzzleInput = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8");

const [nPlayers, nMarbles] = puzzleInput.match(/\d+/g)!.map(Number);

interface Marble {
  value: number;
  prev?: Marble;
  next?: Marble;
}

class Circle {
  currentMarble: Marble;
  constructor() {
    const marble0: Marble = { value: 0 };
    marble0.next = marble0;
    marble0.prev = marble0;
    this.currentMarble = marble0;
  }

  addMarble(n: number): number {
    if (n % 23 === 0) {
      let curr = this.currentMarble;
      for (let i = 0; i < 7; i++) {
        curr = curr.prev!;
      }
      curr.next!.prev = curr.prev;
      curr.prev!.next = curr.next;
      this.currentMarble = curr.next!;
      return curr.value + n;
    }

    const newMarble: Marble = { value: n };
    const place = this.currentMarble.next!;
    place.next!.prev = newMarble;
    newMarble.next = place.next;
    newMarble.prev = place;
    place.next = newMarble;
    this.currentMarble = newMarble;
    return 0;
  }
}

const circle = new Circle();
const scores = Array(nPlayers).fill(0);
for (let i = 1; i <= nMarbles; i++) {
  scores[(i - 1) % nPlayers] += circle.addMarble(i);
}
console.log(Math.max(...scores));

const circle2 = new Circle();
const scores2 = Array(nPlayers).fill(0);
for (let i = 1; i <= nMarbles * 100; i++) {
  scores2[(i - 1) % nPlayers] += circle2.addMarble(i);
}
console.log(Math.max(...scores2));
