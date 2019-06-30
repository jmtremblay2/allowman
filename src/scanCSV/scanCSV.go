package scanCSV

import (
    "bufio"
    _"fmt"
    "log"
    "os"
    "regexp"
    "strings"
    "github.com/kniren/gota/dataframe"
)

type FileLine = string
type FileLines = []string
type CSVField = string
type CSVTable = [][]CSVField

//remove spaces from strings
func RemoveSpaces(s []string) []string{
  s2 := make([]string, len(s))
  for i, ss := range(s){
    s2[i] = strings.Replace(ss," ","",-1)
  }
  return s2
}

// cast strings to lwer case
func ToLower(s []string) []string{
  s2 := make([]string, len(s))
  for i, ss := range(s){
    s2[i] = strings.ToLower(ss)
  }
  return s2
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func ReadLines(path string) (FileLines, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var lines FileLines
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, scanner.Err()
}

func CountFieldsOneLine(line FileLine) int {
  sep := regexp.MustCompile(",")
  matches := sep.FindAllStringIndex(string(line), -1)
  return len(matches)
}

// countFields counts how many comas are in each line of lines
// and return a slice of the results
// for a csv that helps identify when a text file contains multiple sub tables
func CountFields(lines FileLines) []int {
  count := make([]int, len(lines))
   for i, line := range lines {
    count[i] = CountFieldsOneLine(string(line))
  }
  return count
}

// splitTables 
func SplitTables(lines FileLines) []FileLines {
  count := CountFields(lines)
  startnew := true
  tablenum := -1
  tables := make([]FileLines, 0)
  for i, line := range lines {
    // if line is empty it means next non empty line is the start of a table
    if 0 == count[i] {
      startnew = true
      continue
    }
    // initialize the new table
    if startnew {
      tablenum++
      tables = append(tables, make(FileLines, 0))
    }
    tables[tablenum] = append(tables[tablenum], line)

    startnew = false
  }
  return tables
}

func SplitTokens(lines FileLines) CSVTable {
  table := make(CSVTable, len(lines))
  for i, line := range lines {
    table[i] = strings.Split(line, ",")
    for j,token := range table[i]{
      table[i][j] = strings.ReplaceAll(token,"\"","")
    }
  }
  return table
}

/*
func printSlice(lines []string) {
  for _, line := range lines{
    fmt.Println(line)
  }
}

func print2Dstring(tokens [][]string){
  for _, row := range tokens{ 
    for _, token := range row{
      fmt.Print(token)
      fmt.Print("  --  ")
    }
    fmt.Println("")
  }
}
*/

func DeleteEmptyColumns(table CSVTable) CSVTable {
  // figure out which colums are empty -- for now it's columns with an empty colnum name
  emptyCols := make([]bool, len(table[0]))
  nEmpty := 0
  for c, colname := range table[0]{
    isEmpty := "" == colname
    emptyCols[c] = isEmpty
    if isEmpty {
      nEmpty++
    }
  }

  rowLength := len(table[0]) - nEmpty
  newTable := make([][]string, len(table))

  for r, row := range table {
    newTable[r] = make([]string, rowLength)
    index := 0

    for c, col := range row {
      if ! emptyCols[c] {
        newTable[r][index] = col
        index++
      }
    }
  }
  return newTable

}

func CSVTable2DataFrame(table CSVTable) dataframe.DataFrame {
   df2 := dataframe.LoadRecords(DeleteEmptyColumns(table))
   return df2
}

func ProcessCSVFile(fname string) []dataframe.DataFrame {
    stmtLines, err := ReadLines(fname)
    if err != nil {
        log.Fatalf("readLines: %s", err)
    }

    stmtTables := SplitTables(stmtLines)

    tables := make([]CSVTable, len(stmtTables))
    DFs := make([]dataframe.DataFrame, len(stmtTables))
    for i,t := range stmtTables{
      tables[i] = SplitTokens(t)
      // format variable names to kill white spaces and upper case letters
      tables[i][0] = ToLower(RemoveSpaces(tables[i][0]))
      DFs[i] = CSVTable2DataFrame(tables[i])
    }
  return DFs
}

func GetDFElem(DF dataframe.DataFrame, colname string, index int) string {
  return DF.Col(colname).Elem(index).String()
}

func GetDFElems(DF dataframe.DataFrame, colname string) []string {
  nrow := DF.Nrow()
  elems := make([]string, DF.Nrow())
  for i := 0; i < nrow; i++ {
    elems[i] = GetDFElem(DF, colname, i)
  }
  return elems
}

