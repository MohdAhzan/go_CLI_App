package main

import(
  "os/exec"
  "fmt"
  "os"
)


func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}




func CmdExit()error{
  
  os.Exit(0)

  return nil 
}


func CommandHelp()error{

  fmt.Println("use -<options> to run commands \nuse -help for Help")
  return nil
}


