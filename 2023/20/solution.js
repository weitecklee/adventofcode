const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split(" -> "))
  .map(([a, b]) => [a, b.split(", ")]);

class Module {
  constructor(name, destinations) {
    this.name = name;
    this.destinations = destinations;
    this.pulse = false;
  }

  sendPulse() {
    return this.destinations.map((d) => [d, this.pulse, this.name]);
  }
}

class Button extends Module {
  constructor() {
    super("button", ["broadcaster"]);
  }
}

class Broadcaster extends Module {
  constructor(name, destinations) {
    super(name, destinations);
  }
}

class FlipFlop extends Module {
  constructor(name, destinations) {
    super(name, destinations);
    this.state = false;
  }

  sendPulse(srcPulse) {
    let pulses = [];
    if (!srcPulse) {
      this.state = !this.state;
      pulses = this.destinations.map((d) => [d, this.state, this.name]);
    }
    return pulses;
  }
}

class Conjunction extends Module {
  constructor(name, destinations) {
    super(name, destinations);
    this.memory = new Map();
  }

  addSource(srcModuleName) {
    this.memory.set(srcModuleName, false);
  }

  sendPulse(srcPulse, srcModuleName) {
    this.memory.set(srcModuleName, srcPulse);
    if (Array.from(this.memory.values()).every((v) => v)) {
      return this.destinations.map((d) => [d, false, this.name]);
    }
    return this.destinations.map((d) => [d, true, this.name]);
  }
}

const moduleMap = new Map([["button", new Button()]]);
const conjModules = [];

for (const [src, dst] of input) {
  if (src === "broadcaster") {
    moduleMap.set(src, new Broadcaster(src, dst));
  } else {
    const moduleType = src[0];
    const name = src.slice(1);
    if (moduleType === "%") {
      moduleMap.set(name, new FlipFlop(name, dst));
    } else {
      conjModules.push(name);
      moduleMap.set(name, new Conjunction(name, dst));
    }
  }
}

for (const [src, dst] of input) {
  const srcName = src.replaceAll(/[%&]/g, "");
  for (const name of dst) {
    if (conjModules.includes(name)) {
      moduleMap.get(name).addSource(srcName);
    }
  }
}

let lows = 0;
let highs = 0;
for (let i = 0; i < 1000; i++) {
  const queue = moduleMap.get("button").sendPulse();
  let j = 0;
  while (j < queue.length) {
    const [dst, pulse, src] = queue[j];
    if (pulse) {
      highs++;
    } else {
      lows++;
    }
    if (moduleMap.has(dst)) {
      moduleMap
        .get(dst)
        .sendPulse(pulse, src)
        .forEach((p) => queue.push(p));
    }
    j++;
  }
}

console.log(lows * highs);

// Analysis of input shows graph is composed of 4 independent subgraphs,
// branching from 'button' -> 'broadcaster', eventually leading to '&rm' -> 'rx'.
// At end of each subgraph is a conjunction module that periodically sends
// high pulse. Find this period for each subgraph and calculate LCM
// to find when conjunction module pulses will line up to eventually send
// low pulse to 'rx' module.

// Find conjunction module that sends pulse to 'rx'
const conjModuleToRX = moduleMap.get(
  conjModules.find((mdl) => moduleMap.get(mdl).destinations.includes("rx"))
);
// Map to keep track of period for high pulses from each source
const sources = new Map(
  Array.from(conjModuleToRX.memory.keys()).map((k) => [k, 0])
);

let pushes = 1000;
while (Array.from(sources.values()).some((v) => v === 0)) {
  pushes++;
  const queue = moduleMap.get("button").sendPulse();
  let j = 0;
  while (j < queue.length) {
    const [dst, pulse, src] = queue[j];
    if (sources.has(src) && pulse && sources.get(src) === 0) {
      sources.set(src, pushes);
    }
    if (pulse) {
      highs++;
    } else {
      lows++;
    }
    if (moduleMap.has(dst)) {
      moduleMap
        .get(dst)
        .sendPulse(pulse, src)
        .forEach((p) => queue.push(p));
    }
    j++;
  }
}

const periods = Array.from(sources.values());
const gcd = (a, b) => (b === 0 ? a : gcd(b, a % b));
const lcm = (a, b) => (a * b) / gcd(a, b);
console.log(periods.reduce(lcm));
