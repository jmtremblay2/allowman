
package main

import (
  "talktodb"
  "banks"
  "fmt"
  "path"
  "os"
)


func main() {
  /* DB connection stuff
   */ 
  cd := talktodb.ConnData{  
    Host: "localhost",
    Port: 5432,
    User: "postgres",
    Password: "postgres",
    Dbname: "allowman",
  }

  /* files for our first pass
   */
   allowmanpath := os.Getenv("ALLOWMANPATH")
  stmtfolder := "bankcreditcardtransactions"
  var fnames = []struct{
    fname string
    bank string
    // accid is ugly, we'll put application code to handle that eventually but 
    // for now we just assign something
    accid int}{
      {path.Join(allowmanpath,stmtfolder,"/Chase-JM-CC.CSV"),"chase",-1},
      {path.Join(allowmanpath,stmtfolder,"/BOA-Lynna-CC.csv"),"boacredit",-1},
      {path.Join(allowmanpath,stmtfolder,"/BOA-Checking-2.csv"),"boacheck",-1},
      {path.Join(allowmanpath,stmtfolder,"/BOA-Joint-CC.csv"),"boacredit",-1},
      {path.Join(allowmanpath,stmtfolder,"/Chase-Amazon-CC.CSV"),"chase",-1},
      {path.Join(allowmanpath,stmtfolder,"/BOA-Checking-1.csv"),"boacheck",-1},
      {path.Join(allowmanpath,stmtfolder,"/CO-Checking.csv"),"cap1",-1},
      {path.Join(allowmanpath,stmtfolder,"/CO-Parents-CC.csv"),"cap1",-1},
    }

  db := talktodb.ConnectToDB(cd)
  defer db.Close()

  /* create now HOUSEHOLD
   */
  hid := talktodb.RegisterHousehold(db)

  /* create new ACCOUNTS
   */
  for k, fname := range fnames{
    accountName := banks.GetAcctName(fname.fname)
    fnames[k].accid = talktodb.RegisterAccount(db, hid, accountName)
  }

  for _, fname := range fnames{
    fmt.Println(fname)
    switch fname.bank{
      case "chase": 
        cs := banks.CreateChaseStmt(fname.fname)
        banks.ProcessBankStmt(cs, db, fname.accid)
      case "boacheck":
        cs := banks.CreateBOACheckStmt(fname.fname)
        banks.ProcessBankStmt(cs, db, fname.accid)
      case "boacredit":
        cs := banks.CreateBOACreditStmt(fname.fname)
        banks.ProcessBankStmt(cs, db, fname.accid)
      case "cap1":
        cs := banks.CreateCap1Stmt(fname.fname)
        banks.ProcessBankStmt(cs, db, fname.accid)
    }
  }
}  
