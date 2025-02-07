import * as fs from "fs";
import * as path from "path";

const puzzleInput = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n\n")
  .map((a) => a.split("\n"));

class Group {
  nUnits: number;
  hitPoints: number;
  immunities: Set<string>;
  weaknesses: Set<string>;
  initiative: number;
  attackDamage: number;
  attackType: string;
  army: Army;
  target?: Group;
  attacker?: Group;
  id: number;

  constructor(line: string, army: Army, id: number, boost: number) {
    const nums = line.match(/\d+/g)!.map(Number);
    this.nUnits = nums[0];
    this.hitPoints = nums[1];
    this.attackDamage = nums[2] + boost;
    this.initiative = nums[3];
    const immMatch = line.match(/immune to (.*?)(;|\))/);
    if (immMatch) {
      this.immunities = new Set(immMatch[1].split(", "));
    } else {
      this.immunities = new Set();
    }
    const weakMatch = line.match(/weak to (.*?)(;|\))/);
    if (weakMatch) {
      this.weaknesses = new Set(weakMatch[1].split(", "));
    } else {
      this.weaknesses = new Set();
    }
    this.attackType = line.match(/\w+(?= damage)/)![0];
    this.army = army;
    this.id = id;
  }

  get effectivePower(): number {
    return this.nUnits * this.attackDamage;
  }

  damageToTarget(target: Group): number {
    let damage = this.effectivePower;
    if (target.weaknesses.has(this.attackType)) damage *= 2;
    return damage;
  }

  findTarget() {
    let maxDamage = 0;
    for (const group of this.army.enemy!.groups) {
      if (group.attacker) continue;
      if (group.immunities.has(this.attackType)) continue;
      const damage = this.damageToTarget(group);
      if (!this.target) {
        maxDamage = damage;
        this.target = group;
        group.attacker = this;
      } else if (damage > maxDamage) {
        maxDamage = damage;
        this.target!.attacker = undefined;
        this.target = group;
        group.attacker = this;
      } else if (damage === maxDamage) {
        if (group.effectivePower > this.target!.effectivePower) {
          this.target!.attacker = undefined;
          this.target = group;
          group.attacker = this;
        } else if (
          group.effectivePower === this.target!.effectivePower &&
          group.initiative > this.target!.initiative
        ) {
          this.target!.attacker = undefined;
          this.target = group;
          group.attacker = this;
        }
      }
    }
  }

  attackTarget(): number {
    let unitsLost = 0;
    if (this.target) {
      unitsLost = Math.min(
        Math.floor(this.damageToTarget(this.target) / this.target.hitPoints),
        this.target.nUnits
      );
      this.target.nUnits -= unitsLost;
      if (this.target.nUnits <= 0) {
        this.target.army.groups.delete(this.target);
      }
    }
    return unitsLost;
  }

  reset() {
    this.target = undefined;
    this.attacker = undefined;
  }
}

class Army {
  groups: Set<Group>;
  name: string;
  enemy?: Army;

  constructor(lines: string[], boost: number = 0) {
    this.name = lines[0];
    this.groups = new Set(
      lines.slice(1).map((l, i) => new Group(l, this, i + 1, boost))
    );
  }

  get nUnits(): number {
    return Array.from(this.groups).reduce((a, b) => a + b.nUnits, 0);
  }
}

function reset(...armies: Army[]) {
  for (const army of armies) {
    for (const group of army.groups) {
      group.reset();
    }
  }
}

function battle(boost: number = 0): Army | null {
  const army1 = new Army(puzzleInput[0], boost);
  const army2 = new Army(puzzleInput[1]);
  army1.enemy = army2;
  army2.enemy = army1;
  while (army1.groups.size > 0 && army2.groups.size > 0) {
    reset(army1, army2);
    const groups = Array.from(army1.groups).concat(Array.from(army2.groups));
    groups.sort((a, b) => {
      if (a.effectivePower != b.effectivePower)
        return b.effectivePower - a.effectivePower;
      return b.initiative - a.initiative;
    });
    for (const group of groups) {
      group.findTarget();
    }
    groups.sort((a, b) => b.initiative - a.initiative);
    let unitsLost = 0;
    for (const group of groups) {
      if (group.nUnits <= 0) continue;
      unitsLost += group.attackTarget();
    }
    if (unitsLost === 0) {
      // stalemate
      return null;
    }
  }
  if (army1.groups.size) return army1;
  return army2;
}

console.log(battle()!.nUnits);

let i = 0;
while (true) {
  i++;
  const winner = battle(i);
  if (winner?.name === "Immune System:") {
    console.log(winner?.nUnits);
    break;
  }
}
