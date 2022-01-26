package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

var hadError bool

func main() {
	args := os.Args
	if len(args) > 2 {
		fmt.Println("Usage: golox [script]")
		os.Exit(64)
	} else if len(args) == 2 {
		runFile(args[1])
	} else {
		runPrompt()
	}
}

func runFile(path string) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	run(string(bytes))

	if hadError {
		os.Exit(65)
	}
}

func runPrompt() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()
		if len(input) == 0 {
			break
		}
		run(input)
		hadError = false
	}
}

func run(source string) {
	fmt.Println(source)
}

func error(line int, message string) {
	report(line, "", message)
}

func report(line int, where, message string) {
	fmt.Fprintf(os.Stderr, "[line %d] Error %s: %s", line, where, message)
	hadError = true
}
