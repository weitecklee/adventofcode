const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8", (err, data) => {
    if (err) {
      console.log(err);
    } else {
      return data;
    }
  })
  .split("\n");

const orbitalObjectMap = new Map();

class OrbitalObject {
  constructor(name) {
    this.name = name;
    this.orbiters = [];
    this.orbits = null;
    this.checked = false;
  }

  addOrbiter(orbitalObj) {
    this.orbiters.push(orbitalObj);
  }

  setOrbits(orbitalObj) {
    this.orbits = orbitalObj;
  }
}

for (const line of input) {
  const parts = line.split(")");
  if (!orbitalObjectMap.has(parts[0])) {
    orbitalObjectMap.set(parts[0], new OrbitalObject(parts[0]));
  }
  if (!orbitalObjectMap.has(parts[1])) {
    orbitalObjectMap.set(parts[1], new OrbitalObject(parts[1]));
  }
  const orbitee = orbitalObjectMap.get(parts[0]);
  const orbiter = orbitalObjectMap.get(parts[1]);
  orbitee.addOrbiter(orbiter);
  orbiter.setOrbits(orbitee);
}

function part1() {
  let res = 0;
  const queue = [...orbitalObjectMap.values()];
  while (queue.length) {
    const curr = queue.pop();
    res += curr.orbiters.length;
    queue.push(...curr.orbiters);
  }
  return res;
}

console.log(part1());

function part2(obj1, obj2) {
  const orbitalObj1 = orbitalObjectMap.get(obj1);
  orbitalObj1.checked = true;
  const queue = [[orbitalObj1, 0]];
  for (let i = 0; i < queue.length; i++) {
    const [currOrbitalObj, d] = queue[i];
    if (currOrbitalObj.name === obj2) {
      return d - 2;
    }
    currOrbitalObj.checked = true;
    if (currOrbitalObj.orbits && !currOrbitalObj.orbits.checked) {
      queue.push([currOrbitalObj.orbits, d + 1]);
    }
    for (const orbitalObj of currOrbitalObj.orbiters) {
      if (!orbitalObj.checked) {
        queue.push([orbitalObj, d + 1]);
      }
    }
  }
}

console.log(part2("YOU", "SAN"));
