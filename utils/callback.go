package utils

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/MohdAhzan/go_CLI_App/apps/todo"
)

// map the command corresponding to the functions
var CallBackMap = map[string]func() error{
	"commandHelp": CommandHelp,
	"cmdExit":     CmdExit,
	"cmdClear":    ClearScreen,
	"todo":        todo.TodoStart,
	"just":        todo.Just,
  "ShowApps":    ShowApps,
}

func ClearScreen() error {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func CmdExit() error {

  os.Exit(0)

	return nil
}

func CommandHelp() error {
	fmt.Println("\nuse <options> to run commands \nuse help for Help")
	return nil
}

func ShowApps() error {
	fmt.Print("\n\nuse <appname> to run Apps\n")
	return nil
}

