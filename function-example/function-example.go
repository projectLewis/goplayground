package main

import "fmt"

func max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	fmt.Println(max(15, 10))
}
