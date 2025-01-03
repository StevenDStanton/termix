package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var homeDir = ""

func main() {
	reader := bufio.NewReader(os.Stdin)
	var err error
	homeDir, err = os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	err = os.Chdir(homeDir)
	if err != nil {
		log.Fatal(err)
	}

	for {
		printPath()
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

func printPath() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	if strings.HasPrefix(cwd, homeDir) {
		cwd = "~" + cwd[len(homeDir):]
	}

	fmt.Printf("%s> ", cwd)
}
