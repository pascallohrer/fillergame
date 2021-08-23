package client

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func inputLoop(execute func(string) error) error {
	for {
		input := bufio.NewScanner(os.Stdin)
		var command string
		for input.Scan() {
			if err := input.Err(); err != nil {
				return fmt.Errorf("inputLoop: command couldn't be read: %v", err)
			}
			command = input.Text()
			if len(strings.Trim(command, " ")) > 0 {
				break
			}
		}
		err := execute(command)
		if err == fmt.Errorf("User abort") {
			return nil
		}
		if err != nil {
			return fmt.Errorf("inputLoop: couldn't execute the previous command: %v", err)
		}
	}
}

func printUpdate(currentField *field, message ...string) {
	clear()
	if len(message) > 0 {
		fmt.Println(" " + message[0])
	}
	fmt.Println(" --------------------")
	fmt.Println(currentField)
	fmt.Println(" --------------------")
	fmt.Println(currentField.menu())
	fmt.Println(" --------------------")
	if status == "move" {
		fmt.Println(" It's your turn. Enter a digit to make a move.")
	}
}

func clear() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	case "linux":
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
