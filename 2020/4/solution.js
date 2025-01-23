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

const expectedFields = ["byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid"];
const yearPattern = /^\d{4}$/;
const hgtPattern = /^(\d+)(cm|in)$/;
const hclPattern = /^#[0-9a-f]{6}$/;
const eclPattern = /^(amb|blu|brn|gry|grn|hzl|oth)$/;
const pidPattern = /^\d{9}$/;

class Passport {
  constructor(fields) {
    this.fields = new Map();
    for (const [field, val] of fields) {
      this.fields.set(field, val);
    }
    this.validate1();
    this.validate2();
  }

  validate1() {
    const count = expectedFields.reduce((a, b) => a + this.fields.has(b), 0);
    this.isValid = count === 8 || (count === 7 && !this.fields.has("cid"));
  }

  validate2() {
    this.isValid2 = true;
    if (!this.isValid) {
      this.isValid2 = false;
      return;
    }
    if (!yearPattern.test(this.fields.get("byr"))) {
      this.isValid2 = false;
      return;
    }
    this.fields.set("byr", Number(this.fields.get("byr")));
    if (this.fields.get("byr") < 1920 || this.fields.get("byr") > 2002) {
      this.isValid2 = false;
      return;
    }
    if (!yearPattern.test(this.fields.get("iyr"))) {
      this.isValid2 = false;
      return;
    }
    this.fields.set("iyr", Number(this.fields.get("iyr")));
    if (this.fields.get("iyr") < 2010 || this.fields.get("iyr") > 2020) {
      this.isValid2 = false;
      return;
    }
    if (!yearPattern.test(this.fields.get("eyr"))) {
      this.isValid2 = false;
      return;
    }
    this.fields.set("eyr", Number(this.fields.get("eyr")));
    if (this.fields.get("eyr") < 2020 || this.fields.get("eyr") > 2030) {
      this.isValid2 = false;
      return;
    }
    const match = this.fields.get("hgt").match(hgtPattern);
    if (match === null) {
      this.isValid2 = false;
      return;
    }
    const hgt = Number(match[1]);
    if (match[2] === "cm" && (hgt < 150 || hgt > 193)) {
      this.isValid2 = false;
      return;
    }
    if (match[2] === "in" && (hgt < 59 || hgt > 76)) {
      this.isValid2 = false;
      return;
    }
    if (!hclPattern.test(this.fields.get("hcl"))) {
      this.isValid2 = false;
      return;
    }
    if (!eclPattern.test(this.fields.get("ecl"))) {
      this.isValid2 = false;
      return;
    }
    if (!pidPattern.test(this.fields.get("pid"))) {
      this.isValid2 = false;
      return;
    }
  }
}

if (input[input.length - 1].length > 0) {
  input.push("");
}

const fieldPattern = /(\S+):(\S+)/g;
let i = 0;
const passports = [];

while (i < input.length) {
  const currFields = [];
  while (input[i].length > 0) {
    const matches = input[i].matchAll(fieldPattern);
    for (const match of matches) {
      currFields.push([match[1], match[2]]);
    }
    i++;
  }
  passports.push(new Passport(currFields));
  i++;
}

const part1 = passports.reduce((a, b) => a + b.isValid, 0);
console.log(part1);
const part2 = passports.reduce((a, b) => a + b.isValid2, 0);
console.log(part2);
