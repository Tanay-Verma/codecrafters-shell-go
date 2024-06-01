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
	if err == nil {
		op := fmt.Sprintf("%s: command not found", command[:len(command)-1])
		fmt.Fprintln(os.Stdout, op)
	} else {
		fmt.Fprintln(os.Stdout, err)
	}
}
