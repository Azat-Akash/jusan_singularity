package main

import (
	"fmt"
	"task1/greet"
)

func main() {
	res := greet.Greet("Adlet", 4)
	fmt.Println(res)
}
