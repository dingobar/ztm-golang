//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

type Item struct {
	active bool
}

func activateTag(item *Item) {
	item.active = true
}

func deactivateTag(item *Item) {
	item.active = false
}

func checkout(items []Item) {
	for i := range items {
		deactivateTag(&items[i])
	}
}

func main() {
	var items []Item
	for i := 0; i < 5; i++ {
		items = append(items, Item{active: true})
	}

	fmt.Println(items)

	deactivateTag(&items[2])
	fmt.Println(items)

	checkout(items)
	fmt.Println(items)

}
