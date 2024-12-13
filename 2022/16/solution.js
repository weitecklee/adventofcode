const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n");

const regValve = /[A-Z]{2}/g;
const regFlow = /\d+/;

class Valve {
  constructor(name) {
    this.name = name;
    this.flow = 0;
    this.neighbors = new Set();
  }

  addNeighbor(valve) {
    this.neighbors.add(valve);
  }
}

const valveMap = new Map();
const valvesWithFlow = [];

for (const line of input) {
  const valves = line.match(regValve);
  const flow = Number(line.match(regFlow));
  if (!valveMap.has(valves[0])) {
    valveMap.set(valves[0], new Valve(valves[0]));
  }
  const currValve = valveMap.get(valves[0]);
  currValve.flow = flow;
  if (flow > 0) valvesWithFlow.push(currValve);
  for (let i = 1; i < valves.length; i++) {
    if (!valveMap.has(valves[i])) {
      valveMap.set(valves[i], new Valve(valves[i]));
    }
    currValve.addNeighbor(valveMap.get(valves[i]));
  }
}

const valves = Array.from(valveMap.values());
const valveDistances = new Map();

for (const valve of valves) {
  for (const valveWithFlow of valvesWithFlow) {
    if (valve === valveWithFlow) continue;
    if (!valveDistances.has(valve)) valveDistances.set(valve, new Map());
    if (!valveDistances.has(valveWithFlow))
      valveDistances.set(valveWithFlow, new Map());
    if (valveDistances.get(valve).has(valveWithFlow)) continue;
    if (valve.neighbors.has(valveWithFlow)) {
      valveDistances.get(valve).set(valveWithFlow, 1);
      valveDistances.get(valveWithFlow).set(valve, 1);
    }
    const queue = [[valve, 0, new Set()]];
    let i = 0;
    while (i < queue.length) {
      const [currValve, dist, visited] = queue[i];
      i++;
      if (currValve === valveWithFlow) {
        valveDistances.get(valve).set(valveWithFlow, dist);
        valveDistances.get(valveWithFlow).set(valve, dist);
        break;
      }
      for (const neighbor of currValve.neighbors) {
        if (visited.has(neighbor)) continue;
        queue.push([neighbor, dist + 1, new Set([...visited, neighbor])]);
      }
    }
  }
}

const queue = [[0, 30, valveMap.get("AA"), new Set()]];
let part1 = 0;

while (queue.length) {
  const [total, timeRemaining, currValve, openValves] = queue.pop();
  if (openValves.size === valvesWithFlow.length) {
    part1 = Math.max(part1, total);
    continue;
  }
  let progress = false;
  for (const valveWithFlow of valvesWithFlow) {
    if (openValves.has(valveWithFlow)) continue;
    let newTimeRemaining =
      timeRemaining - valveDistances.get(currValve).get(valveWithFlow) - 1;
    if (newTimeRemaining <= 0) {
      continue;
    }
    progress = true;
    let newTotal = total + newTimeRemaining * valveWithFlow.flow;
    let newOpenValves = new Set([...openValves, valveWithFlow]);
    queue.push([newTotal, newTimeRemaining, valveWithFlow, newOpenValves]);
  }
  if (!progress) {
    part1 = Math.max(part1, total);
  }
}

console.log(part1);

const queue2 = [[0, 26, valveMap.get("AA"), new Set()]];
const paths = [];
const pathSet = new Set();

function pathBitmask(openValves) {
  let bitmask = [];
  for (const valve of valvesWithFlow) {
    bitmask.push(openValves.has(valve) ? "1" : "0");
  }
  return parseInt(bitmask.join(""), 2);
}

while (queue2.length) {
  const [total, timeRemaining, currValve, openValves] = queue2.pop();
  if (!pathSet.has(openValves)) {
    pathSet.add(openValves);
    paths.push([total, pathBitmask(openValves)]);
  }
  if (openValves.size === valvesWithFlow.length) {
    continue;
  }
  for (const valveWithFlow of valvesWithFlow) {
    if (openValves.has(valveWithFlow)) continue;
    let newTimeRemaining =
      timeRemaining - valveDistances.get(currValve).get(valveWithFlow) - 1;
    if (newTimeRemaining <= 0) {
      continue;
    }
    let newTotal = total + newTimeRemaining * valveWithFlow.flow;
    let newOpenValves = new Set([...openValves, valveWithFlow]);
    queue2.push([newTotal, newTimeRemaining, valveWithFlow, newOpenValves]);
  }
}

let part2 = 0;
for (let i = 0; i < paths.length; i++) {
  for (let j = i + 1; j < paths.length; j++) {
    const [total1, bitmask1] = paths[i];
    const [total2, bitmask2] = paths[j];
    if ((bitmask1 & bitmask2) === 0) {
      part2 = Math.max(part2, total1 + total2);
    }
  }
}

console.log(part2);

// New strategy: Map every possible path (including incomplete ones)
// while keeping track of totals. Iterate through all possible pairs,
// using bitmask comparison to make sure they don't share any valves.
// If they don't, add totals together and keep track of the maximum.
