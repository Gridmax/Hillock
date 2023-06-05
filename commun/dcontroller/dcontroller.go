package dcontroller

import (
  "github.com/Gridmax/Hillock/connects/mongodb/mdbclient"
  "log"
)

func DataFlow(dflow string, message string){
  if dflow == "mongodb"{
    log.Println("Messages have submitted to data flow for", dflow)
    //log.Println(message)
    mdbclient.InsertData(message)
  }else if dflow == "mysql"{
  }else if dflow == "queue"{
  }else if dflow == "cloud"{
  }else{
  }
}
