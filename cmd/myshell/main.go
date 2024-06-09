package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	readear := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		prompt, err := readear.ReadString('\n')
		if err != nil {
			fmt.Fprint(os.Stdout, err)
		}

		prompt = strings.TrimRight(prompt, "\n")
		args := strings.Fields(prompt)
		if args[0] == "exit" {
			exitCode, err := strconv.Atoi(args[1])
			if err != nil {
				os.Exit(0)
			}
			os.Exit(exitCode)
		} else if args[0] == "echo" {
			fmt.Fprint(os.Stdout, strings.Join(args[1:], " "), "\n")
		} else if args[0] == "type" {
			paths := strings.Split(os.Getenv("PATH"), ":")
			for _, path := range paths {
				fp := filepath.Join(path, args[1])
				_, err = os.Stat(fp)
				if err != nil {
					fmt.Fprint(os.Stdout, args[1], " is ", fp)
					break
				}
			}
			if args[1] == "exit" || args[1] == "echo" || args[1] == "type" {
				fmt.Fprint(os.Stdout, args[1], " is a shell builtin\n")
			} else {
				fmt.Fprint(os.Stdout, args[1], ": not found\n")
			}
		} else {
			fmt.Fprint(os.Stdout, prompt, ": command not found\n")
		}
	}
}
