package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	prompt, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Fprint(os.Stdout, err)
	}
	fmt.Fprint(os.Stdout, strings.TrimRight(prompt, "\n"), ": command not found\n")
}
