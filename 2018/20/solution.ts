import * as fs from "fs";
import * as path from "path";

const puzzleInput = fs.readFileSync(path.join(__dirname, "input.txt"), "utf-8");

function condensePath(path: string): string {
  let len = 0;
  while (len !== path.length) {
    len = path.length;
    path = path.replaceAll(/NS|SN|EW|WE/g, "");
  }
  return path;
}

// function convertRegexToPaths(regex: string): string[] {
//   if (/[^NESW]/.test(regex[0])) regex = regex.slice(1, regex.length - 1);
//   regex = condensePath(regex);
//   if (/^[^\(\)\|]*$/.test(regex)) return [regex];
//   let open = 0;
//   const pipeIndices: number[] = [];
//   for (let i = 0; i < regex.length; i++) {
//     if (regex[i] === "(") open++;
//     else if (regex[i] === ")") open--;
//     else if (regex[i] === "|" && open === 0) pipeIndices.push(i);
//   }
//   if (pipeIndices.length) {
//     const segments: string[] = [];
//     segments.push(regex.slice(0, pipeIndices[0]));
//     for (let i = 1; i < pipeIndices.length; i++) {
//       segments.push(regex.slice(pipeIndices[i - 1] + 1, pipeIndices[i]));
//     }
//     segments.push(regex.slice(pipeIndices[pipeIndices.length - 1] + 1));
//     return segments.map(convertRegexToPaths).flat();
//   }

//   let paths: string[][] = [[]];
//   for (let i = 0; i < regex.length; i++) {
//     if (/[A-Z]/.test(regex[i]))
//       paths.forEach((a) => {
//         a.push(regex[i]);
//       });
//     else {
//       let j = i;
//       let open = 1;
//       while (open > 0) {
//         j++;
//         if (regex[j] === "(") open++;
//         else if (regex[j] === ")") open--;
//       }
//       const res = convertRegexToPaths(regex.slice(i, j + 1));
//       const newPaths: string[][] = [];
//       for (const path of paths) {
//         for (const path2 of res) {
//           newPaths.push(path.concat(path2));
//         }
//       }
//       paths = newPaths;
//       i = j;
//     }
//   }
//   return paths.map((a) => a.join(""));
// }

function getLongestPath(regex: string): number {
  if (/[^NESW]/.test(regex[0])) regex = regex.slice(1, regex.length - 1);
  regex = condensePath(regex);
  if (/^[^\(\)\|]*$/.test(regex)) return regex.length;
  let open = 0;
  const pipeIndices: number[] = [];
  for (let i = 0; i < regex.length; i++) {
    if (regex[i] === "(") open++;
    else if (regex[i] === ")") open--;
    else if (regex[i] === "|" && open === 0) pipeIndices.push(i);
  }
  if (pipeIndices.length) {
    const segments: string[] = [];
    segments.push(regex.slice(0, pipeIndices[0]));
    for (let i = 1; i < pipeIndices.length; i++) {
      segments.push(regex.slice(pipeIndices[i - 1] + 1, pipeIndices[i]));
    }
    segments.push(regex.slice(pipeIndices[pipeIndices.length - 1] + 1));
    return Math.max(...segments.map(getLongestPath));
  }

  let res = 0;
  for (let i = 0; i < regex.length; i++) {
    if (/[A-Z]/.test(regex[i])) res++;
    else {
      let j = i;
      let open = 1;
      while (open > 0) {
        j++;
        if (regex[j] === "(") open++;
        else if (regex[j] === ")") open--;
      }
      res += getLongestPath(regex.slice(i, j + 1));
      i = j;
    }
  }
  return res;
}

console.log(getLongestPath(puzzleInput));
