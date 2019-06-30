package banks

import(
  "database/sql"
  "strings"
  "talktodb"
)

/*
 * BANK STATEMENTS
 */

type BankStmt interface {
  GetDates() []string
  GetDescr() []string
  GetAmount() []string
}

func ProcessBankStmt(bs BankStmt, conn *sql.DB, accid int){
  stmtid := RegisterStatement(bs, conn, accid)
  RegisterTransactions(bs, conn, stmtid)
}

/*
 * HOUSEKEEPING
 */
func GetAcctName(fname string) string {
  tokens := strings.Split(fname,"/")
  name := tokens[len(tokens) - 1]
  name = strings.Replace(name,".CSV","",1)
  name = strings.Replace(name,".csv","",1)
  return name
}

func RegisterStatement(bs BankStmt, conn *sql.DB, accid int) int {
  dates := bs.GetDates()
  stmtid := talktodb.RegisterStatement(conn, accid, dates[0], dates[1], dates[2])
  return stmtid
}

func RegisterTransactions(bs BankStmt, conn *sql.DB, stmtid int){
  //stmtid := cs.GetStmtid()
  descr := bs.GetDescr()
  amount := bs.GetAmount()

  for i := 0; i < len(descr); i++ {
    talktodb.RegisterTransaction(conn, stmtid, descr[i], amount[i])
  }
}

