package main

import "fmt"

func main() {
	for day := 1; day <= 12; day++ {
		if day == 12 {
			fmt.Println("12 Drummers Drumming")
		} else if day == 11 {
			fmt.Println("11 pipers piping")
		} else if day == 10 {
			fmt.Println("Ten Lords a Leaping")
		} else if day == 9 {
			fmt.Println("Nine Ladies Dancing")
		} else if day == 8 {
			fmt.Println("Eight Maids a Milking")
		} else if day == 7 {
			fmt.Println("Seven Swans a Swinng")
		} else if day == 6 {
			fmt.Println("Six Geese a Laying")
		} else if day == 5 {
			fmt.Println("Five Golden Rings")
		} else if day == 4 {
			fmt.Println("Four Calling Birds")
		} else if day == 3 {
			fmt.Println("Three French Hens")
		} else if day == 2 {
			fmt.Println("Two Turtle Doves")
		} else if day == 1 {
			fmt.Println("a Partridge in a Pear Tree")
		}
	}
}
