package banks

import (
  "github.com/kniren/gota/dataframe"
  "scanCSV"
)

/*
 * BOA
 */

type BOAStmt struct{
  DF dataframe.DataFrame
}


func (cs BOAStmt) GetDates() []string {
  return []string{
    "7/4/1776",
    scanCSV.GetDFElem(cs.DF,"date",0),
    scanCSV.GetDFElem(cs.DF,"date", cs.DF.Nrow()-1),
  }
}

func (cs BOAStmt) GetDescr() []string {
  return scanCSV.GetDFElems(cs.DF, "description")
}

func (cs BOAStmt) GetAmount() []string {
  return scanCSV.GetDFElems(cs.DF, "amount")
} 

// For now CreateChaseStmt create its own account at the same time
func CreateBOACheckStmt(fname string) BOAStmt { 
  allDFs := scanCSV.ProcessCSVFile(fname)
  cs := BOAStmt{
    DF: allDFs[1],
  }
  return cs
}

// For now CreateChaseStmt create its own account at the same time
func CreateBOACreditStmt(fname string) BOAStmt { 
  allDFs := scanCSV.ProcessCSVFile(fname)
  cs := BOAStmt{
    DF: allDFs[0].Rename("date","posteddate").Rename("description","payee"),
  }
  return cs
}

