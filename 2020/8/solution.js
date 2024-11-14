const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split(" "))
  .map(([a, b]) => [a, Number(b)]);

function runProgram(lineToChange = -1) {
  let acc = 0;
  let i = 0;
  const visited = new Set();
  while (i >= 0 && i < input.length && !visited.has(i)) {
    visited.add(i);
    let [op, arg] = input[i];
    if (i === lineToChange) {
      op = op === "nop" ? "jmp" : "nop";
    }
    switch (op) {
      case "nop":
        i++;
        break;
      case "acc":
        acc += arg;
        i++;
        break;
      case "jmp":
        i += arg;
        break;
    }
  }
  return [acc, i < 0 || i >= input.length];
}

console.log(runProgram()[0]);

const linesToChange = input.reduce((acc, [op, arg], i) => {
  if (op === "jmp" || op === "nop") {
    acc.push(i);
  }
  return acc;
}, []);

for (const lineToChange of linesToChange) {
  const [acc, success] = runProgram(lineToChange);
  if (success) {
    console.log(acc);
    break;
  }
}
