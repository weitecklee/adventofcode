import * as fs from "fs";
import * as path from "path";
import intcodeGenerator, { IntcodeGenerator } from "../intcode";

const puzzleInput = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split(",")
  .map(Number);

const droid = intcodeGenerator(puzzleInput);

function displayMessage(
  droid: IntcodeGenerator,
  display: boolean = false,
  retVal: number = -9999
): string {
  let message: string[] | string = [];
  while (true) {
    const ret = droid.next();
    if (ret.done) break;
    if (ret.value === retVal) break;
    message.push(String.fromCharCode(ret.value));
  }
  message = message.join("");
  if (display) console.log(message);
  return message;
}

function inputCommand(droid: IntcodeGenerator, command: string): number {
  for (let i = 0; i < command.length; i++) {
    droid.next(command.charCodeAt(i));
  }
  return droid.next(10).value;
}

// Puzzle is essentially a text-based RPG where you go from room to room,
// collecting items along the way. Eventually you will find a
// Pressure-Sensitive Floor where you have to be holding a certain
// combination of items to pass and solve the puzzle. Iterate through the
// combinations until the right one is found.

// Certain items will trigger a failstate if taken. Do not take these.
// These items seem to be the same across inputs. It is the other items
// that are different.
const bannedItems: Set<string> = new Set([
  "giant electromagnet",
  "infinite loop",
  "escape pod",
  "photons",
  "molten lava",
]);

const items: string[] = [];
const itemRegex = /(?<=Items here\:\n\- )[a-z ]+(?=\n)/g;

// temporary hard-coded path through map
// TODO: fix :)
const commands: string[] = [
  "west",
  "west",
  "west",
  "west",
  "west",
  "east",
  "south",
  "west",
  "east",
  "north",
  "east",
  "south",
  "north",
  "east",
  "south",
  "south",
  "west",
  "east",
  "south",
  "north",
  "north",
  "east",
  "north",
  "south",
  "east",
  "south",
  "north",
  "north",
  "east",
];

for (const c of commands) {
  const message = displayMessage(droid);
  const match = message.match(itemRegex);
  if (match && !bannedItems.has(match[0])) {
    items.push(match[0]);
    inputCommand(droid, `take ${match[0]}`);
    displayMessage(droid);
  }
  inputCommand(droid, c);
}

// drop all items in inventory
for (const item of items) {
  displayMessage(droid);
  inputCommand(droid, `drop ${item}`);
}

function generateCombinations(array: string[]): string[][] {
  if (array.length === 1) return [array, []];
  const combos: string[][] = [];
  for (const combo of generateCombinations(array.slice(1))) {
    combos.push(combo.concat(array[0]));
    combos.push(combo);
  }
  return combos;
}

const combinations: string[][] = generateCombinations(items);

for (const combo of combinations) {
  // pick up items in current combination
  for (const item of combo) {
    displayMessage(droid);
    inputCommand(droid, `take ${item}`);
  }

  // go into Pressure-Sensitive Floor to check
  displayMessage(droid);
  inputCommand(droid, "north");

  // if message does not contain "Security Checkpoint", we have passed!
  // Grab the number from the message to solve the puzzle.
  const message = displayMessage(droid);
  if (!/Security Checkpoint/.test(message)) {
    const part1 = message.match(/\d+/)!;
    console.log(part1[0]);
    break;
  }

  inputCommand(droid, "");
  // drop items in current combination
  for (const item of combo) {
    displayMessage(droid);
    inputCommand(droid, `drop ${item}`);
  }
}
