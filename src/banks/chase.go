package banks

import (
  "github.com/kniren/gota/dataframe"
  "scanCSV"
)

/*
 * CHASE
 */
type ChaseStmt struct {
  DF dataframe.DataFrame
  //stmtid int
}

func (cs ChaseStmt) GetDates() []string {
  return []string{
    "7/4/1776",
    scanCSV.GetDFElem(cs.DF,"transactiondate",0),
    scanCSV.GetDFElem(cs.DF,"transactiondate", cs.DF.Nrow()-1),
  }
}

func (cs ChaseStmt) GetDescr() []string {
  return scanCSV.GetDFElems(cs.DF, "description")
}

func (cs ChaseStmt) GetAmount() []string {
  return scanCSV.GetDFElems(cs.DF, "amount")
}

func CreateChaseStmt(fname string) ChaseStmt { 
  allDFs := scanCSV.ProcessCSVFile(fname)
  cs := ChaseStmt{
    DF: allDFs[0],
  }
  return cs
}

