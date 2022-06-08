//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

func main() {
	printGreeting("Joe")
	fmt.Println(greetJoe())

	two, anotherTwo := getTwoTwo()
	fmt.Println(addThree(two, anotherTwo, getTwo()))
}

func getTwoTwo() (int, int) {
	return getTwo(), getTwo()
}

func getTwo() int {
	return 2
}

func addThree(a, b, c int) int {
	return a + b + c
}

func greetJoe() string {
	return "Hi, Joe!"
}

func printGreeting(name string) {
	fmt.Printf("Hi there, %s!", name)
}
