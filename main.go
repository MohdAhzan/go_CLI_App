package main

import (
  "time"
	"bufio"
	"fmt"
	"os"
)

func main(){
  initialize() 
  
    err:=start()
    if err !=nil{
      fmt.Println(err)

  }
}


func start()error{

  scanner:=bufio.NewScanner(os.Stdin)

  for scanner.Scan(){
    command:=scanner.Text()

    data,err:=commandEvents(command)
    if err!=nil{
      return err 
    }

    fmt.Print(data.name,data.description,data.callback())

    time.Sleep(time.Second*3)
  }

  return nil
}

type CliCommand struct{
  name string
  description string
  callback func()error

}
var events = map[string]CliCommand{}

func initialize(){
  
  fmt.Print("here reaches")

  events["-h"]=CliCommand{
    name : "help\n",
    description : "Displays help message here . FOR now please help mee\n",
    callback: commandHelp,
  }
  events["help"]=CliCommand{
    name : "help\n",
    description : "Displays help message here . FOR now please help mee\n",
    callback: commandHelp,
  }

}
func commandHelp()error{
  
  fmt.Println("nothing for now")
  return nil
}

func  commandEvents(command string)(CliCommand,error){
  
  if command==""{
  
    return CliCommand{},fmt.Errorf("Invalid command... MotherF**ker")
  }
  
  data:=events[command]


  return data,nil

}



