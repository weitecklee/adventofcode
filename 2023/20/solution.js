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
  while (queue.length) {
    const [dst, pulse, src] = queue.shift();
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
  }
}

console.log(lows * highs);
