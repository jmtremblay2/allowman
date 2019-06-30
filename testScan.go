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

  fname2 := "src/scanCSV/BOA-Checking-1_test.csv"
  ls,_ := scanCSV.ReadLines(fname2)
  fmt.Println(ls)
  for _,l := range ls{
    fmt.Println(l)
  }
}

