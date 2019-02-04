"use strict";
let fs = require('fs')

function getFileContents(filename) {
  try {
    return fs.readFileSync(filename, 'utf8');
  } catch (e) {
    console.log(e);
  };
}

let input = getFileContents('input.txt').trim();

function removeMatches(polymer) {
  let pass = '';

  for (let i = 0; i < polymer.length; i++) {
    if (Math.abs(polymer.charCodeAt(i) - polymer.charCodeAt(i + 1)) === 32) {
      i++;
    } else {
      pass += polymer[i];
    }
  }

  return pass.length === polymer.length ? pass : removeMatches(pass);
}

function runTests() {
  let tests = [
    ['aA', ''],
    ['Aa', ''],
    ['aaAA', ''],
    ['abAB', 'abAB'],
    ['abcCBa', 'aa'],
    ['ABbaC', 'C'],
    ['dabAcCaCBAcCcaDA', 'dabCBAcaDA'], // example from day 5
    ['nNNnHhnNjyYLlLpPmMnHhNyYSwWovlLVbBEWwtTpPejJnNHhjJiIeLlEcCOhFfjuUeuUEJWVv', 'jLShW'] // snippet of input
  ];
  
  tests.forEach(test => {
    console.log(`testing ${test[0]}:`);
    let testResult = removeMatches(test[0]);
    console.log(`   resulted in '${testResult}'`);
    if (testResult !== test[1]) {
      console.error(`   ... '${test[0]}' did not produce expected outcome of '${test[1]}'`);
    } else {
      console.log(`   !!! '${test[0]}' produced expected outcome of '${test[1]}'`)
    }
  });
}

runTests();
console.log(`\nsolution: polymers leftover - ${removeMatches(input).length}`);
