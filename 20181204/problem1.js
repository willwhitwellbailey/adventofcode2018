"use strict";
let fs = require('fs')

function getFileContents(filename) {
  try {
    return fs.readFileSync(filename, 'utf8');
  } catch (e) {
    console.log(e);
  };
}

class Nap {
  constructor(begin, end) {
    this.begin = begin;
    this.end = end;
    this.duration = end.getMinutes() - begin.getMinutes();
  }
}

class Shift {
  _getRecordDate(record) {
    let recordPieces = record.split(' ');
    return new Date(`${recordPieces[0].slice(1)} ${recordPieces[1].slice(0, recordPieces[1].length - 1)}`);
  }

  /**
   * record is a string that contains a timestamp in brackets and a guard number
   * [1518-05-01 00:03] Guard #1783 begins shift
   * @param {string} record 
   */
  constructor(record) {
    try {
      this.shiftStart = this._getRecordDate(record);
      this.naps = [];
      this.minutes = new Map();
    } catch (e) {
      console.log('Could not figure out start date...');
      console.log('record:', record);
      console.log('error:', e);
    }
  }

  sleepStart(record) {
    this.napStart = this._getRecordDate(record);
  }

  sleepEnd(record) {
    let nap = new Nap(this.napStart, this._getRecordDate(record));
    this.naps.push(nap);
    for (let i = nap.begin.getMinutes(); i < nap.begin.getMinutes() + nap.duration; i++) {
      if (!this.minutes.get(i)) {
        this.minutes.set(i, 0);
      }
      this.minutes.set(i, this.minutes.get(i) + 1);
    }
  }

  getMinutesAsleep() {
    return this.naps.reduce((sum, nap) => {
      return sum + nap.duration;
    }, 0);
  }

  getMinutesMap() {
    return this.minutes;
  }
}

class Guard {
  constructor(guardNumber) {
    this.guardNumber = guardNumber;
    this.shifts = [];
  }

  addShift(shift) {
    this.shifts.push(shift);
  }

  getMinutesAsleep() {
    return this.shifts.reduce((sum, shift) => {
      return sum + shift.getMinutesAsleep();
    }, 0);
  }

  getMinutesMap() {
    return this.shifts.reduce((minutesMap, shift) => {
      let shiftMap = shift.getMinutesMap();
      shiftMap.forEach((value, minute) => {
        if (!minutesMap.has(minute)) {
          minutesMap.set(minute, 0);
        }
        minutesMap.set(minute, minutesMap.get(minute) + value);
      });
      return minutesMap;
    }, new Map());
  }
}

let inputs = getFileContents('input.txt').split('\r\n').sort();
// inputs.forEach(i => console.log(i));

let guards = new Map();

for (let i = 0; i < inputs.length; i++) {
  if (/Guard/ig.test(inputs[i])) {
    let guardNumber = undefined;
    try {
      guardNumber = parseInt(inputs[i].split(' ')[3].slice(1), 10);
    } catch (e) {
      console.log('Could not figure out guard number...');
      console.log('record:', record);
      console.log('error:', e);
    }
    if (!guards.has(guardNumber)) {
      guards.set(guardNumber, new Guard(guardNumber));
    }
    let newShift = new Shift(inputs[i]);
    for (let j = i + 1; !/Guard/ig.test(inputs[j]) && j < inputs.length; j++) {
      if (/asleep/ig.test(inputs[j])) {
        newShift.sleepStart(inputs[j]);
      }
      if (/wakes/ig.test(inputs[j])) {
        newShift.sleepEnd(inputs[j]);
      }
      i = j;
    }
    guards.get(guardNumber).addShift(newShift);
  }
}

let sleepRecord = []
guards.forEach((value, key) => {
  sleepRecord.push({
    guardNumber: key,
    totalSleepTime: value.getMinutesAsleep()
  });
});
sleepRecord.sort((a, b) => {
  if (a.totalSleepTime < b.totalSleepTime) return -1;
  if (a.totalSleepTime > b.totalSleepTime) return 1;
  return 0;
})

// biggest sleeper is!!:
let biggestSleeper = guards.get(sleepRecord[sleepRecord.length - 1].guardNumber)

// greatest minutes is!!:
let greatestMinute = { 
  minute: -1,
  value: -1
};
biggestSleeper.getMinutesMap().forEach((value, minute) => {
  if (value > greatestMinute.value) {
    greatestMinute = {
      minute,
      value
    };
  }
})

console.log(`guardNumber ${biggestSleeper.guardNumber} slept the most during ${greatestMinute.minute} value (${greatestMinute.value})`);
console.log(`solution: ${biggestSleeper.guardNumber * greatestMinute.minute}`);