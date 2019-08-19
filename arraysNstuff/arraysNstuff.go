package main

import "fmt"

const gCap int = 5 // number of array elements
var gLen int
var gGroceries [gCap]string

func addGrocery(a string) {
	if gLen < gCap {
		gGroceries[gLen] = a
		gLen++
	} else {
		fmt.Println("Error! We've reached the upper bound of the array!!!")
	}
	listGroceries()
}

func listGroceries() {
	if gLen > 0 {
		fmt.Println("Grocery list is as follows:")
		for i := 0; i < gLen; i++ {
			fmt.Println(gGroceries[i])
		}
	} else {
		fmt.Println("No groceries in the list...")
	}
	if gLen == gCap {
		fmt.Println("The grocery list is full")
	}
}

func main() {
	listGroceries()
	addGrocery("Coffee")
	addGrocery("Milk")
	addGrocery("Pizza")
	addGrocery("Cereal")
	addGrocery("French Fries")
}
