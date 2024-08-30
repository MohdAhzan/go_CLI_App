
package utils 

type CliCommand struct{

  Key         string      `bson:"key"`
  Name        string      `bson:"name"`
  Description string      `bson:"description"`
  Callback    string      `bson:"callback"`
}

var Events = map[string]CliCommand{}

