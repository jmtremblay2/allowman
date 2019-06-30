
package banks

import (
  "github.com/kniren/gota/dataframe"
  "scanCSV"
  _"fmt"
)

/*
 * BOA
 */

func coalesce(s1, s2 []string) []string{
  s := make([]string, len(s1))
  copy(s, s1)
  for i, value := range(s2){
    if s[i] == "" || s[i] == "NaN"{
      s[i] = value
    }
  }
  return s
}

type Cap1Stmt struct{
  DF dataframe.DataFrame
}

func (cs Cap1Stmt) GetDates() []string {
  return []string{
    "7/4/1776",
    scanCSV.GetDFElem(cs.DF,"date",0),
    scanCSV.GetDFElem(cs.DF,"date", cs.DF.Nrow()-1),
  }
}

func (cs Cap1Stmt) GetDescr() []string {
  return scanCSV.GetDFElems(cs.DF, "description")
}

func (cs Cap1Stmt) GetAmount() []string {
  debit := scanCSV.GetDFElems(cs.DF,"debit")
  credit := scanCSV.GetDFElems(cs.DF,"credit")
  
  return coalesce(debit, credit)
 
}

func CreateCap1Stmt(fname string) Cap1Stmt { 
  allDFs := scanCSV.ProcessCSVFile(fname)
  cs := Cap1Stmt{
    DF: allDFs[0].Rename("date","posteddate"),
  }
  return cs
}

