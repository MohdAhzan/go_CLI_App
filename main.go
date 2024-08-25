package main

import (
	"bufio"
	"fmt"
	"os"
  "sync"
)

func main(){
  initialize() 
  
    err:=start()
    if err !=nil{
      fmt.Println(err)

  }
}


func start()error{

  fmt.Print("MyApp >  ")

  scanner:=bufio.NewScanner(os.Stdin)
  wg:=  &sync.WaitGroup{}

  for scanner.Scan(){
    command:=scanner.Text()
    
    if command == "-exit"{

      return nil
    }
    data,err:=commandEvents(command,wg)
  wg.Done()
    if err!=nil{
      return err 
    }
    wg.Wait()
    if data.callback!=nil{
      
        data.callback()
      
    }

    
  
    fmt.Println(data.name," ")
    fmt.Println(data.description," ")
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
  

  events["-help"]=CliCommand{
    name : "Welcome to MyApp!!\n",
    description : "-help :  Lists help \n-exit : Exit from My App\nHit the <space> and Get a Wish",
    callback: commandHelp,
  }
  events["-exit"]=CliCommand{
    name : "\n",
    description : "Displays help message here . FOR now please help mee\n",
    callback: cmdExit,
  }
   

}

func cmdExit()error{
    
  return nil 
}


func commandHelp()error{
  

  fmt.Println("use -<options> to run commands \nuse -help for Help")
  return nil
}


func  commandEvents(command string,wg *sync.WaitGroup)(CliCommand,error){
 
  wg.Add(1)
  if command==" "{
    return CliCommand{name: "Surprise... MTF**ker....",description: "",callback: nil, },nil
  }

  data,ok:=events[command]
  if !ok{
    return CliCommand{name: "Invalid Command ;)\n",description: "",callback:commandHelp, },nil
  }
 
  

  return data,nil

}



