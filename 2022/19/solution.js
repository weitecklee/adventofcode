const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n");

class Blueprint {
  constructor(line) {
    const numbers = line.match(/\d+/g).map(Number);
    this.id = numbers[0];
    this.oreRobotCost = numbers[1]; // ore
    this.clayRobotCost = numbers[2]; // ore
    this.obsidianRobotCost = [numbers[3], numbers[4]]; // ore, clay
    this.geodeRobotCost = [numbers[5], numbers[6]]; // ore, obsidian
    this.oreRobotLimit = Math.max(
      this.oreRobotCost,
      this.clayRobotCost,
      this.obsidianRobotCost[0],
      this.geodeRobotCost[0]
    );
    this.clayRobotLimit = this.obsidianRobotCost[1];
    this.obsidianRobotLimit = this.geodeRobotCost[1];
  }

  maximumGeodes(time) {
    let max = 0;
    const tried = new Set();
    const queue = [
      {
        timeLeft: time,
        oreRobots: 1,
        clayRobots: 0,
        obsidianRobots: 0,
        geodeRobots: 0,
        ore: 0,
        clay: 0,
        obsidian: 0,
        geode: 0,
      },
    ];
    while (queue.length) {
      let {
        timeLeft,
        oreRobots,
        clayRobots,
        obsidianRobots,
        geodeRobots,
        ore,
        clay,
        obsidian,
        geode,
      } = queue.pop();
      if (timeLeft === 0) {
        max = Math.max(max, geode);
        continue;
      }
      // Instead of branching every minute, branch on building robots.
      // Find out how long it'll take to build each robot (if possible and necessary)
      // and jump ahead to when robot is built.
      // Robot limits are based on resource cost to build each robot (e.g., if any robot
      // type costs max of N ore, then you never need more than N ore robots) since
      // you can only make one robot per second.
      if (oreRobots && clayRobots && obsidianRobots) {
        let time1 = Math.ceil((this.geodeRobotCost[0] - ore) / oreRobots);
        let time2 = Math.ceil(
          (this.geodeRobotCost[1] - obsidian) / obsidianRobots
        );
        let timeToBuildGeodeRobot = Math.max(0, time1, time2);
        if (timeToBuildGeodeRobot < timeLeft) {
          queue.push({
            timeLeft: timeLeft - timeToBuildGeodeRobot - 1,
            oreRobots,
            clayRobots,
            obsidianRobots,
            geodeRobots: geodeRobots + 1,
            ore:
              ore -
              this.geodeRobotCost[0] +
              oreRobots * (timeToBuildGeodeRobot + 1),
            clay: clay + clayRobots * (timeToBuildGeodeRobot + 1),
            obsidian:
              obsidian -
              this.geodeRobotCost[1] +
              obsidianRobots * (timeToBuildGeodeRobot + 1),
            geode: geode + geodeRobots * (timeToBuildGeodeRobot + 1),
          });
        }
      }
      if (oreRobots && clayRobots && obsidianRobots < this.obsidianRobotLimit) {
        let time1 = Math.ceil((this.obsidianRobotCost[0] - ore) / oreRobots);
        let time2 = Math.ceil((this.obsidianRobotCost[1] - clay) / clayRobots);
        let timeToBuildObsidianRobot = Math.max(0, time1, time2);
        if (timeToBuildObsidianRobot < timeLeft) {
          queue.push({
            timeLeft: timeLeft - timeToBuildObsidianRobot - 1,
            oreRobots,
            clayRobots,
            obsidianRobots: obsidianRobots + 1,
            geodeRobots,
            ore:
              ore -
              this.obsidianRobotCost[0] +
              oreRobots * (timeToBuildObsidianRobot + 1),
            clay:
              clay -
              this.obsidianRobotCost[1] +
              clayRobots * (timeToBuildObsidianRobot + 1),
            obsidian:
              obsidian + obsidianRobots * (timeToBuildObsidianRobot + 1),
            geode: geode + geodeRobots * (timeToBuildObsidianRobot + 1),
          });
        }
      }
      if (oreRobots && clayRobots < this.clayRobotLimit) {
        let time1 = Math.ceil((this.clayRobotCost - ore) / oreRobots);
        let timeToBuildClayRobot = Math.max(0, time1);
        if (timeToBuildClayRobot < timeLeft) {
          queue.push({
            timeLeft: timeLeft - timeToBuildClayRobot - 1,
            oreRobots,
            clayRobots: clayRobots + 1,
            obsidianRobots,
            geodeRobots,
            ore:
              ore - this.clayRobotCost + oreRobots * (timeToBuildClayRobot + 1),
            clay: clay + clayRobots * (timeToBuildClayRobot + 1),
            obsidian: obsidian + obsidianRobots * (timeToBuildClayRobot + 1),
            geode: geode + geodeRobots * (timeToBuildClayRobot + 1),
          });
        }
      }
      if (oreRobots < this.oreRobotLimit) {
        let time1 = Math.ceil((this.oreRobotCost - ore) / oreRobots);
        let timeToBuildOreRobot = Math.max(0, time1);
        if (timeToBuildOreRobot < timeLeft) {
          queue.push({
            timeLeft: timeLeft - timeToBuildOreRobot - 1,
            oreRobots: oreRobots + 1,
            clayRobots,
            obsidianRobots,
            geodeRobots,
            ore:
              ore - this.oreRobotCost + oreRobots * (timeToBuildOreRobot + 1),
            clay: clay + clayRobots * (timeToBuildOreRobot + 1),
            obsidian: obsidian + obsidianRobots * (timeToBuildOreRobot + 1),
            geode: geode + geodeRobots * (timeToBuildOreRobot + 1),
          });
        }
      }
    }
    return max;
  }
  get qualityLevel() {
    return this.id * this.maximumGeodes(24);
  }
}

const blueprints = input.map((a) => new Blueprint(a));
console.log(blueprints.reduce((a, b) => a + b.qualityLevel, 0));
