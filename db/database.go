package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/MohdAhzan/go_CLI_App/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB()(*mongo.Database,error){


  clientOptions:=options.Client().ApplyURI("mongodb://localhost:27017")
  client,err:=mongo.Connect(context.TODO(),clientOptions)
  if err!=nil{
    fmt.Println("error connecting Mongoshh")

    return nil,err
  }

  clientdb:=client.Database("MyApp")

  fmt.Println("Connected to Mongoooo")
  err=client.Ping(context.TODO(),nil)

  if err!=nil{
    log.Fatal(err)
  }

  return clientdb,nil

}

//accepts the collectionName string and returns true if data present else false
func checkDbCollection(db *mongo.Database,collectionName string) (bool,error){
  coll, err := db.ListCollectionNames(context.Background(),bson.D{})
  if err!=nil{
    return false,err
  }

  for _,c:=range coll{
    if c==collectionName{
      return true,nil
    }
  }
  return false,nil
}



func AddAppCmd(db *mongo.Database)error{

  exists,err :=checkDbCollection(db,"Clicommand")
  if err!=nil{
    return err
  }
  if !exists{

    models.Events = map[string]models.CliCommand{
      "-help": {
        Key: "-help",
        Name:        "Welcome to MyApp!!\n",
        Description: "-help : Lists help \n-exit : Exit from My App\nHit the <space> and Get a Wish",
        Callback:    "commandHelp",
      },
      "-exit": {
        Key: "-exit",
        Name:        "\n",
        Description: "Displays help message here. FOR now please help me\n",
        Callback:    "cmdExit",
      },
    }

    coll:=db.Collection("Clicommand")

    for _,cmd:=range models.Events{

      res, err:=coll.InsertOne(context.TODO(),cmd)
      if err!=nil{
        return err
      }else{
        fmt.Println("insert result",res)
      }
    }

  }else{
    fmt.Println("already Have the data in dbs")
  }

  return nil

}
func FetchData (DB *mongo.Collection,command string)(models.CliCommand,error){

  ctx,cancel:=context.WithTimeout(context.Background(),time.Second*5) 
  defer cancel() 
  var data models.CliCommand


  filter:= bson.D{
    {Key:"key",Value:command},
  }

  err:=DB.FindOne(ctx,filter).Decode(&data)
  if err!=nil{
    return models.CliCommand{},fmt.Errorf("Invalid Command... Try '-help' ")
  }

  return data ,nil

}


// func loadCommandsFromMongoDB(DB *mongo.Database) (map[string]models.CliCommand, error) {
//   collection := DB.Collection("Clicommand")
//
//   cursor, err := collection.Find(context.TODO(), bson.D{})
//   if err != nil {
//     return nil, err
//   }
//   defer cursor.Close(context.TODO())
//
//   events := make(map[string]models.CliCommand)k
//
//   for cursor.Next(context.TODO()) {
//     var result struct {
//       Key         string `bson:"key"`
//       Name        string `bson:"name"`
//       Description string `bson:"description"`
//       Callback    string `bson:"callback"`
//     }
//     err := cursor.Decode(&result)
//     if err != nil {
//       return nil, err
//     }
//
//     // Map the callback string back to the actual function
//     var callbackFunc func()
//     switch result.Callback {
//     case "commandHelp":
//       callbackFunc = CommandHelp
//     case "cmdExit":
//       callbackFunc = cmdExit
//     default:
//       callbackFunc = nil
//     }
//
//
//     events[result.Key] = models.CliCommand{
//       Name:        result.Name,
//       Description: result.Description,
//       Callback:    callbackFunc,
//     }
//   }
//
//   return events, nil
// }
//