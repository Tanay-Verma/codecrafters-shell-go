package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	command, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
	} else {
		// removing the "/n" from the end of the command and storing the formatted string
		op := fmt.Sprintf("%s: command not found", command[:len(command)-1])
		fmt.Fprintln(os.Stdout, op)
	}
}
