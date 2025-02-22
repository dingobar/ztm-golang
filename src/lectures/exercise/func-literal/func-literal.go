//--Summary:
//  Create a program that can create a report of rune information from
//  lines of text.
//
//--Requirements:
//* Create a single function to iterate over each line of text that is
//  provided in main().
//  - The function must return nothing and must execute a closure
//* Using closures, determine the following information about the text and
//  print a report to the terminal:
//  - Number of letters
//  - Number of digits
//  - Number of spaces
//  - Number of punctuation marks
//
//--Notes:
//* The `unicode` stdlib package provides functionality for rune classification

package main

import (
	"fmt"
	"unicode"
)

type LineCallback func(line string)

func iterate(stringArray []string, op LineCallback) {
	for _, val := range stringArray {
		op(val)
	}
}

func main() {
	lines := []string{
		"There are",
		"68 letters,",
		"five digits,",
		"12 spaces,",
		"and 4 punctuation marks in these lines of text!",
	}

	var (
		letters, numbers, punctuation, spaces int
	)

	countFunc := func(s string) {
		for _, l := range s {
			if unicode.IsLetter(l) {
				letters++
			} else if unicode.IsDigit(l) {
				numbers++
			} else if unicode.IsSpace(l) {
				spaces++
			} else if unicode.IsPunct(l) {
				punctuation++
			}
		}
	}
	iterate(lines, countFunc)
	fmt.Println(
		"There are",
		letters, "letters,",
		numbers, "digits,",
		spaces, "spaces",
		"and", punctuation, "punctuation marks in these lines of text!",
	)
}
