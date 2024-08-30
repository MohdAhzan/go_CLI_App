
package utils

import(
  "os/exec"
  "fmt"
  "os"
  "github.com/MohdAhzan/go_CLI_App/apps/todo"

)

//map the command corresponding to the functions
var CallBackMap =  map[string]func()error{

  "commandHelp":CommandHelp,
  "cmdExit":CmdExit,
  "cmdClear":ClearScreen,
  "todo":todo.TodoStart,

}




func ClearScreen()error {
  cmd := exec.Command("clear")
  fmt.Println(cmd,"klsdfjkldsjfklsdjfldksjflsdlfkjldas")
  cmd.Stdout = os.Stdout
  err:=cmd.Run()
  if err!=nil{

    return err
  }
  return nil
}




func CmdExit()error{

  os.Exit(0)

  return nil 
}


func CommandHelp()error{

  fmt.Println("use -<options> to run commands \nuse -help for Help")
  return nil
}


