/**
 * @author William Provost
 * @version 1.0.0
 * @date 2025-11-24
 * @fileoverview This program sums a set of integers entered by the user.
 */

// variables
let numberCount: number = 0;
let sum: number = 0;
let currentNumber: number = 0;

// get number of integers
numberCount = parseInt(prompt("How many integers will be added: ") || "0");

// loop to get each integer and sum them
for (let i = 1; i <= numberCount; i++) {
  currentNumber = parseInt(prompt("Enter an integer: ") || "0");
  sum += currentNumber;
}

// output result
console.log("The sum is " + sum);

console.log("\nDone.");
