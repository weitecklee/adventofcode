import * as fs from "fs";
import * as path from "path";

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n");

const parenRegex = /\([^\()]+?\)/g;
function evaluate1(eqn: string): number {
  let matches = eqn.match(parenRegex);
  while (matches) {
    for (const match of matches) {
      eqn = eqn.replace(
        match,
        evaluate1(match.slice(1, match.length - 1)).toString()
      );
    }
    matches = eqn.match(parenRegex);
  }
  const parts = eqn.split(" ");
  let op = "add";
  let res = 0;
  for (const part of parts) {
    if (part === "+") {
      op = "add";
    } else if (part === "*") {
      op = "mul";
    } else if (op === "add") {
      res += Number(part);
    } else {
      res *= Number(part);
    }
  }
  return res;
}

console.log(input.reduce((a, b) => a + evaluate1(b), 0));

const addRegex = /\d+ \+ \d+/g;
const numRegex = /\d+/g;

function evaluate2(eqn: string): number {
  let matches = eqn.match(parenRegex);
  while (matches) {
    for (const match of matches) {
      eqn = eqn.replace(
        match,
        evaluate2(match.slice(1, match.length - 1)).toString()
      );
    }
    matches = eqn.match(parenRegex);
  }
  matches = eqn.match(addRegex);
  while (matches) {
    for (const match of matches) {
      const [a, _, b] = match.split(" ");
      eqn = eqn.replace(match, (Number(a) + Number(b)).toString());
    }
    matches = eqn.match(addRegex);
  }
  return eqn.match(numRegex)!.reduce((a, b) => a * Number(b), 1);
}

console.log(input.reduce((a, b) => a + evaluate2(b), 0));
