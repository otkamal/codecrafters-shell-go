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
	KnownCommands := map[string]int{"exit": 0}

	for {
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
		tokenizedInput := strings.Split(input, " ")
		cmd := tokenizedInput[0]

		if fn, exists := KnownCommands[cmd]; !exists {
			fmt.Fprintf(os.Stdout, "%v: command not found\n", input)
		} else {
			switch fn {
			case 0:
				DoExit(tokenizedInput[1:])
			}
		}
	}

}

func DoExit(params []string) {
	os.Exit(0)
}
