package mdbclient

import (
  "log"
  "time"
  "context"
  "github.com/Gridmax/Hillock/utility/messages"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  "go.mongodb.org/mongo-driver/bson"
  
  "github.com/Gridmax/Hillock/utility/configload"

  //"go.mongodb.org/mongo-driver/mongo/readpref"
)


func InsertData(data string) {

  config, err := configload.LoadConfig("config.yaml")
  
  _, kv, err := messages.ConvertToJSONAndKeyValue(data)
  client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://"+config.DatabaseUrl))
  if err != nil {

    panic(err)
  }
  
  ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
  err = client.Connect(ctx)
  if err != nil {
        log.Fatal(err)
  }
  defer client.Disconnect(ctx)
  
  databaseConnects := client.Database("hillock")
  sentinelCollection := databaseConnects.Collection("sentinelCollection")

  sentinelInsert, err := sentinelCollection.InsertOne(ctx, bson.D{
    {Key: "host", Value: kv["host_name"]},
    {Key: "hostgroup", Value: kv["host_group"]},
    {Key: "timestamp", Value: kv["timestamp"]},
    {Key: "cpu_total", Value: kv["cpu_total"]},
    {Key: "cpu_used", Value: kv["cpu_usage"]},
    {Key: "cpu_idle", Value: kv["cpu_idel"]},
    {Key: "ram_total", Value: kv["ram_total"]},
    {Key: "ram_used", Value: kv["ram_usage"]},
    {Key: "ram_idle", Value: kv["ram_usage"]},
    {Key: "uptime", Value: kv["uptime"]},
  }) 

  log.Println("MongoDB object successfully Inserted ", sentinelInsert)
  //fmt.Println(data)

}


