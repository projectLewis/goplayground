package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 10; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
	fmt.Println("We Made It!!")
}
