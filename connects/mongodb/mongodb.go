package mongodb

import (
  "fmt"

  "context"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  //"go.mongodb.org/mongo-driver/mongo/readpref"
)


func InsertData(data string) {

  client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
  if err != nil {
    panic(err)
  }
  fmt.Println(client)
  fmt.Println(client)
}

