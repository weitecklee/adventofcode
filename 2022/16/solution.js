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
  for (const valveWithFlow of valvesWithFlow) {
    if (openValves.has(valveWithFlow)) continue;
    let newTimeRemaining =
      timeRemaining - valveDistances.get(currValve).get(valveWithFlow) - 1;
    if (newTimeRemaining < 0) {
      part1 = Math.max(part1, total);
      continue;
    }
    let newTotal = total + newTimeRemaining * valveWithFlow.flow;
    let newOpenValves = new Set([...openValves, valveWithFlow]);
    queue.push([newTotal, newTimeRemaining, valveWithFlow, newOpenValves]);
  }
}

const queue2 = [
  [
    0,
    26,
    valveMap.get("AA"),
    26,
    valveMap.get("AA"),
    new Set(),
    ["AA"],
    ["AA"],
  ],
];
let part2 = 0;
const pathSet = new Set();

try {
  while (queue2.length) {
    const [
      total,
      timeRemainingMe,
      currValveMe,
      timeRemainingElephant,
      currValveElephant,
      openValves,
      pathMe,
      pathElephant,
    ] = queue2.pop();
    if (openValves.size === valvesWithFlow.length) {
      part2 = Math.max(part2, total);
      console.log(part2);
      continue;
    }
    const pathKey1 = pathMe.join("-") + " " + pathElephant.join("-");
    const pathKey2 = pathElephant.join("-") + " " + pathMe.join("-");
    if (pathSet.has(pathKey1) || pathSet.has(pathKey2)) continue;
    pathSet.add(pathKey1);
    let progress = false;
    for (const valveWithFlow of valvesWithFlow) {
      if (openValves.has(valveWithFlow)) continue;
      const newPathMe = [...pathMe, valveWithFlow.name];
      let newOpenValves = new Set([...openValves, valveWithFlow]);
      let newTimeRemainingMe =
        timeRemainingMe -
        valveDistances.get(currValveMe).get(valveWithFlow) -
        1;
      if (newTimeRemainingMe > 0) {
        progress = true;
        let newTotal = total + newTimeRemainingMe * valveWithFlow.flow;
        queue2.push([
          newTotal,
          newTimeRemainingMe,
          valveWithFlow,
          timeRemainingElephant,
          currValveElephant,
          newOpenValves,
          newPathMe,
          pathElephant,
        ]);
      }
      const newPathElephant = [...pathElephant, valveWithFlow.name];
      let newTimeRemainingElephant =
        timeRemainingElephant -
        valveDistances.get(currValveElephant).get(valveWithFlow) -
        1;
      if (newTimeRemainingElephant > 0) {
        progress = true;
        let newTotal = total + newTimeRemainingElephant * valveWithFlow.flow;
        queue2.push([
          newTotal,
          timeRemainingMe,
          currValveMe,
          newTimeRemainingElephant,
          valveWithFlow,
          newOpenValves,
          pathMe,
          newPathElephant,
        ]);
      }
    }
    if (!progress) {
      part2 = Math.max(part2, total);
      console.log(part2);
    }
  }
} catch (e) {
  console.log(e);
} finally {
  console.log(part1);
  console.log(part2);
}

// Working(?) solution with brute force.
// Have to wrap it in a try block because it actually never finishes
// on the full input (Set maximum size exceeded for pathSet).
// After a minute or so it does give the correct answer but...
// yeah we'll optimize it later.
