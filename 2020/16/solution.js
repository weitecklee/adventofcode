const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n\n");

const [fieldsInput, myTicketInput, nearbyTicketsInput] = input;

class Field {
  constructor(inputLine) {
    const [name, fieldRanges] = inputLine.split(": ");
    const ranges = fieldRanges
      .split(" or ")
      .map((range) => range.split("-").map(Number));
    this.name = name;
    this.ranges = ranges;
    this.validFields = new Set();
    this.position = -1;
  }

  isValid(value) {
    return this.ranges.some(([min, max]) => min <= value && value <= max);
  }

  initializeValidFields(numFields) {
    for (let i = 0; i < numFields; i++) {
      this.validFields.add(i);
    }
  }

  removeValidField(field) {
    this.validFields.delete(field);
  }

  setPosition() {
    this.position = this.validFields.values().next().value;
    return this.position;
  }
}

const fields = fieldsInput.split("\n").map((line) => new Field(line));
const myTicket = myTicketInput.split("\n")[1].split(",").map(Number);

const nearbyTickets = nearbyTicketsInput
  .split("\n")
  .slice(1)
  .map((line) => line.split(",").map(Number));

let part1 = 0;
const validTickets = [];

for (const ticket of nearbyTickets) {
  let isValid = true;
  for (const value of ticket) {
    if (!fields.some((field) => field.isValid(value))) {
      part1 += value;
      isValid = false;
    }
  }
  if (isValid) validTickets.push(ticket);
}

console.log(part1);

const numFields = myTicket.length;

for (const field of fields) {
  field.initializeValidFields(numFields);
}

for (const ticket of validTickets) {
  for (let i = 0; i < ticket.length; i++) {
    for (const field of fields) {
      if (!field.isValid(ticket[i])) {
        field.removeValidField(i);
      }
    }
  }
}

let undeterminedFields = fields.slice();
while (undeterminedFields.length) {
  let tmp = [];
  for (const field of undeterminedFields) {
    if (field.validFields.size === 1) {
      const position = field.setPosition();
      for (const otherField of undeterminedFields) {
        otherField.removeValidField(position);
      }
    } else {
      tmp.push(field);
    }
  }
  undeterminedFields = tmp;
}

const departureFields = fields.filter((field) =>
  field.name.startsWith("departure")
);

const part2 = departureFields.reduce(
  (product, field) => product * myTicket[field.position],
  1
);

console.log(part2);
