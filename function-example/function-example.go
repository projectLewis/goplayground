package main

import "fmt"

func max(i int, j int, k *int) {
	if i > j {
		*k = i
	} else {
		*k = j
	}
}

func main() {
	var c int
	max(15, 10, &c)
	fmt.Println(c)
}
