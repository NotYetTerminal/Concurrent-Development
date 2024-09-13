package main

import (
	"fmt"
	"strings"
)

// Program to convert Roman number to decimal integer number
func main() {
	conversionMap := map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	modifierMap := map[uint8]uint8{
		'V': 'I',
		'X': 'I',
		'L': 'X',
		'C': 'X',
		'D': 'C',
		'M': 'C',
	}
	var romanNumber string

	fmt.Print("Type in your Roman number to convert to decimal number: ")
	_, err := fmt.Scanln(&romanNumber)
	if err != nil {
		fmt.Println("Improper input.\nOnly input the following characters with no spaces: I V X L C D M")
		return
	}
	romanNumber = strings.ToUpper(romanNumber)

	integerNumber := 0
	var previousRomanNumber uint8

	// Loop over string from the back to front
	for i := len(romanNumber) - 1; i > -1; i-- {
		currentRomanNumber := romanNumber[i]
		value, present := modifierMap[previousRomanNumber]

		// If the previous Roman numeral is the key in the modifierMap and the current Roman numeral is the value
		// then the value is subtracted
		// else the number is added
		if present && value == currentRomanNumber {
			integerNumber -= conversionMap[currentRomanNumber]
		} else {
			integerNumber += conversionMap[currentRomanNumber]
		}
		previousRomanNumber = currentRomanNumber
	}

	fmt.Println("The Roman number converted to decimal number is:", integerNumber)
}
