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

class Bot {
  constructor(giveLow, giveHigh) {
    this.lowBot = null;
    this.highBot = null;
    this.lowOutput = null;
    this.highOutput = null;
    this.values = [];
    if (giveLow[0] === "bot") {
      this.lowBot = giveLow[1];
    } else {
      this.lowOutput = giveLow[1];
    }
    if (giveHigh[0] === "bot") {
      this.highBot = giveHigh[1];
    } else {
      this.highOutput = giveHigh[1];
    }
  }
  execute(botMap, outputMap) {
    const res = [];
    if (this.values.length === 2) {
      let lowValue = this.values[0];
      let highValue = this.values[1];
      if (lowValue > highValue) {
        [lowValue, highValue] = [highValue, lowValue];
      }
      if (this.lowBot !== null) {
        botMap.get(this.lowBot).values.push(lowValue);
        if (botMap.get(this.lowBot).values.length === 2) {
          res.push(this.lowBot);
        }
      } else {
        outputMap.set(this.lowOutput, lowValue);
      }
      if (this.highBot !== null) {
        botMap.get(this.highBot).values.push(highValue);
        if (botMap.get(this.highBot).values.length === 2) {
          res.push(this.highBot);
        }
      } else {
        outputMap.set(this.highOutput, highValue);
      }
    }
    return res;
  }
}

const botMap = new Map();
const outputMap = new Map();
const re = /\w+ \d+/g;
const valuesToGiveOut = [];

for (const line of input) {
  const parts = line.match(re).map((a) => {
    a = a.split(" ");
    a[1] = Number(a[1]);
    return a;
  });
  if (parts.length === 3) {
    botMap.set(parts[0][1], new Bot(parts[1], parts[2]));
  } else {
    valuesToGiveOut.push([parts[0][1], parts[1][1]]);
  }
}

const botQueue = [];

for (const [value, bot] of valuesToGiveOut) {
  botMap.get(bot).values.push(value);
  if (botMap.get(bot).values.length === 2) {
    botQueue.push(bot);
  }
}

for (let i = 0; i < botQueue.length; i++) {
  const botsToQueue = botMap.get(botQueue[i]).execute(botMap, outputMap);
  botQueue.push(...botsToQueue);
}

for (const [n, bot] of botMap) {
  if (bot.values.includes(61) && bot.values.includes(17)) {
    console.log(n);
    break;
  }
}

let part2 = 1;
for (let i = 0; i < 3; i++) {
  part2 *= outputMap.get(i);
}

console.log(part2);
