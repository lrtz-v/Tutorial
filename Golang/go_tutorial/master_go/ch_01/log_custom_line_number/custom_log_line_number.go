package main

import (
  "fmt"
  "log"
  "os"
)

var LOG_FILE = "/tmp/tmp.log"

func main() {
  f, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
  if err != nil {
    fmt.Println(err)
    return
  } 
  defer f.Close()

  iLog := log.New(f, "customLogWithLineNumber ", log.LstdFlags)
  iLog.SetFlags(log.LstdFlags|log.Lshortfile)
  iLog.Println("Hello there!")
  iLog.Println("Another log entry!")

}