//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	name  string
	price float64
}

func printShoppingList(shoppingList []Product) {
	var total float64
	fmt.Println("The last item is ", shoppingList[len(shoppingList)-1].name)
	fmt.Printf("There are %d items\n", len(shoppingList))
	for i := 0; i < len(shoppingList); i++ {
		total += shoppingList[i].price
	}
	fmt.Println("The total price is: ", total)
}

func main() {

	var shoppingList [4]Product

	shoppingList[0] = Product{name: "Oatmeal", price: 12.70}
	shoppingList[1] = Product{name: "Bananas", price: 6.90}
	shoppingList[2] = Product{name: "Mangoes", price: 30.10}

	printShoppingList(shoppingList[:])

	shoppingList[3] = Product{name: "Orange Juice", price: 15.49}

	printShoppingList(shoppingList[:])
}
