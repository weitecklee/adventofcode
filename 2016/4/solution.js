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

let count = 0;
const realRooms = [];

for (const line of input) {
  const sections = line.match(/(.*?)-(\d+)\[(.*)\]/);
  const charmap = new Map();
  for (const c of sections[1]) {
    if (c !== "-") {
      if (!charmap.has(c)) {
        charmap.set(c, 0);
      }
      charmap.set(c, charmap.get(c) + 1);
    }
  }
  const arr = [];
  for (const [c, n] of charmap) {
    arr.push([c, n]);
  }
  arr.sort((a, b) => {
    if (a[1] === b[1]) {
      return a[0].localeCompare(b[0]);
    } else {
      return b[1] - a[1];
    }
  });
  let check = "";
  for (let i = 0; i < 5; i++) {
    check += arr[i][0];
  }
  if (check === sections[3]) {
    const sectorID = Number(sections[2]);
    count += sectorID;
    realRooms.push([sections[1], sectorID]);
  }
}

console.log(count);

for (const [name, sectorID] of realRooms) {
  const codes = [];
  for (const c of name) {
    if (c === "-") {
      codes.push(32);
    } else {
      let code = c.charCodeAt(0) - 97;
      code = (code + sectorID) % 26;
      codes.push(code + 97);
    }
  }
  // console.log(String.fromCharCode(...codes), sectorID);
  if (String.fromCharCode(...codes) === "northpole object storage") {
    console.log(sectorID);
    break;
  }
}
