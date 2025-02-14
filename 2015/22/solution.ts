import * as fs from "fs";
import * as path from "path";
import MinHeap from "../../utils/MinHeap";

const puzzleInput = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8");

const bossStats = puzzleInput.match(/\d+/g)!.map(Number);

interface Spell {
  name: string;
  costMP: number;
}

const spells: Spell[] = [
  { name: "Magic Missile", costMP: 53 },
  { name: "Drain", costMP: 73 },
  { name: "Shield", costMP: 113 },
  { name: "Poison", costMP: 173 },
  { name: "Recharge", costMP: 229 },
];

interface QueueEntry {
  playerHP: number;
  bossHP: number;
  playerMP: number;
  spentMP: number;
  effects: Map<string, number>;
  playerTurn: boolean;
  armor: number;
}

function solve(isPart2: boolean = false): number {
  const queue: [number, QueueEntry][] = [
    [
      bossStats[0],
      {
        playerHP: 50,
        bossHP: bossStats[0],
        playerMP: 500,
        spentMP: 0,
        effects: new Map(),
        playerTurn: true,
        armor: 0,
      },
    ],
  ];

  while (queue.length) {
    let [
      _,
      { playerHP, bossHP, playerMP, spentMP, effects, playerTurn, armor },
    ] = MinHeap.pop(queue) as [number, QueueEntry];

    if (isPart2 && playerTurn) {
      playerHP--;
      if (playerHP <= 0) continue;
    }

    for (let [name, timer] of effects) {
      timer--;
      switch (name) {
        case "Shield":
          if (timer === 0) {
            armor -= 7;
          }
          break;
        case "Poison":
          bossHP -= 3;
          break;
        case "Recharge":
          playerMP += 101;
          break;
      }
      if (timer === 0) {
        effects.delete(name);
      } else {
        effects.set(name, timer);
      }
    }

    if (bossHP <= 0) {
      return spentMP;
    }

    if (playerTurn) {
      for (const spell of spells) {
        if (effects.has(spell.name)) continue;
        if (playerMP < spell.costMP) continue;
        const effects2 = new Map(effects);
        let playerHP2 = playerHP;
        let bossHP2 = bossHP;
        let armor2 = armor;
        switch (spell.name) {
          case "Magic Missile":
            bossHP2 -= 4;
            break;
          case "Drain":
            bossHP2 -= 2;
            playerHP2 += 2;
            break;
          case "Shield":
            effects2.set("Shield", 6);
            armor2 += 7;
            break;
          case "Poison":
            effects2.set("Poison", 6);
            break;
          case "Recharge":
            effects2.set("Recharge", 5);
            break;
        }
        MinHeap.push(queue, [
          bossHP2 + spentMP + spell.costMP,
          {
            playerHP: playerHP2,
            bossHP: bossHP2,
            playerMP: playerMP - spell.costMP,
            spentMP: spentMP + spell.costMP,
            effects: effects2,
            playerTurn: !playerTurn,
            armor: armor2,
          },
        ]);
      }
    } else {
      const playerHP2 = playerHP - Math.max(1, bossStats[1] - armor);
      if (playerHP2 <= 0) continue;
      MinHeap.push(queue, [
        bossHP + spentMP,
        {
          playerHP: playerHP2,
          bossHP: bossHP,
          playerMP: playerMP,
          spentMP: spentMP,
          effects: new Map(effects),
          playerTurn: !playerTurn,
          armor: armor,
        },
      ]);
    }
  }

  return -1;
}

console.log(solve());
console.log(solve(true));
