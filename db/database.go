package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/MohdAhzan/go_CLI_App/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("error connecting Mongoshh")

		return nil, err
	}

	clientdb := client.Database("MyApp")

	// fmt.Println("Connected to Mongoooo")
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return clientdb, nil
}

// accepts the collectionName string and returns true if data present else false
func checkDbCollection(db *mongo.Database, collectionName string,updatedCount int64) (bool, error) {

   filter :=bson.M{}

  latestCount,err  :=db.Collection(collectionName).CountDocuments(context.Background(),filter)
  if err!=nil{
    return false,err 
  }
  
  // fmt.Println(updatedCount,"\n",latestCount)
  if latestCount!=updatedCount{
    db.Collection(collectionName).Drop(context.Background())
    // fmt.Println("change in collection droping old collection to add new collection")
    return false,nil
  }
  return true,nil

}

func checkIfCollectionUpdated()(int64,error){
  
 return int64(len(utils.Events)),nil

}



func AddAppCmd(db *mongo.Database) error {

  updatedCount,err :=checkIfCollectionUpdated()
  if err!=nil{
    return err
  }

  exists, err := checkDbCollection(db, "Clicommand",updatedCount)
	if err != nil {
		return err
	}

  if !exists {

    coll := db.Collection("Clicommand")

    for _, cmd := range utils.Events {

      res, err := coll.InsertOne(context.TODO(), cmd)
      if err != nil {
        return err
      } else {
        fmt.Println("insert result", res)
      }
    }
	}

	return nil
}




func FetchData(DB *mongo.Collection, command string) (utils.CliCommand, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var data utils.CliCommand

	filter := bson.D{
		{Key: "key", Value: command},
	}

	err := DB.FindOne(ctx, filter).Decode(&data)
	if err != nil {
		return utils.CliCommand{}, fmt.Errorf("Invalid Command... Try 'help' ")
	}
	// fmt.Println("dataa fetched",data)
	if data.Callback != "" {

		callbackFunc := utils.CallBackMap[data.Callback]

		err := callbackFunc()
		if err != nil {
			return utils.CliCommand{}, fmt.Errorf("Error calling the callback functions\n%v", err)
		}

	}

	return data, nil
}

