//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	x1 float64
	y1 float64
	x2 float64
	y2 float64
}

func rectangleArea(r Rectangle) float64 {
	return math.Abs((r.x1 - r.x2) * (r.y1 - r.y2))
}

func main() {
	rectangle := Rectangle{
		x1: 1.54, y1: 1, x2: 20.124, y2: 2.123,
	}

	fmt.Println(rectangleArea(rectangle))
}
