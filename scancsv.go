package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "regexp"
    "strings"
    "github.com/kniren/gota/dataframe"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
    file, err := os.Create(path)
    if err != nil {
        return err
    }
    defer file.Close()

    w := bufio.NewWriter(file)
    for _, line := range lines {
        fmt.Fprintln(w, line)
    }
    return w.Flush()
}

func countFields(lines []string) []int {
   count := make([]int, len(lines))

    sep := regexp.MustCompile(",")

    for i, line := range lines {
      matches := sep.FindAllStringIndex(line, -1)
      count[i] = len(matches)
      fmt.Println(count[i],line)
    }
    return count
}

func splitTables(lines []string) [][]string{
  count := countFields(lines)
  startnew := true
  tablenum := -1
  tables := make([][]string, 0)
  for i, line := range lines {
    // if line is empty it means next non empty line is the start of a table
    if 0 == count[i] {
      startnew = true
      continue
    }
    // initialize the new table
    if startnew {
      tablenum++
      tables = append(tables, make([]string, 0))
    }
    tables[tablenum] = append(tables[tablenum], line)

    startnew = false
  }
  return tables
}

func splitTokens(lines []string) [][]string{
  table := make([][]string, len(lines))
  for i, line := range lines {
    fmt.Println("")
    table[i] = strings.Split(line, ",")
    for j,token := range table[i]{
      table[i][j] = strings.ReplaceAll(token,"\"","")
    }
  }
  return table
}

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

func deleteEmptyColumns(tokens [][]string) [][]string{
  // figure out which colums are empty -- for now it's columns with an empty colnum name
  emptyCols := make([]bool, len(tokens[0]))
  nEmpty := 0
  for c, colname := range tokens[0]{
    isEmpty := "" == colname
    emptyCols[c] = isEmpty
    if isEmpty {
      nEmpty++
    }
  }


  rowLength := len(tokens[0]) - nEmpty
  newTokens := make([][]string, len(tokens))
  for r, row := range tokens {
    newTokens[r] = make([]string, rowLength)
    index := 0
    for c, col := range row {
      if ! emptyCols[c] {
        newTokens[r][index] = col
        index++
      }
    }
  }
  return newTokens

}

func main() {
    lines, err := readLines("/home/jm/allowman/bankcreditcardtransactions/BOA-Checking-1.csv")
    if err != nil {
        log.Fatalf("readLines: %s", err)
    }

    tables := splitTables(lines)
    printSlice(tables[0])
    fmt.Println("\n\n")
    printSlice(tables[1])

    fmt.Println("\n\n")

    df := dataframe.LoadRecords(
    [][]string{
        []string{"A", "B", "C", "D"},
        []string{"a", "4", "5.1", "true"},
        []string{"k", "5", "7.0", "true"},
        []string{"k", "4", "6.0", "true"},
        []string{"a", "2", "7.1", "false"},
    })
    fmt.Println(df)

   df2 := dataframe.LoadRecords(deleteEmptyColumns(splitTokens(tables[0])))
   fmt.Println(df2)
    /*
    for i, line := range lines {
        fmt.Println(i, line)
    }

    if err := writeLines(lines, "~/foo.out.txt"); err != nil {
        log.Fatalf("writeLines: %s", err)
    }
    */
}
