package main

import (
	"fmt"
	"task1/greet"
	"task1/input"
)

func main() {
	name, table := input.Input()
	res := greet.Greet(name, table)
	fmt.Println(res)
}
