package main

import "fmt"

func daysOfXmas(day int) {
	if day == 1 {
		fmt.Print("On the ", day, "st day of Christmas my true love sent to me: \n")
	} else if day == 2 {
		fmt.Print("On the ", day, "nd day of Christmas my true love sent to me: \n")
	} else if day == 3 {
		fmt.Print("On the ", day, "rd day of Christmas my true love sent to me: \n")
	} else {
		fmt.Print("On the ", day, "th day of Christmas my true love sent to me: \n")
	}
	switch day {
	case 12:
		fmt.Println("12 Drummers Drumming")
		fallthrough
	case 11:
		fmt.Println("11 pipers piping")
		fallthrough
	case 10:
		fmt.Println("Ten Lords a Leaping")
		fallthrough
	case 9:
		fmt.Println("Nine Ladies Dancing")
		fallthrough
	case 8:
		fmt.Println("Eight Maids a Milking")
		fallthrough
	case 7:
		fmt.Println("Seven Swans a Swimming")
		fallthrough
	case 6:
		fmt.Println("Six Geese a Laying")
		fallthrough
	case 5:
		fmt.Println("Five Golden Rings")
		fallthrough
	case 4:
		fmt.Println("Four Calling Birds")
		fallthrough
	case 3:
		fmt.Println("Three French Hens")
		fallthrough
	case 2:
		fmt.Println("Two Turtle Doves and a")
		fallthrough
	case 1:
		fmt.Println("a Partridge in a Pear Tree")
	}

}
func main() {
	daysOfXmas(11)
}
