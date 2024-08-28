package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"

	clientDb "github.com/MohdAhzan/go_CLI_App/db"
	"github.com/MohdAhzan/go_CLI_App/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func main(){


  err:=start()
  if err !=nil{
    fmt.Println(err)

  }


}



func start()error{

  db,err:=clientDb.ConnectDB()
  if err!=nil{
    fmt.Println("error here in connecting ")
    log.Fatal(err)
  }


  if err=clientDb.AddAppCmd(db);err!=nil{
    fmt.Println("error here")
    log.Fatal(err)

  }


  fmt.Print("MyApp >  ")

  scanner:=bufio.NewScanner(os.Stdin)
  wg:=  &sync.WaitGroup{}

  for scanner.Scan(){
    command:=scanner.Text()

    if command == "-exit"{

      return nil
    }
    data,err:=commandEvents(command,wg,db)
    wg.Done()
    if err!=nil{
      return err 
    }
    wg.Wait()
    // if len(data.Callback)>0{
    //
    //   if data.Callback==command{
    //
    //
    //
    //   } 
    //
    // }


    fmt.Println(data.Name," ")
    fmt.Println(data.Description," ")
  }

  return nil
}



// need to get the data from mongo according to the users input commands 

func  commandEvents(command string,wg *sync.WaitGroup,db *mongo.Database)(models.CliCommand,error){
  //
  wg.Add(1)
  //   if command==" "{
  //     return models.CliCommand{Name: "Surprise... MTF**ker....",Description: "",Callback: nil, },nil
  //   }
  //

  coll:=db.Collection("Clicommand")  
    
  
  data,err:= clientDb.FetchData(coll,command)
  if err!=nil{
    return models.CliCommand{},err
  }
  // data,ok:=models.Events[command]
  // if !ok{
  //   return models.CliCommand{Name: "Invalid Command ;)\n",Description: "",Callback:"commandHelp", },nil
  // }
  //
  //
  //
  return data,nil

}



