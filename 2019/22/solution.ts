import * as fs from "fs";
import * as path from "path";

const puzzleInput = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n");

const dealWithRegex = /(?<=increment )\d+/;
const cutRegex = /(?<=cut )-?\d+/;

const nCards = 10007;

let deck = Array(nCards)
  .fill(0)
  .map((_, i) => i);

for (const line of puzzleInput) {
  let match: RegExpMatchArray | null;
  if ((match = line.match(dealWithRegex))) {
    const inc = Number(match[0]);
    let i = 0;
    const deck2 = Array(nCards).fill(0);
    for (const card of deck) {
      deck2[i] = card;
      i += inc;
      i %= nCards;
    }
    deck = deck2;
  } else if ((match = line.match(cutRegex))) {
    let n = Number(match[0]);
    if (n < 0) n += nCards;
    deck = deck.slice(n).concat(deck.slice(0, n));
  } else {
    deck.reverse();
  }
}

console.log(deck.indexOf(2019));
