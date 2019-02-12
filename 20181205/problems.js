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

function inPlace(polymer) {
  for (let i = 0; i < polymer.length; i++) {
    if (Math.abs(polymer.charCodeAt(i) - polymer.charCodeAt(i + 1)) === 32) {
      polymer = polymer.slice(0, i) + polymer.slice(i + 2);
      i -= 2;
    }
  }

  return polymer;
}

function inPlaceWithoutChar(polymer, char) {
  for (let i = 0; i < polymer.length; i++) {
    if (polymer.charCodeAt(i) === char.charCodeAt(0) || polymer.charCodeAt(i) + 32 === char.charCodeAt(0)) {
      polymer = polymer.slice(0, i) + polymer.slice(i + 1);
      i--;
    }
    if (Math.abs(polymer.charCodeAt(i) - polymer.charCodeAt(i + 1)) === 32) {
      polymer = polymer.slice(0, i) + polymer.slice(i + 2);
      i -= 2;
    }
  }

  return polymer;
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
    let testResult = inPlace(test[0]);
    console.log(`   resulted in '${testResult}'`);
    if (testResult !== test[1]) {
      console.error(`   ... '${test[0]}' did not produce expected outcome of '${test[1]}'`);
    } else {
      console.log(`   !!! '${test[0]}' produced expected outcome of '${test[1]}'`)
    }
  });
}

function runTestsWithoutChar() {
  let tests = [
    ['aA', 'b', ''],
    ['Aa', 'a', ''],
    ['aaAA', 'b', ''],
    ['abAB', 'b', ''],
    ['abcCBa', 'a', ''],
    ['abcCBa', 'b', 'aa'],
    ['ABbaC', 'a', 'C'],
    ['ABbaC', 'c', ''],
    ['dabAcCaCBAcCcaDA', 'a', 'dbCBcD'], // example from day 5
    ['dabAcCaCBAcCcaDA', 'c', 'daDA'], // example from day 5
    ['nNNnHhnNjyYLlLpPmMnHhNyYSwWovlLVbBEWwtTpPejJnNHhjJiIeLlEcCOhFfjuUeuUEJWVv', 'a', 'jLShW'] // snippet of input
  ];
  
  tests.forEach(test => {
    console.log(`testing ${test[0]}:`);
    let testResult = inPlaceWithoutChar(test[0], test[1]);
    console.log(`   resulted in '${testResult}'`);
    if (testResult !== test[2]) {
      console.error(`   ... '${test[0]}' did not produce expected outcome of '${test[2]}'`);
    } else {
      console.log(`   !!! '${test[0]}' produced expected outcome of '${test[2]}'`)
    }
  });
}

// runTests();
// runTestsWithoutChar();
// console.log(`\nsolution: polymers leftover (removeMatches) - ${removeMatches(input).length}`);
// console.log(`\nsolution: polymers leftover (inPlace) - ${removeMatches(input).length}`);

function getLetters() {
  let letters = [];
  for (let i = 'a'.charCodeAt(0); i <= 'z'.charCodeAt(0); i++) {
    letters.push(String.fromCharCode(i));
  }
  return letters;
}


let minLenghth = getLetters().reduce((memo, letter) => {
  let s = inPlaceWithoutChar(input, letter);
  console.log(`...info: ${letter}--${s.length}`)
  return Math.min(memo, s.length);
}, input.length);

console.log(`part 2 solution: ${minLenghth}`);
