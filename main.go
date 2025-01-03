package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err = processCommand(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func processCommand(command string) error {
	cleanCommand := strings.TrimRight(command, " \t\n\r")

	args := strings.Split(cleanCommand, " ")

	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return os.Chdir("/")
		}
		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)
	default:
	}

	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
