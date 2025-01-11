const fs = require("fs");
const path = require("path");

const input = fs
  .readFileSync(path.join(__dirname, "input.txt"), "utf-8")
  .split("\n")
  .map((a) => a.split(" "));

/*
  Actually solved first by hand by going through input and noticing patterns.
  Input is comprised of 14 subprograms (one for each digit) that are
  largely similar except the 5th, 6th, and 16th lines. (1-index)
  5th line is always "div z N" where N is either 1 or 26.
  6th line is always "add x N" where N is either positive or negative number.
  Crucially, whenever 5th line is "div z 1", 6th line is always
  "add x A" where A is a positive number. Whenever 5th line is
  "div z 26", 6th line is always "add x B" where B is a negative number.
  16th line is always "add y N" where N is positive number.
  What ends up happening is that for z to end up equal to 0 at the end of the
  program, there will be constraints placed on pairs of digits (e.g.,
  digits[0] + 4 = digits[13], digits[1] + 8 = digits[12], etc.)
  Figure out the constraints, figure out the maximum and minimum values of
  the digits within the constraints (digits[i] != 0) and there's your answer.

  I can only assume that other inputs have a similar setup with just the lines
  mentioned above differing in the values of the numbers.
*/
