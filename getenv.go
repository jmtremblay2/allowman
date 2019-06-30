package main

import (
  "fmt"
  "os"
  "path"
)

func main(){
  allowpath := os.Getenv("ALLOWMANPATH")
  fname := path.Join(allowpath,"/file.csv")
  fmt.Println(fname)
}
