package main

import (
  "fmt"
  "errors"
)

func returnError(a, b int) (error) {
  if a == b {
    return errors.New("Error in returnError() function!")
  }
  return nil
}

func main() {
  err := returnError(1, 2)
  if err == nil {
    fmt.Println("reutrnError() ended normally!")
  } else {
    fmt.Println(err)
  }

  err = returnError(10, 10)
  if err == nil {
    fmt.Println("reutrnError() ended normally!")
  } else {
    fmt.Println(err)
  }


  if err.Error() == "Error in returnError() function!" {
    fmt.Println("!!")
  }
}

