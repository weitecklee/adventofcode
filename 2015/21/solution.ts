import * as fs from "fs";
import * as path from "path";

const puzzleInput = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8");

const shopInput = `Weapons:    Cost  Damage  Armor
Dagger        8     4       0
Shortsword   10     5       0
Warhammer    25     6       0
Longsword    40     7       0
Greataxe     74     8       0

Armor:      Cost  Damage  Armor
Leather      13     0       1
Chainmail    31     0       2
Splintmail   53     0       3
Bandedmail   75     0       4
Platemail   102     0       5

Rings:      Cost  Damage  Armor
Damage +1    25     1       0
Damage +2    50     2       0
Damage +3   100     3       0
Defense +1   20     0       1
Defense +2   40     0       2
Defense +3   80     0       3`
  .split("\n\n")
  .map((a) => a.split("\n"));

class Character {
  hp: number;
  damage: number;
  armor: number;
  cost: number;
  constructor(equipment: Item[]) {
    this.hp = 0;
    this.damage = 0;
    this.armor = 0;
    this.cost = 0;
    for (const item of equipment) {
      this.equip(item);
    }
  }

  equip(item: Item) {
    this.hp += item.hp;
    this.damage += item.damage;
    this.armor += item.armor;
    this.cost += item.cost;
  }

  attack(enemy: Character) {
    enemy.hp -= Math.max(this.damage - enemy.armor, 1);
  }
}

interface Item {
  name: string;
  hp: number;
  cost: number;
  damage: number;
  armor: number;
}

const itemRegex = /^(.*?)\s{2,}(\d+)\s+(\d+)\s+(\d+)$/;

function parseShopInput(shopInput: string[]): Item[] {
  const items: Item[] = [];
  for (const line of shopInput) {
    const parts = line.match(itemRegex);
    if (!parts) continue;
    items.push({
      name: parts[1],
      hp: 0,
      cost: Number(parts[2]),
      damage: Number(parts[3]),
      armor: Number(parts[4]),
    });
  }
  return items;
}

const weapons = parseShopInput(shopInput[0].slice(1));
const armors = parseShopInput(shopInput[1].slice(1));
armors.push({ name: "No Armor", hp: 0, cost: 0, damage: 0, armor: 0 });
const rings = parseShopInput(shopInput[2].slice(1));
rings.push({ name: "No Ring", hp: 0, cost: 0, damage: 0, armor: 0 });

const ringCombos: [Item, Item][] = [];
for (const ring1 of rings) {
  for (const ring2 of rings) {
    ringCombos.push([ring1, ring2]);
  }
}

const playerBase: Item = {
  name: "Player",
  hp: 100,
  cost: 0,
  damage: 0,
  armor: 0,
};

const bossStats = puzzleInput.match(/\d+/g)!.map(Number);

const bossBase: Item = {
  name: "Boss",
  hp: bossStats[0],
  cost: 0,
  damage: bossStats[1],
  armor: bossStats[2],
};

function simulate(player: Character, boss: Character): boolean {
  while (true) {
    player.attack(boss);
    if (boss.hp <= 0) return true;
    boss.attack(player);
    if (player.hp <= 0) return false;
  }
}

const scenarios: [Character, Character][] = [];

for (const weapon of weapons) {
  for (const armor of armors) {
    for (const ringCombo of ringCombos) {
      const player = new Character([playerBase, weapon, armor, ...ringCombo]);
      const boss = new Character([bossBase]);
      scenarios.push([player, boss]);
    }
  }
}

const playerWinCosts: number[] = [];
const playerLossCosts: number[] = [];

for (const [player, boss] of scenarios) {
  const playerWon = simulate(player, boss);
  if (playerWon) playerWinCosts.push(player.cost);
  else playerLossCosts.push(player.cost);
}

console.log(Math.min(...playerWinCosts), Math.max(...playerLossCosts));
