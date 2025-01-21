import * as fs from "fs";
import * as path from "path";

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map(Number);

const cardPublicKey = input[0];
const doorPublicKey = input[1];

let subjNumber = 7;

function transform(value: number, subjNumber: number): number {
  return (value * subjNumber) % 20201227;
}

function decrypt(subjNumber: number, publicKey: number): number {
  let loops = 0;
  let value = 1;
  while (value != publicKey) {
    loops++;
    value = transform(value, subjNumber);
  }
  return loops;
}

function encrypt(subjNumber: number, loops: number): number {
  let value = 1;
  for (let i = 0; i < loops; i++) {
    value = transform(value, subjNumber);
  }
  return value;
}

const cardLoops = decrypt(subjNumber, cardPublicKey);
// const doorLoops = decrypt(subjNumber, doorPublicKey);

// console.log(encrypt(cardPublicKey, doorLoops));
console.log(encrypt(doorPublicKey, cardLoops));
