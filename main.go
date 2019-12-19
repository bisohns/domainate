package main

import (
   "flag"
   "github.com/deven96/gosock/pkg/custlog"
   "github.com/bisoncorps/domainate/pkg/generate"
  )

  var LogFile = flag.String("log", "domainate.log", "Name of the log file to save to")


  func main(){
    flag.Parse() //parse the flags
//    get the default writers for log
    defwriters := custlog.DefaultWriters(*LogFile, false)
    custlog.LogInit(defwriters)
    custlog.Info.Println("Domainate started")
    generate.GenerateDomains("bisoncorps")
  }
