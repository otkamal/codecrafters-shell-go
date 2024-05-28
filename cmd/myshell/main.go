package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var KnownCommands = map[string]int{"exit": 0, "echo": 1, "type": 2}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println("Logs from your program will appear here!")

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
		fn := KnownCommands[cmd]
		fmt.Printf("\n%v\n", fn)
		switch fn {
		case 0:
			DoExit(tokenizedInput[1:])
		case 1:
			DoEcho(tokenizedInput[1:])
		case 2:
			DoType(tokenizedInput[1:])
		default:
			DoRun(tokenizedInput)
			fmt.Println("GOT HERE")
		}
	}
}

func DoExit(params []string) {
	os.Exit(0)
}

func DoEcho(params []string) {
	output := strings.Join(params, " ")
	fmt.Fprintf(os.Stdout, "%v\n", output)
}

func DoType(params []string) {
	item := params[0]
	if _, exists := KnownCommands[item]; exists {
		class := "builtin"
		fmt.Fprintf(os.Stdout, "%v is a shell %v\n", item, class)
	} else {
		env := os.Getenv("PATH")
		paths := strings.Split(env, ":")
		for _, path := range paths {
			exec := path + "/" + item
			if _, err := os.Stat(exec); err == nil {
				fmt.Fprintf(os.Stdout, "%v is %v\n", item, exec)
				return
			}
		}
		fmt.Fprintf(os.Stdout, "%v not found\n", item)
	}
}

func DoRun(params []string) {
	item := params[0]
	env := os.Getenv("PATH")
	paths := strings.Split(env, ":")
	for _, path := range paths {
		executable := path + "/" + item
		if _, err := os.Stat(executable); err == nil {
			out, err := exec.Command(executable, params[1]).Output()
			if err != nil {
				fmt.Println("err: ", err)
			}
			fmt.Println(string(out))
			return
		}
	}
	fmt.Fprintf(os.Stdout, "%v: command not found", item)
}
