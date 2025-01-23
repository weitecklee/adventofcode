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

const connections = new Map();

for (const line of input) {
  const [a, b] = line.split("-");
  if (!connections.has(a)) {
    connections.set(a, []);
  }
  if (!connections.has(b)) {
    connections.set(b, []);
  }
  connections.get(a).push(b);
  connections.get(b).push(a);
}

const q = [["start", new Set(["start"])]];
let res = 0;

for (const [cave, visited] of q) {
  for (const connection of connections.get(cave)) {
    if (connection === "end") {
      res++;
      continue;
    }
    if (!visited.has(connection)) {
      const newVisited = new Set(visited);
      if (connection === connection.toLowerCase()) {
        newVisited.add(connection);
      }
      q.push([connection, newVisited]);
    }
  }
}

console.log(res);

const q2 = [["start", new Set(["start"]), ""]];
res = 0;

for (const [cave, visited, twice] of q2) {
  for (const connection of connections.get(cave)) {
    if (connection === "end") {
      res++;
      continue;
    }
    if (!visited.has(connection)) {
      const newVisited = new Set(visited);
      if (connection === connection.toLowerCase()) {
        newVisited.add(connection);
      }
      q2.push([connection, newVisited, twice]);
    } else if (twice === "" && connection !== "start") {
      const newVisited = new Set(visited);
      q2.push([connection, newVisited, connection]);
    }
  }
}

console.log(res);
