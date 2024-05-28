package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println("Logs from your program will appear here!")
	KnownCommands := map[string]int{}

	// Uncomment this block to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}

	// remove user enter
	input = strings.TrimRight(input, "\n")

	if _, exists := KnownCommands[input]; !exists {
		fmt.Fprintf(os.Stdout, "%v: command not found", input)
	}

}
