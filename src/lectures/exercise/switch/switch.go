//--Summary:
//  Create a program to display a classification based on age.
//
//--Requirements:
//* Use a `switch` statement to print the following:
//  - "newborn" when age is 0
//  - "toddler" when age is 1, 2, or 3
//  - "child" when age is 4 through 12
//  - "teenager" when age is 13 through 17
//  - "adult" when age is 18+

package main

func lifeStage(age int) string {

	switch {
	case age == 0:
		return "newborn"
	case age >= 1 && age < 4:
		return "toddler"
	case age >= 4 && age < 13:
		return "child"
	case age >= 13 && age < 18:
		return "teenager"
	case age > 18:
		return "adult"
	default:
		return "alien"
	}

}

func main() {
	var age int = 10
	println(lifeStage(age))
}
