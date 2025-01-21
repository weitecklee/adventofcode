import * as fs from "fs";
import * as path from "path";

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n\n");

function historyKey(player1: Player, player2: Player): string {
  return player1.deck.join(",") + "|" + player2.deck.join(",");
}

class Player {
  deck: number[];
  constructor(deck: number[]) {
    this.deck = deck;
  }

  draw(): number {
    if (this.deck.length === 0) return -1;
    return this.deck.shift()!;
  }

  combat(player2: Player): number {
    while (true) {
      const card1 = this.draw();
      const card2 = player2.draw();
      if (card1 > card2) {
        this.deck.push(card1, card2);
        if (player2.deck.length === 0) return 1;
      } else {
        player2.deck.push(card2, card1);
        if (this.deck.length === 0) return 2;
      }
    }
  }

  recursiveCombat(player2: Player): number {
    const combatHistory: Set<string> = new Set();
    while (true) {
      const currKey = historyKey(this, player2);
      if (combatHistory.has(currKey)) {
        return 1;
      }
      combatHistory.add(currKey);
      const card1 = this.draw();
      const card2 = player2.draw();
      if (this.deck.length >= card1 && player2.deck.length >= card2) {
        const player1Copy = this.copy(card1);
        const player2Copy = player2.copy(card2);
        const winner = player1Copy.recursiveCombat(player2Copy);
        if (winner === 1) this.deck.push(card1, card2);
        else player2.deck.push(card2, card1);
      } else {
        if (card1 > card2) this.deck.push(card1, card2);
        else player2.deck.push(card2, card1);
      }
      if (this.deck.length === 0) return 2;
      if (player2.deck.length === 0) return 1;
    }
  }

  get score(): number {
    return this.deck.reduce((a, b, i) => a + b * (this.deck.length - i), 0);
  }

  copy(n: number): Player {
    return new Player(this.deck.slice(0, n));
  }
}

const player1 = new Player(input[0].split("\n").slice(1).map(Number));
const player2 = new Player(input[1].split("\n").slice(1).map(Number));

let winner = player1.combat(player2);
if (winner === 1) console.log(player1.score);
else console.log(player2.score);

const player1Recur = new Player(input[0].split("\n").slice(1).map(Number));
const player2Recur = new Player(input[1].split("\n").slice(1).map(Number));

winner = player1Recur.recursiveCombat(player2Recur);
if (winner === 1) console.log(player1Recur.score);
else console.log(player2Recur.score);
