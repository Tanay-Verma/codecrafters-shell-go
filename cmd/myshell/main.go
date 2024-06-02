package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var builtinCommands map[string]func([]string)

func init() {
	builtinCommands = map[string]func([]string){"echo": echo, "exit": exit, "type": typeCommand, "pwd": pwd, "cd": cd}
}

func main() {
	for {
		repl()
	}
}

func repl() {
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	input, err := (bufio.NewReader(os.Stdin).ReadString('\n'))
	if err != nil {
		log.Println("Error reading input:", err)
		return
	}

	input = strings.TrimSpace(input)
	if input == "" {
		return
	}
	// removing the "/n" from the end of the command and storing the formatted string
	inputArr := strings.Split(input, " ")
	command := inputArr[0]
	args := inputArr[1:]

	if handler, exists := builtinCommands[command]; exists {
		handler(args)
	} else {
		executeExternalCommand(command, args)
	}
}

func echo(args []string) {
	fmt.Fprintln(os.Stdout, strings.Join(args, " "))
}

func pwd(args []string) {
	if len(args) != 0 {
		log.Println("Usage pwd")
		return
	}
	dir, err := os.Getwd()
	if err != nil {
		log.Println("Error get current directory:", err)
		return
	}
	fmt.Fprintln(os.Stdout, dir)
}

func cd(args []string) {
	if len(args) != 1 {
		log.Println("Usage: cd <directory>")
		return
	}
	err := os.Chdir(args[0])
	if err != nil {
		fmt.Fprintln(os.Stdout, args[0]+": No such file or directory")
	}
}

func exit(args []string) {
	if len(args) != 1 {
		log.Println("Usage: exit <exit code>")
		return
	}
	exitCode, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal("Invalid exit code")
		return
	}
	os.Exit(exitCode)
}

func typeCommand(args []string) {
	if len(args) != 1 {
		log.Println("Usage: type <command>")
		return
	}
	argsCommand := args[0]
	if _, exists := builtinCommands[argsCommand]; exists {
		op := fmt.Sprintf("%s is a shell builtin", argsCommand)
		fmt.Fprintln(os.Stdout, op)
	} else {
		path, err := exec.LookPath(argsCommand)
		if err != nil {
			fmt.Fprintln(os.Stdout, argsCommand+": not found")
		} else {
			fmt.Fprintln(os.Stdout, argsCommand+" is "+path)
		}
	}
}

func executeExternalCommand(command string, args []string) {
	cmd := exec.Command(command, args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		if _, ok := err.(*exec.ExitError); ok {
			return
		}
		fmt.Fprintln(os.Stdout, command+": command not found")
	}
}
