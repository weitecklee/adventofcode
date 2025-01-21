import * as fs from "fs";
import * as path from "path";

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("")
  .map(Number);

interface CircleNode {
  value: number;
  prev?: CircleNode;
  next?: CircleNode;
}

class Circle {
  current: CircleNode;
  minValue: number;
  maxValue: number;
  cupMap: Map<number, CircleNode>;
  constructor(nums: number[]) {
    this.minValue = Math.min(...nums);
    this.maxValue = Math.max(...nums);
    this.current = { value: nums[0] };
    this.cupMap = new Map([[nums[0], this.current]]);
    let curr = this.current;
    for (let i = 1; i < nums.length; i++) {
      const next: CircleNode = { value: nums[i] };
      this.cupMap.set(nums[i], next);
      curr.next = next;
      next.prev = curr;
      curr = next;
    }
    curr.next = this.current;
    this.current.prev = curr;
  }

  move() {
    const first = this.current.next!;
    let last = first;
    const pickedUp: Set<number> = new Set([first.value]);
    for (let i = 0; i < 2; i++) {
      last = last.next!;
      pickedUp.add(last.value);
    }
    first.prev = undefined;
    this.current.next = last.next!;
    last.next!.prev = this.current;
    last.next = undefined;
    let dest = this.current.value - 1;
    if (dest < this.minValue) dest = this.maxValue;
    while (pickedUp.has(dest)) {
      dest--;
      if (dest < this.minValue) dest = this.maxValue;
    }
    const destCup = this.cupMap.get(dest)!;
    last.next = destCup.next;
    destCup.next!.prev = last;
    destCup.next = first;
    first.prev = destCup;
    this.current = this.current.next!;
  }

  get labels(): string {
    const values: number[] = [];
    let curr = this.cupMap.get(1)!.next!;
    while (curr.value != 1) {
      values.push(curr.value);
      curr = curr.next!;
    }
    return values.join("");
  }
}

const circle = new Circle(input);

for (let i = 0; i < 100; i++) {
  circle.move();
}

console.log(circle.labels);

class Circle1000000 extends Circle {
  constructor(nums: number[]) {
    super(nums);
    let curr = this.current.prev!;
    for (let i = this.maxValue + 1; i <= 1000000; i++) {
      const next: CircleNode = { value: i };
      this.cupMap.set(i, next);
      next.prev = curr;
      curr.next = next;
      curr = next;
    }
    curr.next = this.current;
    this.current.prev = curr;
    this.maxValue = 1000000;
  }
}

const circle1000000 = new Circle1000000(input);
for (let i = 0; i < 10000000; i++) {
  circle1000000.move();
}

console.log(
  circle1000000.cupMap.get(1)!.next!.value *
    circle1000000.cupMap.get(1)!.next!.next!.value
);
