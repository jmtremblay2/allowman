package scanCSV

import (
  "testing"
  _ "fmt"
)


func TestReadLines(t *testing.T) {
  fname1 := "BOA-Checking-1_test.csv" 
  lines1 := FileLines{
    "Description,,Summary Amt.",
    "Beginning balance as of 05/30/2019,,\"500.00\"",
    "Total credits,,\"200.00\"",
    "Total debits,,\"100.00\"",
    "Ending balance as of 01/01/2019,,\"600.00\"",
    "",
    "Date,Description,Amount,Running Bal.",
    "05/30/2019,Beginning balance as of 05/30/2019,,\"500.00\"",
    "06/04/2019,\"PAYPAL DES:INST XFER ID\",\"-75.00\",\"400.30\"",
    "06/07/2019,\"PATH-E-TECH TECHOLOGIES\",\"250.00\",\"650.30\"",
  }
  var tests = []struct{
    fname string
    lines FileLines
  }{
    {fname1, lines1},
  }

  for _,test := range tests{
    lines,_ := ReadLines(test.fname)
    for i,line := range lines{
      if test.lines[i] != line {
        t.Errorf("error in ReadLines(%q)",test.fname)
      }
    }
  }

}


func TestCountFieldsOneLine(t *testing.T) {
  var tests = []struct{
    line FileLine
    count int
  }{
    {"",0},
    {",,,",3},
    {"1,1,1,1",3},
    {",111,111,111,",4},
    {",,3,,3,,",6},
    {"            ",0},
    {"",0},
    {",,,",3},
    {"1,b,b",2},
    {",1,2,3,",4},
    {",,,1,,,",6},
    {"123,345,432",2},
    {",aaa,aaa,aaa,aaa,",5},
  }

  for _,test := range tests{
    if CountFieldsOneLine(test.line) != test.count {
      t.Errorf("error in CountFieldsOneLine(%q)",test.line)
    }
  }
}

// we already test CountFieldsOneLine
// func TestCountFields(t *testing.T)

func TestSplitTables(t *testing.T) {
  lines1 := FileLines{
    "Description,,Summary Amt.",
    "Beginning balance as of 05/30/2019,,\"500.00\"",
    "Total credits,,\"200.00\"",
    "Total debits,,\"100.00\"",
    "Ending balance as of 01/01/2019,,\"600.00\"",
    "",
    "Date,Description,Amount,Running Bal.",
    "05/30/2019,Beginning balance as of 05/30/2019,,\"500.00\"",
    "06/04/2019,\"PAYPAL DES:INST XFER ID\",\"-75.00\",\"400.30\"",
    "06/07/2019,\"PATH-E-TECH TECHOLOGIES\",\"250.00\",\"650.30\"",
  }
  linesSplit1 := []FileLines{
    {
    "Description,,Summary Amt.",
    "Beginning balance as of 05/30/2019,,\"500.00\"",
    "Total credits,,\"200.00\"",
    "Total debits,,\"100.00\"",
    "Ending balance as of 01/01/2019,,\"600.00\""},
    {
    "Date,Description,Amount,Running Bal.",
    "05/30/2019,Beginning balance as of 05/30/2019,,\"500.00\"",
    "06/04/2019,\"PAYPAL DES:INST XFER ID\",\"-75.00\",\"400.30\"",
    "06/07/2019,\"PATH-E-TECH TECHOLOGIES\",\"250.00\",\"650.30\""},
  }
  
  var tests = []struct{
    lines FileLines
    linesSplit []FileLines
  }{
    {lines1, linesSplit1},
  }

  for i,test := range tests{
    ls := SplitTables(test.lines)
    for i,line := range ls{
      if test.linesSplit[i] != line {
        t.Errorf("error in SplitLines")
      }
    }
  }

}
func TestSplitTokens(t *testing.T) {

}
func TestDeleteEmptyColumns(t *testing.T) {

}

// too complicated to test now
//func TestCSVTable2DataFrame(t *testing.T)

// too complicated to tests for now
//func TestProcessCSVFile(t *testing.T)
