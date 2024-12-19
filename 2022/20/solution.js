const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map(Number);

class Node {
  constructor(val) {
    this.val = val;
    this.prev = null;
    this.next = null;
  }
}

class CircularList {
  constructor(arr) {
    this.list = [];
    this.zeroNode;
    for (const n of arr) {
      const node = new Node(n);
      this.list.push(node);
      if (n === 0) {
        this.zeroNode = node;
      }
    }
    for (let i = 1; i < this.list.length; i++) {
      const curr = this.list[i];
      const prev = this.list[i - 1];
      curr.prev = prev;
      prev.next = curr;
    }
    this.list[0].prev = this.list[this.list.length - 1];
    this.list[this.list.length - 1].next = this.list[0];
  }

  mix(k = 1) {
    for (let i = 0; i < k; i++) {
      for (const node of this.list) {
        node.prev.next = node.next;
        node.next.prev = node.prev;
        const n = node.val % (this.list.length - 1);
        if (n >= 0) {
          let curr = node.next;
          for (let j = 0; j < n; j++) {
            curr = curr.next;
          }
          node.prev = curr.prev;
          node.next = curr;
          node.prev.next = node;
          curr.prev = node;
        } else {
          let curr = node.prev;
          for (let j = 0; j > n; j--) {
            curr = curr.prev;
          }
          node.next = curr.next;
          node.prev = curr;
          node.next.prev = node;
          curr.next = node;
        }
      }
    }
  }

  print(n = this.list.length) {
    const res = [this.zeroNode.val];
    let curr = this.zeroNode.next;
    let i = 0;
    while (curr != this.zeroNode && i < n) {
      res.push(curr.val);
      curr = curr.next;
      i++;
    }
    console.log(res.join(","));
  }

  getNthNumberAfterZero(k) {
    let curr = this.zeroNode;
    for (let i = 0; i < k; i++) {
      curr = curr.next;
    }
    return curr.val;
  }

  get groveCoordinates() {
    return [
      this.getNthNumberAfterZero(1000),
      this.getNthNumberAfterZero(2000),
      this.getNthNumberAfterZero(3000),
    ];
  }

  decrypt(decryptionKey) {
    this.list.forEach((node) => {
      node.val *= decryptionKey;
    });
  }
}

const circularList = new CircularList(input);
circularList.mix();
console.log(circularList.groveCoordinates.reduce((a, b) => a + b));

const decryptionKey = 811589153;
const circularList2 = new CircularList(input);
circularList2.decrypt(decryptionKey);
circularList2.mix(10);

console.log(circularList2.groveCoordinates.reduce((a, b) => a + b));
