package generate

import (
  "fmt"
  "math/rand"
  "bufio"
  "time"
  "github.com/deven96/gosock/pkg/custlog"
 )

 func init(){
   defwriters := custlog.DefaultWriters(*LogFile, true)
   custlog.LogInit(defwriters)
}

var tlds = []string{
  "com",
  "net"
}

const allowedChars = "abcdefghijklmnopqrstuvwxyz01234556789_-"

func transform(keyword string) []string{
  return []string{
    keyword,
    keyword + "app",
    "go" + keyword,
    keyword + "site",
    keyword + "time",
    "get" + keyword,
    keyword + "hq",
    "lets" + keyword,
  }
 }

func GenerateDomains(keyword string){
  rand.Seed(time.Now().UTC().UnixNano())
  wordarray := transform(keyword)
  fmt.Printf("%v\n", wordarray)
 }

