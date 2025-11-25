// Author: William Provost
// Version: 1.0.0
// Date: 2025-11-24
// Fileoverview: This program sums a set of integers entered by the user.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// variable declaration
	var numberCount int
	var sum int = 0
	var currentInput string
	var currentNumber int

	reader := bufio.NewReader(os.Stdin)

	// get number of integers
	fmt.Print("How many integers will be added: ")
	currentInput, _ = reader.ReadString('\n')
	currentInput = strings.TrimSpace(currentInput)
	numberCount, _ = strconv.Atoi(currentInput)

	// loop through each integer
	for i := 1; i <= numberCount; i++ {
		fmt.Print("Enter an integer: ")
		currentInput, _ = reader.ReadString('\n')
		currentInput = strings.TrimSpace(currentInput)
		currentNumber, _ = strconv.Atoi(currentInput)
		sum += currentNumber
	}

	// output result
	fmt.Printf("The sum is %d\n", sum)
	fmt.Println("\nDone.")
}
