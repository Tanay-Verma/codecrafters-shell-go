package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		repl()
	}
}

func repl() {
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	input, err := (bufio.NewReader(os.Stdin).ReadString('\n'))
	input = strings.Trim(input, "\n")
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
	} else {
		// removing the "/n" from the end of the command and storing the formatted string

		inputArr := strings.Split(input, " ")

		if inputArr[0] == "exit" && len(inputArr) == 2 {
			if inputArr[1] == "0" {
				os.Exit(0)
			}
		}
		op := fmt.Sprintf("%s: command not found", inputArr[0][:len(inputArr[0])])

		fmt.Fprintln(os.Stdout, op)
	}
}
