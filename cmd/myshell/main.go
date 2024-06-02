package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
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
	if err != nil {
		log.Fatal(err)
	}

	input = strings.Trim(input, "\n")

	// removing the "/n" from the end of the command and storing the formatted string
	inputArr := strings.Split(input, " ")
	command := inputArr[0]
	args := inputArr[1:]

	executeCommand(command, args)
}

func executeCommand(command string, args []string) {
	switch command {
	case "exit":
		if len(args) == 1 {
			exitCode, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatal("Invalid exit code")
			} else {
				os.Exit(exitCode)
			}
		}
	case "echo":
		fmt.Fprintln(os.Stdout, strings.Join(args, " "))
		return
	case "type":
		if len(args) == 1 {
			argsCommand := args[0]
			if builtinCommands[argsCommand] {
				op := fmt.Sprintf("%s is a shell builtin", argsCommand)
				fmt.Fprintln(os.Stdout, op)
			} else {
				paths := strings.Split(os.Getenv("PATH"), ":")
				for _, path := range paths {
					fp := filepath.Join(path, argsCommand)
					_, err := os.Stat(fp)

					if err == nil {
						fmt.Fprintln(os.Stdout, argsCommand+" is "+fp)
						return
					}
				}
				op := fmt.Sprintf("%s: not found", argsCommand)
				fmt.Fprintln(os.Stdout, op)
			}
		}
	default:
		_, err := os.Stat(command)
		if err != nil {
			op := fmt.Sprintf("%s: command not found", command)
			fmt.Fprintln(os.Stdout, op)
		} else {
			op, err := exec.Command(command, args...).Output()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Fprintln(os.Stdout, strings.Trim(string(op[:]), "\n"))
			return

		}
	}
}
