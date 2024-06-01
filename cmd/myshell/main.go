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
		command := inputArr[0]
		args := inputArr[1:]

		if command == "exit" && len(args) == 1 {
			if args[0] == "0" {
				os.Exit(0)
			}
		}
		if command == "echo" {
			fmt.Fprintln(os.Stdout, strings.Join(args, " "))
			return;
		}
		op := fmt.Sprintf("%s: command not found", command)

		fmt.Fprintln(os.Stdout, op)
	}
}
