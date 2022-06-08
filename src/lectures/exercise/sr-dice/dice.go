//--Summary:
//  Create a program that can perform dice rolls using various configurations
//  of number of dice, number of rolls, and number of sides on the dice.
//  The program should report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these cirsumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must use variables to configure:
//  - number of times to roll the dice
//  - number of dice used in the rolls
//  - number of sides of all the dice (6-sided, 10-sided, etc determined
//    with a variable). All dice must use the same variable for number
//    of sides, so they all have the same number of sides.
//
//--Notes:
//* Use packages from the standard library to complete the project
//* Try using different values for your variables to check the results

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func rollDie(nDieSides int) int {
	return rand.Intn(nDieSides) + 1
}

func rollDice(nDice int, nDieSides int) []int {
	var rolls []int
	for i := 0; i < nDice; i++ {
		rolls = append(rolls, rollDie(nDieSides))
	}
	return rolls
}

func arraySum(num ...int) int {
	result := 0
	for _, n := range num {
		result += n
	}
	return result
}

func arrayDiceRollCharacteristic(num ...int) string {
	sum := arraySum(num...)
	switch {
	case len(num) == 2 && sum == 2:
		return "Snake Eyes"
	case sum == 7:
		return "Lucky Seven"
	case sum%2 == 0:
		return "Even"
	default:
		return "Odd"
	}
}

func describeDiceRoll(num ...int) string {
	rollPrint := "You rolled: "
	for _, roll := range num {
		rollPrint += fmt.Sprintf("[%d]", roll)
	}
	rollPrint += fmt.Sprintf("\nRoll sum: %d", arraySum(num...))
	rollPrint += fmt.Sprintf("\nRoll characteristic: %s", arrayDiceRollCharacteristic(num...))
	return rollPrint
}

func main() {
	nDice := 2      // Number of dice to roll
	nDiceSides := 6 // Number of sides on each die
	rolls := rollDice(nDice, nDiceSides)

	fmt.Println(describeDiceRoll(rolls...))
}
