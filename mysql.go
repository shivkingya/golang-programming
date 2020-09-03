package main

import (
     "fmt"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

func insert(coinId string, coinSymbol string, coinname string) {
	fmt.Println("hello")
db, err :=sql.Open("mysql","root:admin@tcp(127.0.0.1:3306)/sampledb")

if err != nil {
 	fmt.Println("Connection error in mysql")
}
//db.Close()
fmt.Println("connection successfull")
  
//insert, err := db.Query("INSERT INTO dashboard_coin VALUES(",coinId,coinSymbol,coinname,")")
insert, err := db.Query("INSERT INTO table1 VALUES(7,'data2')")
if err != nil{
	panic(err.Error())
}

defer insert.Close()
fmt.Println("Insertion done successfully")

}

