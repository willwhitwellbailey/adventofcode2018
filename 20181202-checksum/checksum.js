"use strict";
let fs = require('fs')

function getFileContents(filename) {
  try {
    return fs.readFileSync(filename, 'utf8')
  } catch (e) {
    console.log(e);
  };
}

function createMap(s) {
  let m = new Map();

  for (let i = 0; i < s.length; i++) {
    let c = s[i];
    if (m.has(c)) {
      m.set(c, m.get(c) + 1);
    } else {
      m.set(c, 1);
    }
  }

  return m;
}

function evalCharCount(map, count) {
  for (let [k, v] of map) {
    if (v === count) {
      return true;
    }
  }
  return false;
}

let inputMaps = getFileContents("input.txt").split("\n").map(createMap);

let doublesCount = 0, triplesCount = 0
inputMaps.forEach(m => {
  if (evalCharCount(m, 2)) {
    doublesCount++;
  }

  if (evalCharCount(m, 3)) {
    triplesCount++;
  }
})

console.log(`doubles ${doublesCount} || triples ${triplesCount}`);
console.log(`checksum = ${doublesCount * triplesCount}`);
