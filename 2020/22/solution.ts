import * as fs from "fs";
import * as path from "path";

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n\n");

class Player {
  deck: number[];
  constructor(deck: number[]) {
    this.deck = deck;
  }

  draw(): number {
    if (this.deck.length === 0) return -1;
    return this.deck.shift()!;
  }

  play(player2: Player): number {
    const card1 = this.draw();
    const card2 = player2.draw();
    if (card1 > card2) {
      this.deck.push(card1, card2);
      if (player2.deck.length === 0) return 1;
    } else {
      player2.deck.push(card2, card1);
      if (this.deck.length === 0) return 2;
    }
    return 0;
  }

  get score(): number {
    return this.deck.reduce((a, b, i) => a + b * (this.deck.length - i), 0);
  }
}

const player1 = new Player(input[0].split("\n").slice(1).map(Number));
const player2 = new Player(input[1].split("\n").slice(1).map(Number));

let winner = 0;
while ((winner = player1.play(player2)) === 0) {}
if (winner === 1) console.log(player1.score);
else console.log(player2.score);
