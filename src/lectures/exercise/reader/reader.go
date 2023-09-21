//--Summary:
//  Create an interactive command line application that supports arbitrary
//  commands. When the user enters a command, the program will respond
//  with a message. The program should keep track of how many commands
//  have been ran, and how many lines of text was entered by the user.
//
//--Requirements:
//* When the user enters either "hello" or "bye", the program
//  should respond with a message after pressing the enter key.
//* Whenever the user types a "Q" or "q", the program should exit.
//* Upon program exit, some usage statistics should be printed
//  ('Q' and 'q' do not count towards these statistics):
//  - The number of non-blank lines entered
//  - The number of commands entered
//
//--Notes
//* Use any Reader implementation from the stdlib to implement the program

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	cmdHello = "hello"
	cmdBye   = "bye"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	numCommands := 0
	numLines := 0
	for scanner.Scan() {
		if strings.ToUpper(scanner.Text()) == "Q" {
			break
		} else {
			text := strings.TrimSpace(scanner.Text())
			switch text {
			case cmdHello:
				numCommands += 1
				numLines += 1
				fmt.Println("Hello user!!")
			case cmdBye:
				numCommands += 1
				numLines += 1
				fmt.Println("Goodbye user!!")
			case "":
			default:
				numLines += 1
			}
		}
	}
	fmt.Printf("Total Program executions: %d\n", numCommands)
	fmt.Printf("Total Lines read: %d\n", numLines)
	time.Sleep(time.Second)
	os.Exit(0)
}
