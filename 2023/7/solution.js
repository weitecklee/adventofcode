const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split(" "));

const cardStrengths = {
  2: 2,
  3: 3,
  4: 4,
  5: 5,
  6: 6,
  7: 7,
  8: 8,
  9: 9,
  T: 10,
  J: 11,
  Q: 12,
  K: 13,
  A: 14,
};

const cardStrengths2 = {
  J: 1,
  2: 2,
  3: 3,
  4: 4,
  5: 5,
  6: 6,
  7: 7,
  8: 8,
  9: 9,
  T: 10,
  Q: 12,
  K: 13,
  A: 14,
};

class Hand {
  constructor(hand, bid) {
    this.hand = hand;
    this.handSorted = this.sortHand(hand);
    this.bid = Number(bid);
    this.type = this.evalHand(this.handSorted);
    this.type2 = this.evalHand2(this.handSorted);
  }

  sortHand(hand) {
    return hand
      .split("")
      .sort((a, b) => cardStrengths[b] - cardStrengths[a])
      .join("");
  }

  evalHand(hand) {
    if (/(\w)\1{4}/.test(hand)) return 7; // Five of a kind
    if (/(\w)\1{3}/.test(hand)) return 6; // Four of a kind
    if (/(\w)\1{2}(\w)\2|(\w)\3(\w)\4{2}/.test(hand)) return 5; // Full house
    if (/(\w)\1{2}/.test(hand)) return 4; // Three of a kind
    if (/(\w)\1.*?(\w)\2/.test(hand)) return 3; // Two pair
    if (/(\w)\1/.test(hand)) return 2; // One pair
    return 1; // High card
  }

  evalHand2(hand) {
    let hand2 = hand.replace(/J/g, "");
    if (hand2.length === 5) return this.type; // No jokers
    const type = this.evalHand(hand2);
    if (hand2.length === 4) {
      // One joker
      if (type === 6) return 7; // Four of a kind => Five of a kind
      if (type === 4) return 6; // Three of a kind => Four of a kind
      if (type === 3) return 5; // Two pair => Full house
      if (type === 2) return 4; // One pair => Three of a kind
      return 2; // High card => One pair
    }
    if (hand2.length === 3) {
      // Two jokers
      if (type === 4) return 7; // Three of a kind => Five of a kind
      if (type === 2) return 6; // One pair => Four of a kind
      return 4; // High card => Three of a kind
    }
    if (hand2.length === 2) {
      // Three jokers
      if (type === 2) return 7; // One pair => Five of a kind
      return 6; // High card => Four of a kind
    }
    return 7; // Four or all jokers => Five of a kind
  }
}

const hands = input.map(([a, b]) => new Hand(a, b));

hands.sort((a, b) => {
  if (a.type !== b.type) return a.type - b.type;
  for (let i = 0; i < 5; i++) {
    if (a.hand[i] !== b.hand[i])
      return cardStrengths[a.hand[i]] - cardStrengths[b.hand[i]];
  }
  return 0;
});

const part1 = hands.reduce((a, b, i) => a + b.bid * (i + 1), 0);
console.log(part1);

hands.sort((a, b) => {
  if (a.type2 !== b.type2) return a.type2 - b.type2;
  for (let i = 0; i < 5; i++) {
    if (a.hand[i] !== b.hand[i])
      return cardStrengths2[a.hand[i]] - cardStrengths2[b.hand[i]];
  }
  return 0;
});

const part2 = hands.reduce((a, b, i) => a + b.bid * (i + 1), 0);
console.log(part2);
