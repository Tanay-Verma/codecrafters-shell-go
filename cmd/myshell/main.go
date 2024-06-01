package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	for {
		repl()
	}
}

var builtinCommands = map[string]bool{"echo": true, "exit": true, "type": true}

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

		switch command {
		case "exit":
			if len(args) == 1 {
				i, err := strconv.Atoi(args[0])
				if err != nil {
				} else {
					os.Exit(i)
				}
			}
		case "echo":
			fmt.Fprintln(os.Stdout, strings.Join(args, " "))
			return
		case "type":
			if len(args) == 1 {
				command = args[0]
				if builtinCommands[command] {
					op := fmt.Sprintf("%s is a shell builtin", command)
					fmt.Fprintln(os.Stdout, op)
					return
				} else {
					op := fmt.Sprintf("%s not found", command)
					fmt.Fprintln(os.Stdout, op)
					return
				}
			}
		default:
			op := fmt.Sprintf("%s: command not found", command)
			fmt.Fprintln(os.Stdout, op)
		}
	}
}
