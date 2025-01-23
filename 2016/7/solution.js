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
let count2 = 0;
const re = /(\w)(\w(?<!\1))\2\1/;
const re2 = /\[\w*?(\w)(\w(?<!\1))\2\1\w*?\]/;
const re3 = /\[\w*\]/g;

for (const line of input) {
  if (re.test(line) && !re2.test(line)) {
    count++;
  }
  const brackets = line.match(re3);
  const outsideBrackets = line.replace(re3, "+");
  let counted = false;
  for (const match of brackets) {
    for (let i = 0; i < match.length - 2; i++) {
      if (match[i] === match[i + 2] && match[i] !== match[i + 1]) {
        const re4 = new RegExp(match[i + 1] + match[i] + match[i + 1]);
        if (re4.test(outsideBrackets)) {
          count2++;
          counted = true;
          break;
        }
      }
    }
    if (counted) {
      break;
    }
  }
}

console.log(count);
console.log(count2);
