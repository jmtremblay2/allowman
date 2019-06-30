package talktodb

import (
  "database/sql"
  _ "github.com/lib/pq"
  "fmt"
  "log"
)

/*
 * DATABASE
 */
type ConnData struct {
  Host string
  Port int
  User string
  Password string
  Dbname string
}

func ConnectToDB(cd ConnData) *sql.DB {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    cd.Host, cd.Port, cd.User, cd.Password, cd.Dbname)

  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  //defer db.Close()

  err = db.Ping()
  if err != nil {
    panic(err)
  }
  fmt.Println("Successfully connected!")
  return  db
}

/*
 * QUERIES
 */
func GetNewlyCreatedKey(tname string, keyname string, conn *sql.DB) int {
  // query to get the new key
  getIdQuery := fmt.Sprintf(`SELECT currval(pg_get_serial_sequence('%s','%s'))`,tname,keyname)

  // query
  var id int
  err := conn.QueryRow(getIdQuery).Scan(&id)
  if err != nil {
    log.Fatal(err)
  }

  return id
}

func RegisterHousehold(conn *sql.DB) int {
  createHHQuery := `INSERT INTO households(hid) VALUES (DEFAULT)`
  _, error := conn.Exec(createHHQuery)
  if error != nil{
    fmt.Println(error)
    panic(error)
  }
  return GetNewlyCreatedKey("households","hid",conn)
}

func RegisterAccount(conn *sql.DB, hid int, name string) int {
  createAcctQuery := `INSERT INTO accounts(accid, hid, name) VALUES (DEFAULT, $1, $2)`
  _, error := conn.Exec(createAcctQuery,hid,name)
  if error != nil{
    fmt.Println(error)
    panic(error)
  }
  return GetNewlyCreatedKey("accounts","accid",conn)
}

func RegisterStatement(conn *sql.DB, accid int, stmtdate, stmtstart, stmtend string) int {
  createStmtQuery := `INSERT INTO statements(stmtid, accid, stmtdate, stmtstart, stmtend)
                      VALUES (DEFAULT, $1, $2, $3, $4)`
  _, error := conn.Exec(createStmtQuery,accid,stmtdate, stmtstart, stmtend)//stmtdate)
  if error != nil{
    fmt.Println(error)
    panic(error)
  }
  return GetNewlyCreatedKey("statements","stmtid",conn)
}

func RegisterTransaction(conn *sql.DB, stmtid int, descr string, ammount string) int {
  // for now I leave it up to the DB to convert ammount to a float
  createTransQuery := `INSERT INTO transactions(trid, stmtid, descr, ammount)
                      VALUES (DEFAULT, $1, $2, $3)`
  _, error := conn.Exec(createTransQuery,stmtid,descr, ammount)
  if error != nil{
    fmt.Println(error)
    panic(error)
  }
  return GetNewlyCreatedKey("transactions","trid",conn)
}

