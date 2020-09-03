package main

import (
	 "fmt"
	 "time"
	"database/sql"
	_"github.com/go-sql-driver/mysql"  // go get -u github.com/go-sql-driver/mysql
)

func dbConn() (db *sql.DB) {
    dbDriver := "mysql"
    dbUser := "root"
    dbPass := "admin"
    dbName := "sampledb"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    if err != nil {
        panic(err.Error())
    }
    return db
}

func insert(id string, symbol string, name string)  {
   db:=dbConn()
   t := time.Now()
   created := t.Format("2006-01-02 15:04:05")
   fmt.Println(created)
   insForm, err := db.Prepare("INSERT INTO dashboard_coin(symbol,id,name,created) VALUES(?,?,?,?)")
   if err != nil {
	   panic(err.Error())
   }
   insForm.Exec(symbol,id,name,created)
  
if err != nil{
	panic(err.Error())
}
fmt.Println("Insertion done successfully")
//return "Insertion done successfully"
db.Close()
}

