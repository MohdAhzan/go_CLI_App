
package utils

import(
  "os/exec"
  "fmt"
  "os"
)

//map the command corresponding to the functions
var CallBackMap =  map[string]func()error{

  "commandelp":CommandHelp,
  "cmdExit":CmdExit,
  "cmdClear":ClearScreen,

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


