//--Summary:
//  Create a program that can read text from standard input and count the
//  number of letters present in the input.
//
//--Requirements:
//* Count the total number of letters in any chosen input
//* The input must be supplied from standard input
//* Input analysis must occur per-word, and each word must be analyzed
//  within a goroutine
//* When the program finishes, display the total number of letters counted
//
//--Notes:
//* Use CTRL+D (Mac/Linux) or CTRL+Z (Windows) to signal EOF, if manually
//  entering data
//* Use `cat FILE | go run ./exercise/sync` to analyze a file
//* Use any synchronization techniques to implement the program:
//  - Channels / mutexes / wait groups

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"unicode/utf8"
)

func checkTerminateCommand(cmd string) bool {
	if strings.ToLower(strings.TrimSpace(cmd)) == "q" {
		return true
	}
	return false
}

func countLetters(word string, wg *sync.WaitGroup, out *int) {
	fmt.Println("Started to count word", word)
	defer wg.Done()
	*out += utf8.RuneCountInString(word)
}

func main() {
	var wg sync.WaitGroup
	scanner := bufio.NewScanner(os.Stdin)
	counter := 0
	for scanner.Scan() {
		input := scanner.Text()
		if checkTerminateCommand(input) {
			break
		}
		for _, word := range strings.Split(input, " ") {
			wg.Add(1)
			go countLetters(word, &wg, &counter)
		}
		wg.Wait()
	}
	fmt.Println("N letters total: ", counter)
	fmt.Println("THanks for playing.")
}
