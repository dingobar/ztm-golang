//--Summary:
//  Create an interactive command line application that supports arbitrary
//  commands. When the user enters a command, the program will respond
//  with a message. The program should keep track of how many commands
//  have been ran, and how many lines of text was entered by the user.
//
//--Requirements:
//* When the user enters either "hello" or "bye", the program
//  should respond with a message after pressing the enter key.
//* Whenever the user types a "Q" or "q", the program should exit.
//* Upon program exit, some usage statistics should be printed
//  ('Q' and 'q' do not count towards these statistics):
//  - The number of non-blank lines entered
//  - The number of commands entered
//
//--Notes
//* Use any Reader implementation from the stdlib to implement the program

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var nLines, nCommands int

	for scanner.Scan() {
		val := scanner.Text()
		if val == "" {
			continue
		}
		val = strings.TrimSpace(val)
		if strings.ToLower(val) == "q" {
			fmt.Println(
				"Exiting, thanks for playing. You entered",
				nLines,
				"lines and executed",
				nCommands,
				"commands.")
			break
		}

		nLines++
		if val == "hello" {
			fmt.Println("And hello to you, young traveller")
			nCommands++
		} else if val == "bye" {
			fmt.Println("Farewell, and godspeed")
			nCommands++
		}

		fmt.Println("You entered: ", "\""+strings.ToLower(val)+"\"")
	}
}
