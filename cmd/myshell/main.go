package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var BuiltIns = map[string]int{
	"exit": 1,
	"echo": 2,
	"type": 3,
	"pwd":  4,
	"cd":   5,
}

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
		fn := BuiltIns[cmd]

		switch fn {
		case 1:
			DoExit(tokenizedInput[1:])
		case 2:
			DoEcho(tokenizedInput[1:])
		case 3:
			DoType(tokenizedInput[1:])
		case 4:
			DoPwd(tokenizedInput[1:])
		case 5:
			DoCd(tokenizedInput[1:])
		default:
			DoRun(tokenizedInput)
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
	if _, exists := BuiltIns[item]; exists {
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
	if _, err := os.Stat(item); err == nil {
		out, err := exec.Command(item, params[1]).Output()
		if err != nil {
			fmt.Println("err: ", err)
		}
		fmt.Printf("%v", string(out))
		return
	}
	fmt.Fprintf(os.Stdout, "%v: command not found\n", item)
}

func DoPwd(params []string) {
	currentDirectory, err := os.Getwd()
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Fprintf(os.Stdout, "%v\n", currentDirectory)
}

func DoCd(params []string) {
	err := os.Chdir(params[0])
	if err != nil {
		fmt.Fprintf(os.Stdout, "cd: %v: No such file or directory\n", params[0])
	}
}
