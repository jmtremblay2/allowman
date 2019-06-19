package main

import (
  "scanCSV"
  "fmt"
  //"log"
  //"github.com/kniren/gota/dataframe"

)




func main() {
  fname := "/home/jm/allowman/bankcreditcardtransactions/BOA-Checking-1.csv"
  DF := scanCSV.ProcessCSVFile(fname)
  fmt.Println(DF)
}

