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

  constructor(line: string, army: Army, id: number) {
    const nums = line.match(/\d+/g)!.map(Number);
    this.nUnits = nums[0];
    this.hitPoints = nums[1];
    this.attackDamage = nums[2];
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
    this.target = undefined;
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

  attackTarget() {
    if (this.target) {
      const unitsLost = Math.floor(
        this.damageToTarget(this.target) / this.target.hitPoints
      );
      this.target.nUnits -= unitsLost;
      if (this.target.nUnits <= 0) {
        this.target.army.groups.delete(this.target);
      }
      this.target.attacker = undefined;
    }
  }
}

class Army {
  groups: Set<Group>;
  name: string;
  enemy?: Army;

  constructor(lines: string[]) {
    this.name = lines[0];
    this.groups = new Set(
      lines.slice(1).map((l, i) => new Group(l, this, i + 1))
    );
  }
}

const immune = new Army(puzzleInput[0]);
const infection = new Army(puzzleInput[1]);
immune.enemy = infection;
infection.enemy = immune;

while (immune.groups.size > 0 && infection.groups.size > 0) {
  const groups = Array.from(immune.groups).concat(Array.from(infection.groups));
  groups.sort((a, b) => {
    if (a.effectivePower != b.effectivePower)
      return b.effectivePower - a.effectivePower;
    return b.initiative - a.initiative;
  });
  for (const group of groups) {
    group.findTarget();
  }
  groups.sort((a, b) => b.initiative - a.initiative);
  for (const group of groups) {
    if (group.nUnits <= 0) continue;
    group.attackTarget();
  }
}

console.log(
  Array.from(immune.groups).reduce((a, b) => a + b.nUnits, 0) +
    Array.from(infection.groups).reduce((a, b) => a + b.nUnits, 0)
);
