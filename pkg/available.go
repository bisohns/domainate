package available

import (
    "github.com/deven96/gosock/pkg/custlog"
   )

func init(){
  // call logger in append mode
  defwriters := custlog.DefaultWriters(*LogFile, true)
  custlog.LogInit(defwriters)
}
