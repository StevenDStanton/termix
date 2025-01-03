package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	printPrompt()
	for {
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
	cleanCommand := strings.TrimSuffix(command, "\n")

	if cleanCommand == "" {
		return nil
	}

	switch cleanCommand {
	case "exit":
		os.Exit(0)
	default:
	}
	printPrompt()
	return nil
}

func printPrompt() {
	fmt.Print("> ")

}
