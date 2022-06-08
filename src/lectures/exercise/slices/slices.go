//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

func printAssemblyLine(assemblyLine []int) {
	for _, v := range assemblyLine {
		fmt.Print("[", v, "]")
	}
	fmt.Print("\n")
}

func main() {
	assemblyLine := []int{1, 2, 3}

	printAssemblyLine(assemblyLine)

	assemblyLine = append(assemblyLine, 4, 5)

	printAssemblyLine(assemblyLine)

	printAssemblyLine(assemblyLine[len(assemblyLine)-2:])
}
