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
				i, _ := strconv.Atoi(args[0])
				os.Exit(i)
			}
		}
		if command == "echo" {
			fmt.Fprintln(os.Stdout, strings.Join(args, " "))
			return
		}
		op := fmt.Sprintf("%s: command not found", command)

		fmt.Fprintln(os.Stdout, op)
	}
}
