package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Input() (string, int) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hello!")

	for {
		fmt.Print("\nEnter your name: ")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		fmt.Print("Enter your table number: ")
		tableInput, _ := reader.ReadString('\n')
		tableInput = strings.TrimSpace(tableInput)

		table, err := strconv.Atoi(tableInput)
		if err != nil {
			fmt.Println("\nInvalid input. Please enter a valid number.")
			continue
		}

		return name, table
	}
}
