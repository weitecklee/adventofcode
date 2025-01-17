import * as fs from "fs";
import * as path from "path";

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split(": "));

class Die {
  current: number;
  rolls: number;
  constructor() {
    this.current = 0;
    this.rolls = 0;
  }
  roll(): number {
    let sum = 0;
    for (let i = 0; i < 3; i++) {
      this.current++;
      if (this.current > 100) this.current = 1;
      sum += this.current;
    }
    this.rolls += 3;
    return sum;
  }
}

class Player {
  currentSpace: number;
  score: number;
  constructor(startingSpace: number) {
    this.currentSpace = startingSpace;
    this.score = 0;
  }
  advance(spaces: number) {
    this.currentSpace += spaces % 10;
    if (this.currentSpace > 10) this.currentSpace -= 10;
    this.score += this.currentSpace;
  }
}

const die = new Die();
const player1StartingSpace = Number(input[0][1]);
const player2StartingSpace = Number(input[1][1]);
const player1 = new Player(player1StartingSpace);
const player2 = new Player(player2StartingSpace);

let isPlayer1Turn = true;
while (player1.score < 1000 && player2.score < 1000) {
  if (isPlayer1Turn) player1.advance(die.roll());
  else player2.advance(die.roll());
  isPlayer1Turn = !isPlayer1Turn;
}

console.log(Math.min(player1.score, player2.score) * die.rolls);

/*
  After 3 rolls, `n` possible universes with dice sum `s`:
  n = 1, s = 3;
  n = 3, s = 4;
  n = 6, s = 5;
  n = 7, s = 6;
  n = 6, s = 7;
  n = 3, s = 8;
  n = 1, s = 9.
*/

const possibilities: number[][] = [
  [1, 3],
  [3, 4],
  [6, 5],
  [7, 6],
  [6, 7],
  [3, 8],
  [1, 9],
];

// universeMapKey is [player1Score, player1Space, player2Score, player2Space] joined into string
// value is number of universes with that state

let universeMap: Map<string, number> = new Map([
  [[0, player1StartingSpace, 0, player2StartingSpace].join(","), 1],
]);

let player1Wins = 0;
let player2Wins = 0;
isPlayer1Turn = true;

while (universeMap.size) {
  const universeMap2: Map<string, number> = new Map();
  for (const [universeKey, num] of universeMap) {
    const parts = universeKey.split(",");
    const player1Score = Number(parts[0]);
    const player1Space = Number(parts[1]);
    const player2Score = Number(parts[2]);
    const player2Space = Number(parts[3]);

    if (player1Score >= 21) {
      player1Wins += num;
      continue;
    }
    if (player2Score >= 21) {
      player2Wins += num;
      continue;
    }

    for (const [n, s] of possibilities) {
      let player1SpaceNew = player1Space + (isPlayer1Turn ? s : 0);
      if (player1SpaceNew > 10) player1SpaceNew -= 10;
      let player2SpaceNew = player2Space + (isPlayer1Turn ? 0 : s);
      if (player2SpaceNew > 10) player2SpaceNew -= 10;
      const player1ScoreNew =
        player1Score + (isPlayer1Turn ? player1SpaceNew : 0);
      const player2ScoreNew =
        player2Score + (isPlayer1Turn ? 0 : player2SpaceNew);
      const universeKey2 = [
        player1ScoreNew,
        player1SpaceNew,
        player2ScoreNew,
        player2SpaceNew,
      ].join(",");
      universeMap2.set(
        universeKey2,
        (universeMap2.get(universeKey2) || 0) + n * num
      );
    }
  }
  isPlayer1Turn = !isPlayer1Turn;
  universeMap = universeMap2;
}

console.log(Math.max(player1Wins, player2Wins));
