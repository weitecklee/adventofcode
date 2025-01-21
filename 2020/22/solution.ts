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

  copy(n: number): Player {
    return new Player(this.deck.slice(0, n));
  }
}

const player1 = new Player(input[0].split("\n").slice(1).map(Number));
const player2 = new Player(input[1].split("\n").slice(1).map(Number));

let winner = 0;
while ((winner = player1.play(player2)) === 0) {}
if (winner === 1) console.log(player1.score);
else console.log(player2.score);

function historyKey(player1: Player, player2: Player): string {
  return player1.deck.join(",") + "|" + player2.deck.join(",");
}

function recursiveCombat(player1: Player, player2: Player): number {
  const combatHistory: Set<string> = new Set();
  while (true) {
    const currKey = historyKey(player1, player2);
    if (combatHistory.has(currKey)) {
      return 1;
    }
    combatHistory.add(currKey);
    const card1 = player1.draw();
    const card2 = player2.draw();
    if (player1.deck.length >= card1 && player2.deck.length >= card2) {
      const winner = recursiveCombat(player1.copy(card1), player2.copy(card2));
      if (winner === 1) player1.deck.push(card1, card2);
      else player2.deck.push(card2, card1);
    } else {
      if (card1 > card2) player1.deck.push(card1, card2);
      else player2.deck.push(card2, card1);
    }
    if (player1.deck.length === 0) return 2;
    if (player2.deck.length === 0) return 1;
  }
}

const player1Recur = new Player(input[0].split("\n").slice(1).map(Number));
const player2Recur = new Player(input[1].split("\n").slice(1).map(Number));

winner = recursiveCombat(player1Recur, player2Recur);
if (winner === 1) console.log(player1Recur.score);
else console.log(player2Recur.score);
